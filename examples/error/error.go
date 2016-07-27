package main

import "fmt"

// BEGIN OMIT
type Error struct{ Code int }

func (err Error) Error() string { return "error!" }

func NoError() *Error { return nil }

func WrapError(err error) error { return err }

func main() {
	err := NoError()
	fmt.Println("NoError():", err == nil)
	fmt.Println("WrapError():", WrapError(err) == nil)
}
