package main

import "fmt"

func Hello() string {
	return "Hello world"
}

func main() {
	result := Hello()
	fmt.Println(result)
}
