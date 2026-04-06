package main

import "fmt"

func main() {
	fmt.Println(isMetagramm("text", "test")) // true
	fmt.Println(isMetagramm("кот", "кит"))   // true
	fmt.Println(isMetagramm("кот", "кот"))   // false
	fmt.Println(isMetagramm("кот", "коты"))  // false
	fmt.Println(isMetagramm("кот", "мит"))   // false
}

func isMetagramm(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	diff := 0

	for i := 0; i < len(s); i++ {
		if s[i] != t[i] {
			diff++
			if diff > 1 {
				return false
			}
		}
	}

	return diff == 1
}
