package customstring

import (
	"fmt"
	"time"
)

type Task struct {
	Title    string    `json:"title,omitempty"`
	Status   status    `json:"status,omitempty"`
	Daedline *Daedline `json:"daedline,omitempty"`
	Priority int       `json:"priority, omitempty"`
	SubTasks []Task    `json:"subTasks,omitempty"`
}

type status int

const (
	UNKNOWN status = iota
	TODO
	DONE
)

type Daedline struct {
	time.Time
}

func (t Task) String() string {
	check := "v"
	if t.Status != DONE {
		check = " "
	}
	return fmt.Sprintf("[%s] %s %s", check, t.Title, t.Daedline)
}

type IncludeSubTasks Task

func (t IncludeSubTasks) indentedString(prefix string) string {
	str := prefix + Task(t).String()
	for _, st := range t.SubTasks {
		str += "\n" + IncludeSubTasks(st).indentedString(prefix+"  ")
	}
	return str
}

func (t IncludeSubTasks) String() string {
	return t.indentedString("")
}
