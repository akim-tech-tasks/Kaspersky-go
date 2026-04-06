package main

import "fmt"

func main() {
	fmt.Println(isMetagrammN("text", "test", 1)) // true
	fmt.Println(isMetagrammN("кот", "кит", 1))   // true
	fmt.Println(isMetagrammN("кот", "мит", 2))   // true
	fmt.Println(isMetagrammN("кот", "кот", 0))   // true
	fmt.Println(isMetagrammN("кот", "кот", 1))   // false
}

func isMetagrammN(s, t string, n int) bool {
	rs := []rune(s)
	rt := []rune(t)

	if len(rs) != len(rt) {
		return false
	}

	diff := 0

	for i := 0; i < len(rs); i++ {
		if rs[i] != rt[i] {
			diff++
			if diff > n {
				return false
			}
		}
	}

	return diff == n
}
