package func_practice

import "fmt"

func ExampleNewIntGenerator() {
	gen := NewIntGenerator()
	fmt.Println(gen(), gen(), gen(), gen(), gen())
	fmt.Println(gen(), gen(), gen(), gen(), gen())
	gen2 := NewIntGenerator()
	fmt.Println(gen2())
	// Output:
	// 1 2 3 4 5
	// 6 7 8 9 10
	// 1
}
