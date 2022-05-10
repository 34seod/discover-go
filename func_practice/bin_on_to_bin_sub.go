package func_practice

import "fmt"

type BinOp func(int, int) int
type BinSub func(int, int) int

func BinOnToBinSub(f BinOp) BinSub {
	var count int
	return func(a, b int) int {
		fmt.Println(f(a, b))
		count++
		return count
	}
}
