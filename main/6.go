package main

import "fmt"

func _6() {
	num := []int{1,2,3}
	num = append(num, 4)
	fmt.Printf("Len: %v, Capacity: %v\n", len(num), cap(num))

	_ = append(num, 5)
	
	// the underlying array will have 1,2,3,4,5 but the len=4 because the new header which had the updated length was not captured
	fmt.Printf("num: %v, len: %v, cap: %v\n", num, len(num), cap(num))

	peeked := num[:5]
	fmt.Println(peeked)
}
