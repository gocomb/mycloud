package main

import (
	"github.com/spf13/pflag"
	"fmt"
	"os"
	"github.com/emicklei/go-restful"
	"runtime"
	"math/rand"
	"net/http"
	"github.com/jiangchengzi/mycloud/api/rest"
	"github.com/jiangchengzi/mycloud/controller"
	"github.com/jiangchengzi/mycloud/etcd"
	"github.com/jiangchengzi/mycloud/util/logger"

	"time"
	"github.com/golang/glog"
	"k8s.io/apimachinery/pkg/util/wait"
)



func main(){

	ops:=NewApiOptions()
	ops.AddFlags(pflag.CommandLine)
	pflag.Parse()

	logger.InitLogs()
	ops.InitEtcd()
	controller.InitController()

	s:=NewApiServer(ops)

	if err:=s.Run();err!=nil{
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}


func(s *ApiServer) Run() error{

	runtime.GOMAXPROCS(runtime.NumCPU())

	rand.Seed(time.Now().UTC().UnixNano())

	httpServer:=s.ApiServerOptions.NewHttpServer()

	go controller.Run(wait.NeverStop)

	glog.V(0).Infoln("now start api server ......")
	if err:=httpServer.ListenAndServe();err!=nil{
		return err
	}
	return nil
}

func(s *ApiServerOptions) NewHttpServer() *http.Server{
	wsContainer:=restful.NewContainer()
	rest.InitApiserver(wsContainer)
	ListenningAddress:=s.ApiHost
	server:=&http.Server{Addr:ListenningAddress,Handler:wsContainer}
	return server
}



func(s *ApiServerOptions) InitEtcd(){
	etcd.Init([]string{s.EtcdEndPoint})
}


func(s *ApiServerOptions) InitStorage(){


}


