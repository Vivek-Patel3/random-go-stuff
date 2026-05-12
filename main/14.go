package main

import "fmt"

type Byte int64

func convert(b int64) (float64, string) {
	var i int
	size := float64(b)

	for i=0;b/1024 != 0;i++ {
		size /= 1024
		b /= 1024
	}

	var unit string
	switch i {
	case 0:
		unit = "B"
	case 1:
		unit = "KB"
	case 2: 
		unit = "MB"
	case 3: 
		unit = "GB"
	case 4:
		unit = "TB"
	case 5:
		unit = "PB"
	case 6:
		unit = "EB"
	case 7:
		unit = "ZB"
	case 8:
		unit = "YB"
	}
	return size, unit
}

func (b Byte) String() string {
	val, unit := convert(int64(b))
	return fmt.Sprintf("%.2f %v", val, unit)
}

func _14() {
	fmt.Println(Byte(500))
	fmt.Println(Byte(2048))
	fmt.Println(Byte(1536870912))
}