package main

import (
	"fmt"
	"os"
)

var users = map[string]string{
	"305": "Sue",
	"204": "Bob",
	"631": "Jake",
	"073": "Tracy",
}

func deleteUser(id string) {
	delete(users, id)
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Used ID not passed")
		os.Exit(1)
	}
	users := map[string]string{
		"305": "Sue",
		"204": "Bob",
		"631": "Jake",
		"073": "Tracy",
	}
	delete(users, os.Args[1])

	if len(os.Args) < 2 {
		fmt.Println("User ID not passed")
		os.Exit(1)
	}
	userID := os.Args[1]
	fmt.Println(userID)
	deleteUser(userID)
	fmt.Println("Users:", users)
}
