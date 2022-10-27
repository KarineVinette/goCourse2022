package main

import "fmt"

func main() {
	greetings := IntMin(12, 15)
	fmt.Println(greetings)
}
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

