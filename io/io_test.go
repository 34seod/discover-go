package io

import (
	"fmt"
	"os"
	"strings"
)

func ExampleWriteTo() {
	lines := []string{
		"bil@gmail.com",
		"tom@gmail.com",
		"jane@gmail.com",
	}
	if err := WriteTo(os.Stdout, lines); err != nil {
		fmt.Println(err)
	}
	// Output:
	// bil@gmail.com
	// tom@gmail.com
	// jane@gmail.com
}

func ExampleReadFrom() {
	r := strings.NewReader("bill\ntom\njane\n")
	var lines []string
	if err := ReadFrom(r, &lines); err != nil {
		fmt.Println(err)
	}
	fmt.Println(lines)
	// Output:
	// [bill tom jane]
}
