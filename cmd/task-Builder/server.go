package task_Builder

import (
	"github.com/spf13/pflag"
	"fmt"
	"os"
)

/*
为何存在？
首先，这一部分本身是可以放在api-server中实现的，但是出于以下考虑，将这一部分拧出来：
	1、分布式构建：在构建过程中需要下载用户代码，执行构建指令，生成构建产物，这些操作都非常消耗资源，
				 网络带宽、CPU、内存，将构建节点放在一台机器上，然后将构建产物放到仓库里，然后再执
				 行部署脚本，都很难满足实际需求，分布式构建可以充分协调集群中的资源分配，最大化集群的构建效率
	2、可靠性：....


任务划分
	1、响应构建请求：获取到任务之后，去访问api-server，获取到etcd中存储的task，解析task任务：
						1、git clone url、token
						2、调用docker api,利用Dockerfile进行build,
							TODO：改成允许用户自定义assemble.sh来进行构建
						3、调用docker api,将构建产物根据用户自定义策略push到私有镜像仓库或者公有镜像仓库
							Notes:如果是用户的私有镜像仓库，比如用户自己的机器上的registry，首先要保证可以访问，其次是需要用户填写task的时候
							也附带registry的用户名和密码，否则无法上传成功
						到此，task-builder的基本任务已经完成，用户可以在镜像仓库中看到已经构建好的产物，如果用户填写了部署策略的化，task-builder
						需要在构建完成之后，调用api-server的接口，传达开始部署的指令。

	2、通过websocket与controllers建立链接,传递heartbeat
*/

func main(){

	ops := NewTaskBuilderOption()
	ops.AddFlags(pflag.CommandLine)
	pflag.Parse()

	if err:=ops.Run();err!=nil{
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}











