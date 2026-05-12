package main

import "fmt"

func _3() {
	type Person struct { Name string }
	p := Person{Name: "Alice"}
	fmt.Printf("%v\n", p)  // Output: {Alice}
	fmt.Printf("%+v\n", p) // Output: {Name:Alice}
	fmt.Printf("%#v\n", p) // Output: main.Person{Name:"Alice"}
}
