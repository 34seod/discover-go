package practice3

import "fmt"

func ExampleSort() {
	list := []int{10, 6, 1, 5, 3}
	Sort(&list)
	fmt.Println(list)
	// Output:
	// [1 3 5 6 10]
}
