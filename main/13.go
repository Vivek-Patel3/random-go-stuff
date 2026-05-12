package main

import "fmt"

const (
	Red TrafficLight = iota
	Yellow TrafficLight = iota
	Green TrafficLight = iota
)

// now defining the type which will implement 
type TrafficLight int

func (t TrafficLight) String() string {
	switch t {
	case Red:
		return "Red"
	case Yellow:
		return "Yellow"
	case Green:
		return "Green"
	default:
		return "Invalid Color"
	}
}

func _13() {
	fmt.Println(Red)
	fmt.Println(Yellow)
	fmt.Println(Green)
	
}