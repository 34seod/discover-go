package func_practice

import "fmt"

func ExampleBinOnToBinSub() {
	sub := BinOnToBinSub(func(a, b int) int {
		return a + b
	})
	sub(5, 7)
	sub(5, 7)
	count := sub(5, 7)
	fmt.Println("count:", count)
	// Output:
	// 12
	// 12
	// 12
	// count: 3
}
