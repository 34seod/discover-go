package enum

import (
	"errors"
	"fmt"
	"time"
)

type Daedline struct {
	time.Time
}
type Task struct {
	Title    string    `json:"title"`
	Status   status    `json:"status"`
	Daedline *Daedline `json:"daedline" time_format:"sql_date" time_utc:"true"`
}

func (s status) MarshalJSON() ([]byte, error) {
	switch s {
	case UNKNOWN:
		return []byte(`"UNKNOWN"`), nil
	case TODO:
		return []byte(`"TODO"`), nil
	case DONE:
		return []byte(`"DONE"`), nil
	default:
		return nil, errors.New("status.MarshalJSON: unknown value")
	}
}

func (s *status) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case `"UNKNOWN"`:
		*s = UNKNOWN
	case `"TODO"`:
		*s = TODO
	case `"DONE"`:
		*s = DONE
	default:
		return errors.New("status.UnmarshalJSON: unknown value")
	}
	return nil
}
func NewDaedline(t time.Time) *Daedline {
	return &Daedline{t}
}

type status int

const (
	UNKNOWN status = iota
	TODO
	DONE
)

func aaa() {
	t := Task{Title: "aaa", Status: TODO}
	fmt.Println(t)
}
