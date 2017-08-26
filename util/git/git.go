package git

import (
	"errors"
	"fmt"
	"strings"
	"os/exec"
	"bytes"
)

type CloneOptions struct {
	//git clone 的地址
	Url string
	//git clone 下载到目标目录
	Directory string
	//认证信息
	Auth AuthOptions
	//分支信息
	Branch string
}

type AuthOptions struct {
	//username
	UserName string
	//token
	Token string
}
func (co *CloneOptions) GitClone() error{
	if err := co.isValid();err!=nil{
		return err
	}
	cmd := exec.Command("git","clone",co.Url,co.Directory)
	e := bytes.NewBuffer(nil)
	cmd.Stderr = e
	if err := cmd.Run();err != nil{
		stderror := string(e.Bytes())
		return errors.New(fmt.Sprintf("error occured in git clone,error message:'%s'",stderror))
	}
	return nil
}


func (co *CloneOptions) isValid() error{
	if co.Directory == "" {
		co.Directory = "."
	}
	if co.Url == "" {
		return errors.New("url is empty")
	}
	gitUrl := strings.Split(co.Url,"://")
	if gitUrl[0] != "https"{
		return errors.New(fmt.Sprintf("url is wrong,should be like `https://github.com/xxx.git`,but now is %s",co.Url))
	}
	co.Url =  fmt.Sprintf("%s://%s:%s@%s",gitUrl[0],co.Auth.UserName,co.Auth.Token,gitUrl[1])
	return nil
}