package main

import (
	"errors"
	"fmt"
)

var DivideByZero = errors.New("division by zero")

func _8(a, b int) (result int, err error) {
	if b == 0 {
		err = fmt.Errorf("operation failed: %w", DivideByZero)
		return
	}
	result = a/b
	return
}
