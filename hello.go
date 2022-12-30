package main

import "fmt"

func test(test string) {
	fmt.Println(test)
}

func main() {
	var testvar = "test"

	fmt.Println("Hello, World!")
	test(testvar)
}
