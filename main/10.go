package main

import (
	"fmt"
)

type Dog struct {}

func (d *Dog) Speak() string {
	return "woof"
}

func _10() {
	d1 := Dog{}
	d2 := &Dog{}
	fmt.Println(d1.Speak())
	fmt.Println(d2.Speak())

	m := make(map[string]*Dog)
	m["bob"] = &Dog{}
	m["bob"].Speak()
}