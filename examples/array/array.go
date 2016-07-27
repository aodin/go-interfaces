package main

import "fmt"

// BEGIN OMIT
func Print(args ...interface{}) {
	fmt.Println(args...)
}

func main() {
	names := []string{"Alice", "Bob", "Carol"}
	Print(names...)
}
