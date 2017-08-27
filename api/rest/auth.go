package rest

import (
	"github.com/emicklei/go-restful"
)

func PrePare(req *restful.Request, resp *restful.Response, chain *restful.FilterChain){
	req.SetAttribute("UserID","user-1")
	req.SetAttribute("Namespace","namespace1")
	chain.ProcessFilter(req, resp)
}


func Checktoken(req *restful.Request, resp *restful.Response, chain *restful.FilterChain){
	chain.ProcessFilter(req, resp)
}
