package main

import "fmt"

func getWeek() []string {
	week := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	week = append(week[6:], week[:6]...)
	return week
}

func main() {
	//fmt.Println(getWeek())
	week := []string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday",
	}
	// slice initialization
	weekSunday1 := append([]string{}, week[6])
	weekSunday1 = append(weekSunday1, week[0:6]...)
	fmt.Println(weekSunday1)
	// slice initialization (more memory-efficient)
	weekSunday2 := append(week[:0:0], week[6])
	weekSunday2 = append(weekSunday2, week[0:6]...)
	fmt.Println(weekSunday2)
	// slice with range notation
	week = append(week[6:], week[:6]...)
	fmt.Println(week)
}
