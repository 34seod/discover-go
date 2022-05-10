package receiver

import "fmt"

type VertexId int

func (id VertexId) String() string {
	return fmt.Sprintf("VertexID(%d)", id)
}
