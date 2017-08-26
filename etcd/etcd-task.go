package etcd

import (
	"github.com/golang/glog"
	"github.com/pkg/errors"
	"fmt"
	"github.com/jiangchengzi/mycloud/api"
	"encoding/json"
)

//存放需要处理的task
const WORKDIR string = `mycloud/tasks`

func (ec *Client) TaskExists(taskName string) bool{
	taskKey :=  WORKDIR + taskName
	taskValue,err := ec.Get(taskKey)
	if err!=nil{
		glog.Infoln("IsTaskExists","error occured when check whether task exists")
		return false
	}
	if "" == taskValue{
		return false
	}
	return true
}

func (ec *Client) InsertTask(task *api.Task) error{
	taskKey :=  WORKDIR + task.TaskID
	taskValue,err := json.Marshal(task)
	if err != nil{
		glog.Infoln("InsertTask","error occured when marshal task body,error message is:%v",err)
		return errors.New(fmt.Sprintf("error occured when marshal task body,error message is:%v",err))
	}
	if err := ec.Create(taskKey,string(taskValue));err !=nil {
		glog.Infoln("InsertTask","error occured when insert task to etcd,error message is:%v",err)
		return errors.New(fmt.Sprintf("error occured when insert task to etcd,error message is:%v",err))
	}
	return nil
}