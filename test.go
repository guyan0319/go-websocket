package main

import "fmt"

func main() {
	fmt.Printf("%d\n", cal(1, 2))
	fmt.Printf("%d\n", cal(5, 2))
	fmt.Printf("%d\n", cal(5, 0))
	fmt.Printf("%d\n", cal(9, 2))
}

func cal(a, b int) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("%s\n", err)
		}
	}()
	return a / b
}