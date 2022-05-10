package practice3

import "fmt"

func ExampleBinarySearch() {
	target1 := "ghi"
	target2 := "abc"
	target3 := "zzzz"
	target4 := "jkl"
	src := []string{
		"abc",
		"def",
		"ghi",
		"jkl",
	}
	fmt.Println(BinarySearch(src, target1))
	fmt.Println(BinarySearch(src, target2))
	fmt.Println(BinarySearch(src, target3))
	fmt.Println(BinarySearch(src, target4))
	// Output:
	// true
	// true
	// false
	// true
}
