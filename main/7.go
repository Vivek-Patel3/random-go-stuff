package main

import (
	"fmt"
	"strings"
)

type user struct {
	name string
	age int
	location string
}

func createUser(name string, args ...interface{}) user {
	u := user{name: name, age: 18, location: "New York"}

	if len(args) == 0 {
		return u
	}

	for i:=0;i<len(args);i+=2 {
		key := args[i].(string)
		val := args[i+1]

		switch strings.ToLower(key) {
		case "age":
			u.age = val.(int)
		case "location":
			u.location = val.(string)
		}
	}
	return u
}

func _7() {
	user1 := createUser("Steve Harrington")
	user2 := createUser("Robin Buckley", "age", 18, "location", "Hawkins")

	fmt.Printf("%+v\n", user1)
	fmt.Printf("%+v\n", user2)
}
