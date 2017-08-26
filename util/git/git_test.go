package git

import "testing"

func TestClone(t *testing.T){
	c := CloneOptions{
		Url:"https://github.com/jiangchengzi/istio_cn.git",
		Branch:"dev",
		Directory:"./m2",
		Auth:AuthOptions{
			UserName:"jiangchengzi",
			Token:"1d4e8062cefad249f71a1806b4bb9c311a9ff164",
		},
	}
	if err := c.GitClone();err!=nil{
		t.Log(err)
	}
	t.Logf("success")
}
