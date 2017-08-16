package rest

import (
	"github.com/emicklei/go-restful"
	"github.com/jiangchengzi/mycloud/api"
)

func CreateTask(request *restful.Request,response *restful.Response){

	response.WriteEntity(api.TaskResponse{
		ErrorMessage:"None",
		TaskId:"123",
	})
}