package main

import (
	"github.com/spf13/pflag"
)

type ApiServerOptions struct {
	ApiHost  string `json:"apiserver_host,omitempty"`
	EtcdEndPoint string `json:"etcd_endpoint,omitempty"`
}

type ApiServer struct {
	ApiServerOptions *ApiServerOptions
}


func NewApiServer(opts *ApiServerOptions) *ApiServer {
	return &ApiServer{
		ApiServerOptions:opts,
	}
}

func NewApiOptions() *ApiServerOptions{
	return &ApiServerOptions{}
}


func (a *ApiServerOptions) AddFlags(fs *pflag.FlagSet){
	fs.StringVar(&a.ApiHost,"host","127.0.0.1:28000","set api-server host")
	fs.StringVar(&a.EtcdEndPoint,"etcd-endpoint","http://127.0.0.1:2379","set etcd endpoint")
	return
}







