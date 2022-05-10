package receiver

import "fmt"

func ExampleVertexID_print() {
	i := VertexId(100)
	fmt.Println(i)
	// Outout:
	// VertexID(100)
}
