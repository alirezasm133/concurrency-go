//practice1
//package main

// import "fmt"

// func sendNumber(ch chan<- int) {
// 	for i := 0; i < 5; i++ {
// 		ch <- i
// 	}
// 	close(ch)
// }
// func main() {
// 	ch := make(chan int)
// 	go sendNumber(ch)
// 	for value := range ch {
// 		fmt.Println(value)
// 	}

// }

//practice2
// package main

// import "fmt"

// func printNumbers(ch <-chan int) {
// 	for value := range ch {
// 		fmt.Println(value)
// 	}
// }
// func main() {
// 	ch := make(chan int)
// 	go printNumbers(ch)
// 	for i := 0; i < 5; i++ {
// 		ch <- i
// 	}
// 	close(ch)

// }

//practice5
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func producer(ch chan<- int) {
// 	for i := 1; i <= 10; i++ {
// 		ch <- i
// 	}
// }
// func consumer(ch <-chan int) {
// 	value := 0
// 	for i := 0; i < 10; i++ {
// 		value += <-ch
// 	}
// 	fmt.Println(value)
// }
// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(2)

// 	ch := make(chan int)
// 	go func() {
// 		defer wg.Done()
// 		producer(ch)
// 	}()
// 	go func() {
// 		defer wg.Done()
// 		consumer(ch)
// 	}()
// 	wg.Wait()

// }

// practice5
package main

import (
	"fmt"
	"sync"
)

func producerHigh(ch chan<- int) {
	for i := 1; i <= 3; i++ {
		ch <- i * 100
	}
}
func producerLow(ch chan<- int) {
	for i := 1; i <= 3; i++ {
		ch <- i
	}
}
func consumer(high <-chan int, low <-chan int) {
	for i := 0; i < 6; i++ {
		select {
		case value := <-high:
			fmt.Println(value)
		case value := <-low:
			fmt.Println(value)
		}
	}

}
func main() {
	high := make(chan int)
	low := make(chan int)
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		producerHigh(high)
	}()
	go func() {
		defer wg.Done()
		producerLow(low)
	}()
	go func() {
		defer wg.Done()
		consumer(high, low)
	}()
	wg.Wait()
}
