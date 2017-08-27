package logger

import (
	"flag"
	"log"
	"github.com/golang/glog"
	"github.com/jiangchengzi/mycloud/util/wait"
	"time"
)

func init() {
	flag.Set("logtostderr", "true")
}

type GlogWriter struct{}

func (writer GlogWriter) Write(data []byte) (n int, err error) {
	glog.Info(string(data))
	return len(data), nil
}

func InitLogs() {
	log.SetPrefix("mycloud")
	log.SetOutput(GlogWriter{})
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)
	// The default glog flush interval is 30 seconds, which is frighteningly long.
	go wait.Until(glog.Flush,5*time.Second, wait.NeverStop)
}
