package controller

import (
	"github.com/jiangchengzi/mycloud/etcd"
	"github.com/jiangchengzi/mycloud/util/queue"
	"github.com/jiangchengzi/mycloud/api"
	"context"
	"github.com/golang/glog"
	"sync"
	"time"
)

type TaskConfig struct {
	Status chan queue.Node
}

var (
	tc *TaskConfig
	q *queue.Queue
)



// 存放TaskFlow，TaskFlow是一串task的集合，可以由多种task进行任意结合，满足用户丰富的构建、部署需求
// 在TASKFLOWDIR目录下存放着完整的taskflow的yaml文档，api-server接收到taskflow的请求后，会将第一
// 批(parent)节点(task)同时存入到WORKDIR目录下，触发构建，在父节点完成构建后，task-builder会调用
// api-server提供的rest-api将其构建状态置成successful，controller会捕捉到这一动作，在相应逻辑处理代码
// 处，会分析此构建完的task是否需要部署，以及是否有父节点，如果有部署则调用rest-api将阶段置为deployment，状态
// 置为start,触发部署，如果有子节点，则从TASKFLOWDIR读出子节点(task),调用rest-api创建task节点
// 未来可能还需要开发一个状态管理器，收集taskflow中所有任务的状态，形成更丰富的状态调节机制
const TASKFLOWDIR string = `mycloud/taskflow`


func (ts *TaskConfig) Update() <-chan queue.Node{
	return ts.Status
}

func NewTaskConfig() *TaskConfig{
	return &TaskConfig{
		Status:make(chan queue.Node),
	}
}

func GetTaskConfig() *TaskConfig{
	return tc
}

func InitController(){
	e := etcd.GetClient()
	if !e.IsDirExist(etcd.WORKDIR){
		e.CreateDir(etcd.WORKDIR)
	}
	tasks,err:=e.ListNodes(etcd.WORKDIR)
	if err!=nil{
		glog.Errorf("list tasks error in etcd")
		return
	}
	q = queue.NewQueue()
	tc = NewTaskConfig()
	for _,task := range tasks{
		data,err:=api.Parse([]byte(task.Value))
		if err!=nil{
			glog.Errorf("init tasks queue error")
			return
		}
		q.Push(&queue.Node{
			Data:data,
		})
	}
}


func Run(stopCh <-chan struct{} ){
	go watchTaskFlow()
	go handleTasks(tc.Update())
	<-stopCh
}

func watchTaskFlow(){
	e := etcd.GetClient()
	w,err := e.CreateWatcher(etcd.WORKDIR)
	if err!=nil{
		glog.Infof("error occored,error is %v",err)
	}
	ctx := context.Background()
	wg := sync.WaitGroup{}
	wg.Add(1) //如果状态更新循环退出，则直接结束，所以设为1
	go func(){
		defer wg.Done()
		for {
			node,err := w.Next(ctx)
			if err!=nil{
				continue
			}
			task,err:= api.Parse([]byte(node.Node.Value))
			if err!=nil{
				continue
			}
			q.Push(&queue.Node{
				Data:task,
				Operation:node.Action,
			})
		}
	}()
	go SyncTaskStatus()
	wg.Wait()
}

func handleTasks(taskstatus <-chan queue.Node){
	for {
		select {
		case task:=<-taskstatus:
			switch task.Operation {
			case etcd.WatchActionCreate:
				glog.Infof("create")
			case etcd.WatchActionSet:
				glog.Infof("set")
			case etcd.WatchActionDelete:
				glog.Infof("delete")
			default:
				/*
					默认分支可能出现在以下情况(可以补充)：
						1、初始化时同步的状态，可能是由于平台突然间中断、重启等异常情况导致需要重新从etcd加载状态，此时的动作可能已经过期
						(处于expire)，那么我们需要去读取task的state,


				*/
			}
		default:
			time.Sleep(100*time.Millisecond)
		}
	}
}


func SyncTaskStatus(){
	t := time.NewTicker(1*time.Second)
	for {
		<-t.C
		task := q.Pop()
		if task == nil{
			continue
		}
		tc.Status <- *task
	}
}
