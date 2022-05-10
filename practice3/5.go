package practice3

import "fmt"

func NewMultiSet() map[string]int {
	return make(map[string]int)
}

func Insert(m map[string]int, val string) {
	m[val]++
}

func Erase(m map[string]int, val string) {
	m[val]--
}

func Count(m map[string]int, val string) int {
	return m[val]
}

func String(m map[string]int) (result string) {
	result += "{"

	for key, value := range m {
		for i := 0; i < value; i++ {
			result += fmt.Sprintf(" %s", key)
		}
	}

	result += " }"
	return
}
