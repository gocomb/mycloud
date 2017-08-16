package queue

import (
	"testing"
	"github.com/jiangchengzi/mycloud/api"
)

func TestQueue(t *testing.T){
	q := NewQueue()
	q.Push(&Node{
		Data:api.Task{
			TaskID:"1",
			TaskName:"TaskName1",
		},
		Operation:"create",
	})
	q.Push(&Node{
		Data:api.Task{
			TaskID:"2",
			TaskName:"TaskName2",
		},
		Operation:"create",
	})
	q.Push(&Node{
		Data:api.Task{
			TaskID:"3",
			TaskName:"TaskName3",
		},
		Operation:"create",
	})
	t.Log(q)

	q.Update(&Node{
		Data:api.Task{
			TaskID:"1",
			TaskName:"TaskName3",
		},
		Operation:"update",
	})
	t.Log(q)

	q1 := q.Pop()
	t.Log(q1.Data)

	t.Log(q)
	q.Del(&Node{
		Data:api.Task{
			TaskID:"3",
			TaskName:"TaskName3",
		},
		Operation:"delete",
	})
	t.Log(q)

}
