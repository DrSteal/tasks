package proj

import "fmt"

func LetsGo(a int) {
	a++
	fmt.Print(a)
}

// func main() {
// 	// func binary_search(list []int, item int) {

// 	// 	low := 0
// 	// 	high := len(list) - 1
// 	// 	count := 0
// 	// 	for low <= high {
// 	// 		mid := (low + high) / 2
// 	// 		guess := list[mid]
// 	// 		if guess == item {
// 	// 			fmt.Println(guess)
// 	// 			fmt.Println(mid)
// 	// 			fmt.Println(count)
// 	// 			break
// 	// 		} else if guess > item {
// 	// 			high = mid - 1
// 	// 			count++
// 	// 		} else if guess < item {
// 	// 			low = mid + 1
// 	// 			count++
// 	// 		} else {
// 	// 			return
// 	// 		}
// 	// 	}
// 	// }

// }
