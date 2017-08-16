package controller

import (
	"github.com/jiangchengzi/mycloud/api"
	"github.com/jiangchengzi/mycloud/etcd"
	"github.com/jiangchengzi/mycloud/util/queue"
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

const WORKDIR string = `mycloud/tasks`


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
	if !e.IsDirExist(WORKDIR){
		e.CreateDir(WORKDIR)
	}
	tasks,err:=e.ListNodes(WORKDIR)
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
	w,err := e.CreateWatcher(WORKDIR)
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
			glog.Infoln(node.Node,node.Node.Value)
			task,err:= api.Parse([]byte(`{"id":"xssxs"}`))
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
