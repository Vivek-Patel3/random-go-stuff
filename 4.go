package main

import "fmt"

func _4() {
	var s string
	fmt.Printf("s %T: %v\n", s, s)
	fmt.Printf("s %T: %q\n", s, s)

	s = "alice\n"
	fmt.Printf("s %T: %s\n", s, s)
	fmt.Printf("s %T: %q\n", s, s)
}
