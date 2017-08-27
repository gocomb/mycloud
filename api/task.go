package api

import (
	"github.com/jiangchengzi/mycloud/etcd"
	"encoding/json"
	"github.com/golang/glog"
	"github.com/pkg/errors"
	"fmt"
	"path/filepath"
)

func (t *Task) CheckTaskParams() error{
	if t.TaskName == "" {
		errMessage := "taskName is empty"
		glog.Infoln("CheckTaskParams",errMessage)
		return errors.New(errMessage)
	}
	if t.TaskExists(){
		errMessage := "taskName has been existed"
		glog.Infoln("CheckTaskParams",errMessage)
		return errors.New(errMessage)
	}
	//TODO：检查其他项，假如需要的话
	return nil
}

func (t *Task) Insert() error{
	e := etcd.GetClient()
	taskKey := filepath.Join(etcd.WORKDIR,t.TaskID)
	taskValue,err := json.Marshal(t)
	if err != nil{
		errMessage := fmt.Sprintf("error occured when marshal task body,error message is:%v",err)
		glog.Infoln("InsertTask",errMessage)
		return errors.New(errMessage)
	}
	if err := e.Create(taskKey,string(taskValue));err !=nil {
		errMessage := fmt.Sprintf("error occured when insert task to etcd,error message is:%v",err)
		glog.Infoln("InsertTask",)
		return errors.New(errMessage)
	}
	return nil
}


func Parse(from []byte) (to Task,err error){
	if err = json.Unmarshal(from,&to);err!=nil{
		glog.Info("error occured when unmarshal task")
		return
	}
	return
}


func (t *Task) TaskExists() bool{
	e := etcd.GetClient()
	taskKey :=  filepath.Join(etcd.WORKDIR,t.TaskID)
	taskValue,err := e.Get(taskKey)
	if err!=nil{
		glog.Infoln("IsTaskExists","error occured when check whether task exists")
		return false
	}
	if "" == taskValue{
		return false
	}
	return true
}