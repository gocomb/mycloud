package api

import (
	"encoding/json"
	"github.com/golang/glog"
)

type Task struct {
	TaskID      string       `json:"id,omitempty"`
	TaskName    string       `json:"task_name,omitempty"`
	UserID      string       `json:"user_id,omitempty"`
	UserName    string       `json:"user_name,omitempty"`
	Description string       `json:"description,omitempty"`
	Repository  Repository   `json:"repository,omitempty"`
	Hooks       []Hook       `json:"hook,omitempty"`
	Strategy    TaskStrategy `json:"strategy,omitempty"`
	OutCome     TaskOutCome  `json:"outcome,omitempty"`
	Children    []string     `json:"children,omitempty"` //taskIDs
}

type TaskOutCome struct {
	OutComeType string `json:"type"`
}

type TaskStrategy struct {
	Name    string      `json:"name"`
	Trigger TaskTrigger `json:"trigger"`
}

type TaskTrigger struct {
	TriggerType string `json:"type"`
}

type TaskFlow struct {
	Name  string `json:"name"`
	Tasks []Task `json:"tasks"`
}

type Hook struct {
	CallBack string `json:"callback,omitempty"`
	Token    string `json:"token,omitempty"`
}

type Repository struct {
	Url      string `json:"url,omitempty"`
	UserName string `json:"user_name,omitempty"`
	Password string `json:"password,omitempty"`
	Webhook  string `json:"webhook,omitempty"`
}

type TaskResponse struct {
	TaskId       string `json:"TaskId,omitempty"`
	ErrorMessage string `json:"ErrorMessage,omitempty"`
}

type ErrorInternelResponse struct {
	ErrorMessage string `json:"message"`
}

func Parse(from []byte) (to Task,err error){
	if err = json.Unmarshal(from,&to);err!=nil{
		glog.Info("error occured when unmarshal taskFlow")
		return
	}
	return
}