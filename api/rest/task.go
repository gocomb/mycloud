package rest

import (
	"github.com/emicklei/go-restful"
	"github.com/jiangchengzi/mycloud/api"
	ResError "github.com/jiangchengzi/mycloud/error"
	"github.com/golang/glog"
	"github.com/jiangchengzi/mycloud/util/id"
	"fmt"
)

func CreateTask(request *restful.Request,response *restful.Response){
	task := api.Task{}
	glog.Infoln(request.Attribute("UserID"))
	err := request.ReadEntity(&task)
	if err != nil{
		errMessage := "error when read entity from requst body"
		glog.Infoln("createTask",errMessage)
		ResError.ErrorBadRequest(response,errMessage)
		return
	}
	//create task in etcd
	if err := task.CheckTaskParams();err !=nil {
		ResError.ErrorBadRequest(response,fmt.Sprintf("errMessage:%v",err))
		return
	}
	task.TaskID = id.NewTaskID()
	if err := task.Insert();err != nil{
		ResError.InternalServerError(response,fmt.Sprintf("errMessage:%v",err))
		return
	}
	glog.Infoln("createTask","create task successfully",task)
	response.WriteEntity(api.CreateTaskResponse{
		TaskId:task.TaskID,
		Message:"create task successfully",
	})
}

func ListTasks(request *restful.Request,response *restful.Response){

}