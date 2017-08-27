package id

import (
	"github.com/renstrom/shortuuid"
)

const (
	TaskID = `TID-`

	TaskFlowID = `TFID-`
)

const IDSHORTLENGTH = 10

func NewTaskID() string {
	return NewShortID(TaskID)
}

func NewTaskFlowID() string {
	return NewShortID(TaskFlowID)
}

func NewShortID(prefix string) string {
	id := shortuuid.New()
	return prefix + id[:IDSHORTLENGTH]
}
