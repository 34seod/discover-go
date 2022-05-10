package practice3

import (
	"fmt"

	"github.com/34seod/discover-go/hangul"
)

func Example_array() {
	fruits := [3]string{"과일", "바나나", "토마토"}
	for _, fruit := range fruits {
		if hangul.HasConsonantSuffix(fruit) {
			fmt.Printf("%s은 맛있다.\n", fruit)
		} else {
			fmt.Printf("%s는 맛있다.\n", fruit)
		}
	}
	// Output:
	// 과일은 맛있다.
	// 바나나는 맛있다.
	// 토마토는 맛있다.
}
