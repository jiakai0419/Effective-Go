package main

import "fmt"

func shouldEscape(c byte) bool {
	switch c {
	case ' ', '?', '&', '=', '#', '+', '%':
		return true
	}
	return false
}

func main() {
	fmt.Println(shouldEscape(' '))
	fmt.Println(shouldEscape('='))
	fmt.Println(shouldEscape('+'))
	fmt.Println(shouldEscape('a'))
}
