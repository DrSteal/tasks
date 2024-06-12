package main

import (
	"fmt"
	"strings"
)

func repeatString(str string) string {
	rep := []rune(str)
	d := str
	var k string
	k = string(d[0])

	str = strings.Replace(str, k, "", -1)
	for i := 0; i < len(rep); i ++ {
			if k = str[i] {
				// проверяй первый символ. Если он повторяется удаляй
			}
	}
	// for _, r := range str {

	// 	if d == r {
	// 		strings.Trim(str, string(r))
	// 	}
	// }
	// str = string(rep)
	return string(str)
}

func main() {
	k := "zaabcbd"
	fmt.Println(repeatString(k))
}

// func smallestNum(arr []int) []int {
// 	a := arr[0]
// 	chArr := []int{}
// 	// for i := 0; i <= len(arr); {
// 	for _, r := range arr {
// 		if r < a {
// 			chArr = append(chArr, r)

// 		} else if len(chArr) == len(arr) {
// 			arr[0] = 0
// 		}
// 	}
// 	// }
// 	// fmt.Println(arr)
// 	return chArr
// }

// func evenSum(from, to int, ch chan int) {
// 	result := 0
// 	for i := from; i <= to; i++ {
// 		if i%2 == 0 {
// 			result += i
// 		}
// 	}
// 	ch <- result
// }
// func squareSum(from, to int, ch chan int) {
// 	result := 0
// 	for i := from; i <= to; i++ {
// 		if i%2 == 0 {
// 			result += i * i
// 		}
// 	}
// 	ch <- result
// }

// func main() {
// 	evenCh := make(chan int)
// 	sqCh := make(chan int)

// 	go evenSum(0, 100, evenCh)
// 	go squareSum(0, 100, sqCh)

// 	fmt.Println(<-evenCh + <-sqCh)
// }

// func doTask(n int, ch chan int, wg *sync.WaitGroup) {
// 	time.Sleep(time.Second)
// 	for i := 0; i < n; i++ {
// 		fmt.Scan(&ch)
// 		t := <-ch
// 		t *= t
// 		ch <- t
// 	}
// 	defer wg.Done()
// }

// func doTask(n int, ch chan int, wg *sync.WaitGroup) {
// 	// time.Sleep(time.Second)
// 	for i := 0; i < n; i++ {
// 		var t int
// 		fmt.Scan(&t)
// 		t
// 		ch <- t
// 	}

// 	defer wg.Done()
// }

// func main() {
// 	var n int

// 	fmt.Scan(&n)
// 	tasks := make(chan int)
// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	go doTask(n, tasks, &wg)
// 	<-tasks
// 	wg.Wait()

// 	res := <-tasks
// 	fmt.Println(res)

// }

// outCh := make(chan int)

// var wg sync.WaitGroup

// for i := 0; i < k; i++ {
// 	wg.Add(1)
// 	go func() {
// 		res := <-tasks
// 		outCh <- res
// 		wg.Done()
// 	}()
// }
// wg.Wait()

// output := make([]int, 0, len(tasks))
// for t := range outCh {
// 	output = append(output, t)
// }
// close(outCh)

// for _, v := range output {
// 	fmt.Print(v, " ")
// }

// func generator(num int, ch chan int, wg *sync.WaitGroup) {

// 	for i := 1; i <= num; i++ {
// 		ch <- i

// 	}

// 	wg.Done()

// }

// func main() {
// 	var n int

// 	ch := make(chan int)

// 	var wg sync.WaitGroup

// 	wg.Add(1)

// 	fmt.Scan(&n)

// 	go generator(n, ch, &wg)

// 	wg.Wait()

// 	b := <-ch

// 	fmt.Println(b)

// }

// func generator(n int) chan int {
// 	ch := make(chan int)
// 	ch <- n
// 	return ch
// }

// func main() {
// 	var n int
// 	fmt.Scan(&n)

// 	for i := 0; i < n; i++ {
// 		gen := generator(i)
// 		fmt.Println(gen)

// 	}
// }
