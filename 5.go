package main

import "fmt"

const (
	years = 10
	leapYears = int32(20)
)

func _5() {
	result := years * leapYears
	fmt.Printf("Result %T: %v\n", result, result)
}