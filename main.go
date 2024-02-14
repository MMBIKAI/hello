package main

import "fmt"

func hi(name string) string {

	return "Hello, " + name
}

func main() {

	fmt.Println(hi("world"))
}
