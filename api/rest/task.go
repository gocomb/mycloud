package rest

import (
	"github.com/emicklei/go-restful"
	"github.com/jiangchengzi/mycloud/api"
	ResError "github.com/jiangchengzi/mycloud/error"
	"github.com/golang/glog"
)

func CreateTask(request *restful.Request,response *restful.Response){
	task := api.Task{}
	err := request.ReadEntity(&task)
	if err == nil{
		glog.Infoln("createTask","error when read entity from requst body")
		ResError.ErrorBadRequest(response,"error when read entity from requst body")
		return
	}
	//create task in etcd
	if err := task.CheckTaskParams();err !=nil {
		ResError.ErrorBadRequest(response,string(err))
		return
	}
	if err := task.Insert();err != nil{
		ResError.InternalServerError(response,string(err))
		return
	}
	response.WriteEntity(api.TaskResponse{
		TaskId:"123",
	})
}

