package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type money struct {
	currency string
	value float64
}

type Invoice struct {
	Item string `json:"item"`
	Price money  `json:"price"`
	Discount money  `json:"discount"`
	Tax money  `json:"tax"`
}

/*
	input json format: 
	{
    "item": "Laptop",
    "price": "$1,299.99",
    "discount": "$150.00",
    "tax": "$112.49"
	}
*/

func (m *money) UnmarshalJSON(data []byte) (err error) {
	// first convert the data to string and strip it off its quotations
	v := string(data[1:len(data)-1])
	
	v = strings.ReplaceAll(v, ",", "")

	m.currency = string(v[0])
	m.value, err = strconv.ParseFloat(v[1:], 64)
	return
}

func (i *Invoice) Total() (float64, error) {
	if i.Price.currency == i.Discount.currency && i.Discount.currency == i.Tax.currency {
		return i.Price.value - i.Discount.value + i.Tax.value, nil
	}
	return 0.0, fmt.Errorf("currency mismatch: %s, %s, %s", i.Price.currency, i.Discount.currency, i.Tax.currency)
}

var input = `
	{
		"item": "Laptop",
    "price": "$1,299.99",
    "discount": "$150.00",
    "tax": "$112.49"
	}
`

func _11() {
	data := Invoice{}

	if err := json.Unmarshal([]byte(input), &data); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", data)
	cost, _ := data.Total()
	fmt.Printf("Total: %v\n", cost)
}