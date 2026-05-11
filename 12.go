package main

import "fmt"

/*
Interfaces in golang provide "dynamic" feel.
Golang is statically typed language so the types of variables is determined at runtime
But for interfaces, only the method set for the interface is determined at compile time, the actual concrete type held by the interface is determined at runtime
Thus which method to call is determined at runtime (because its type is determined at runtime), giving golang polymorphism.
*/

type Person struct {
	Name string
	Age int
}

func (p Person) Describe() {
	fmt.Printf("%v [age: %v]\n", p.Name, p.Age)
}

func _12() {
	var i interface{}
	i = Person{
		Name: "Jane Hopper",
		Age: 15,
	}

	// at compile time, i will be decoded as interface. Therefore if we want to use fields or method of the underlying concrete value that it is holding, we need to extract the underlying value. The reason being, for interfaces, at compile time only the method set of the interface is fixed - the method set of the underlying value and its fields are not fixed at compile time
	// fmt.Println(i.Name) -> compile time error because for compiler i is an interface and therefore doesn't have any fields
	// i.Describe() -> compile time error because for compiler i is an interface with no methods (empty interface)

	fmt.Println(i.(Person).Name)
	fmt.Println(i.(Person).Age)

	i.(Person).Describe()
}
