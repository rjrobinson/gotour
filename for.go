package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		fmt.Println("Hello! The next number is....")
		sum += sum
		fmt.Println(sum)

	}
	fmt.Println(sum)
}
