package queue

import (
	"sync"
	"bytes"
	"fmt"
	"github.com/jiangchengzi/mycloud/api"
)

type Node struct {
	Data api.Task
	Operation string
	next *Node
}

type Queue struct {
	head *Node
	tail *Node
	mu sync.RWMutex
}

type QueueImpl interface {
	Push(n *Node)
	Pop() *Node
	Del(n *Node)
	indexOf(n *Node)(index int)
	Update(n *Node)
}

//初始化为空
func NewQueue() *Queue{
	return &Queue{
		head:nil,
		tail:nil,
	}
}

func(q *Queue)Push(n *Node){
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.head == nil{
		q.tail = n
		q.head = n
		return
	}
	q.tail.next = n
	q.tail = n
}

func(q *Queue)Pop() *Node{
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.head == nil{
		return nil
	}
	if q.head.next == nil{//only one Node
		q.tail = q.tail.next
	}
	out := q.head
	q.head = q.head.next
	return out
}

func(q *Queue)Del(n *Node) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.head == nil{
		return
	}
	if q.head.Equal(n){
		if q.head.next == nil{ //only one Node
			q.head = nil
			q.tail = nil
			return
		}
		q.head = q.head.next
		return
	}
	item := q.head
	for item.next!=nil{
		if item.next.Equal(n) {
			if item.next.next == nil{
				q.tail = item
			}
			item.next = item.next.next
			return
		}
		item = item.next
	}
}

func(q *Queue) String() string{
	buf := bytes.Buffer{}
	item := q.head
	if item == nil{
		return "Queue{}"
	}
	buf.WriteString(fmt.Sprintf("Queue{%v",item.Data))
	for item.next != nil{
		item = item.next
		buf.WriteString(fmt.Sprintf(" %v",item.Data))
	}
	return buf.String()+"}"
}


func(q *Queue)indexOf(n *Node)(index int){
	index = 0
	if q.head.Equal(n){
		return index
	}
	item := q.head
	for item.next!=nil{
		index++
		if item.next.Equal(n) {
			return index
		}
		item = item.next
	}
	if !item.Equal(n){
		return -1
	}
	return index
}

func(n *Node)Equal(m *Node) bool{
	if (n == nil ) != (m == nil){
		return false
	}
	if m.Data.TaskID == n.Data.TaskID {
		return true
	}
	return false
}

func(q *Queue)Update(n *Node){
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.head == nil{
		return
	}
	if q.head.Equal(n){
		if q.head.next == nil{ //only one Node
			q.head = n
			q.tail = n
		}
		n.next = q.head.next
		q.head = n
		return
	}
	item := q.head
	for item.next!=nil{
		if item.next.Equal(n) {
			if item.next.next == nil{
				q.tail = n
			}
			n.next = item.next.next
			item.next = n
			return
		}
		item = item.next
	}
}

