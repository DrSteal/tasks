package main

import "fmt"

func smallestNum(arr []int) []int {
	a := arr[0]
	chArr := []int{}
	// for i := 0; i <= len(arr); {
	for _, r := range arr {
		if r < a {
			chArr = append(chArr, r)

		} else if len(chArr) == len(arr) {
			arr[0] = 0
		}
	}
	// }
	// fmt.Println(arr)
	return chArr
}

func main() {
	arr := []int{14, 7, 9, 8, 17, 5}
	fmt.Println(smallestNum(arr))

}
