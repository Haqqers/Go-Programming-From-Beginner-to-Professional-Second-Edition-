package main

import "fmt"

func main() {
	nums := []int{10, 112, 2, 3, 5, 8, 0, 9, 7, 6, 9, 1, 10, 11}
	fmt.Printf("Before : %v\n", nums)

	for i := 0; i < len(nums); i++ {
		num1 := nums[i]
		for j := 0; j < len(nums); j++ {
			num2 := nums[j]
			if num2 > num1 {
				nums[i], nums[j] = nums[j], nums[i]
				fmt.Printf("swapped %v and %v\n", nums[i], nums[j])
				i--
				break
			}
		}
	}
	fmt.Printf("After : %v\n", nums)
}

// func main() {
// 	nums := []int{5, 8, 2, 4, 0, 1, 3, 7, 9, 6}
// 	fmt.Println("Before:", nums)
// 	for swapped := true; swapped; {
// 		swapped = false
// 		for i := 1; i < len(nums); i++ {
// 			if nums[i-1] > nums[i] {
// 				nums[i], nums[i-1] = nums[i-1], nums[i]
// 				swapped = true
// 			}
// 		}
// 	}
// 	fmt.Println("After :", nums)
// }
