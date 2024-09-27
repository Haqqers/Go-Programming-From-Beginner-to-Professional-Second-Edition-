package main

import "fmt"

var slice = []string{
	"Good",
	"Good",
	"Bad",
	"Good",
	"Good",
}

func removeBad() []string {
	sli := []string{"Good", "Good", "Bad", "Good", "Good"}
	sli = append(sli[:2], sli[3:]...)
	return sli
}

func main() {
	//fmt.Println(removeBad())
	slice = append(slice[:2], slice[3:]...)
	fmt.Println(slice)
}
