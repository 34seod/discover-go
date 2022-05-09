package Map

import "fmt"

func ExampleHasDupeRune() {
	fmt.Println(hasDupeRune("숨바꼭질"))
	fmt.Println(hasDupeRune("다시합시다"))
	// Output:
	// false
	// true
}

func ExampleHasDupeRune2() {
	fmt.Println(hasDupeRune2("숨바꼭질"))
	fmt.Println(hasDupeRune2("다시합시다"))
	// Output:
	// false
	// true
}
