package main

import (
	"fmt"
)

type Getter interface {
	Get() int
	N() int
}

type Outer struct {
	Getter
}

var _ Getter = Outer{}

func (outer Outer) Get() int { return 4 * outer.N() }

type Inner struct{}

var _ Getter = Inner{}

func (inner Inner) Get() int { return 2 * inner.N() }

func (inner Inner) N() int { return 1 }

func Display(getter Getter) {
	fmt.Println(getter.Get())
}

func main() {
	i := Inner{}
	o := Outer{Getter: i}
	fmt.Println(o.Get())
	Display(o)
}
