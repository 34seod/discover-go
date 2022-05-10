package practice3

import "fmt"

func ExampleQueue() {
	q := Queue{q: []int{1, 2, 3}}
	fmt.Println(q.q)
	q.Add(4)
	fmt.Println(q.q)
	fmt.Println(q.Pop())
	fmt.Println(q.q)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.q)

	// Output:
	// [1 2 3]
	// [1 2 3 4]
	// 1
	// [2 3 4]
	// 2
	// 3
	// 4
	// 0
	// 0
	// []
}
