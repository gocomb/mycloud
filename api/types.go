package api

type Task struct {
	TaskID      string       `json:"-"`
	TaskName    string       `json:"task_name"`
	// type 表示task的类型，目前规划有四种类型的task，分别是
	//  即时task:完成即时任务，不需要构建产物，更加关注执行结果
	//  构建task:完成构建任务，形成镜像，存放于镜像仓库
	//  部署task:完成部署任务，致力于将mycloud平台或者第三方平台的构建产物进行合理、高效地进行部署
	//  定时task:完成定时任务，满足用户定时触发任务执行的需求
	Type        string       `json:"type"`
	UserID      string       `json:"user_id,omitempty"`
	Namespace   string       `json:"namespace"`
	Description string       `json:"description,omitempty"`
	// State 表示task的状态，分为以下几种：
	//    Creating  刚开始创建
	//    Scheduled controller接收到处理指令，已经将task调度到某一个node上
	//    Running   task已经开始执行
	//    Successful task执行成功
	//    Failed    执行失败
	//    Pending   由于某些原因未能成功执行，处于暂态
	State       string		 `json:"-"`
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
	ID    string `json:"id"`
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

type CreateTaskResponse struct {
	TaskId       string `json:"taskId,omitempty"`
	Message      string `json:"message,omitempty"`
}

type GetTasksResponse struct {
	Message      string `json:"message"`
	Tasks        []Task `json:"tasks"`
	Meta         ListMeta `json:"meta"`
}

type ListMeta struct {
	Total int `json:"total"`
	Data  []interface{} `json:"data"`
}

type ErrorInternelResponse struct {
	ErrorMessage string `json:"message"`
}

