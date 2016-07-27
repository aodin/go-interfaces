package main

import "fmt"

// BEGIN OMIT
type Getter interface {
	Get() int
}

type Thing struct{}

func (thing Thing) Get() int { return 42 }

func main() {
	var thing *Thing

	func(getter Getter) {
		fmt.Println(getter.Get())
	}(thing)
}
