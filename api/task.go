package api

import (
	"github.com/jiangchengzi/mycloud/etcd"
	"encoding/json"
	"github.com/golang/glog"
	"github.com/pkg/errors"
)

func (t *Task) CheckTaskParams() error{
	if t.TaskName == "" {
		glog.Infoln("CheckTaskParams","taskName is empty")
		return errors.New("taskName is empty")
	}
	e := etcd.GetClient()
	if e.TaskExists(t.TaskName){
		glog.Infoln("CheckTaskParams","taskName has been existed")
		return errors.New("taskName has been existed")
	}
	//TODO：检查其他项，假如需要的话
	return nil
}

func (t *Task) Insert() error{
	e := etcd.GetClient()
	if err := e.InsertTask(t);err != nil{
		return err
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