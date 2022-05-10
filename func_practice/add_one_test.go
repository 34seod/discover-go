package func_practice

import "fmt"

func ExampleAddOne() {
	n := []int{1, 2, 3, 4}
	AddOne(n)
	fmt.Println(n)
	// Output:
	// [2 3 4 5]
}

func Example_funcLiteralVar() {
	printHello := func() {
		fmt.Println("Hello!")
	}
	printHello()
	// Output:
	// Hello!
}
