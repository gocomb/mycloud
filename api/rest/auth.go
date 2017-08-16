package rest

import (
	"github.com/emicklei/go-restful"
	"github.com/golang/glog"
	"fmt"
)

func Checktoken(req *restful.Request, resp *restful.Response, chain *restful.FilterChain){
	glog.Infof("[check-token-filter (logger)] %s,%s\n", req.Request.Method, req.Request.URL)
	fmt.Sprintf("sxsxs")
	chain.ProcessFilter(req, resp)
}
