package task_Builder

import (
	"github.com/spf13/pflag"
	"runtime"
	"math/rand"
	"time"
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

type TaskBuilderOption struct {
	// 设定api-server的地址
	ApiServer string

	// ID 用于标示task_builder,在访问api-server时，
	// 利用加密算法+ID生成token，以便api-server进行鉴别，
	// 后期可以改为使用证书，类似于k8s中的serviceaccount
	ID string
}

func NewTaskBuilderOption() *TaskBuilderOption{
	return &TaskBuilderOption{}
}

func (tbo *TaskBuilderOption)AddFlags(fs *pflag.FlagSet){
	fs.StringVar(&tbo.ApiServer,"api-server","127.0.0.1:28000","set api-server host")
	fs.StringVar(&tbo.ID,"id","","task build id")
	return
}


func (tbo *TaskBuilderOption)Run() error{
	runtime.GOMAXPROCS(runtime.NumCPU())

	rand.Seed(time.Now().UTC().UnixNano())

	return nil
}


func GenToken(ID string,Url string) (token string){
	id,_ := strconv.Atoi(ID)
	url,_ := strconv.Atoi(Url)
	// 将id与url相与，然后再经过md5编码，api-server在接收到请求后会首先对token进行比对(也进行一次这样的操作)，
	// 如果token不相同，则记录操作审计中，返回badrequest，这样就可以避免信息在传输过程中被人恶意串改
	input := id&url
	h := md5.New()
	h.Write([]byte(strconv.Itoa(input)))
	token = hex.EncodeToString(h.Sum(nil))
	return
}


