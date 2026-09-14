//excersize1
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	ch1 := make(chan int)
// 	ch2 := make(chan int)

// 	go func() {
// 		time.Sleep(1 * time.Second)
// 		ch1 <- 8
// 	}()

// 	go func() {
// 		time.Sleep(1 * time.Second)
// 		ch2 <- 20
// 	}()
// 	select {
// 	case value := <-ch1:
// 		fmt.Println("ch1:", value)

// 	case value := <-ch2:
// 		fmt.Println("ch2: ", value)
// 	}
// }

// excersize2
// package main

// import "fmt"

// func main() {
// 	ch := make(chan int)

// 	select {
// 	case value := <-ch:
// 		fmt.Println("value is ", value)
// 	default:
// 		fmt.Println("No value available")
// 	}
// }

// excersize3
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	ch := make(chan string)
// 	timeout := time.After(2 * time.Second)
// 	go func() {
// 		time.Sleep(3 * time.Second)
// 		ch <- "Operation completed"
// 	}()
// 	select {
// 	case value := <-ch:
// 		fmt.Println(value)

// 	case <-timeout:
// 		fmt.Println("Operation timed out")
// 	}

// }

//excersize4

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	ch1 := make(chan int)
// 	ch2 := make(chan int)
// 	ch3 := make(chan int)
// 	go func() {
// 		time.Sleep(1 * time.Second)
// 		ch1 <- 10
// 	}()
// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		ch2 <- 20
// 	}()
// 	go func() {
// 		time.Sleep(3 * time.Second)
// 		ch3 <- 30
// 	}()
// 	for j := 0; j < 3; j++ {
// 		select {
// 		case value := <-ch1:
// 			fmt.Println("channle one: ", value)
// 		case value := <-ch2:
// 			fmt.Println("channle two: ", value)
// 		case value := <-ch3:
// 			fmt.Println("channle three: ", value)
// 		}
// 	}
// }

//practice5

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	done := make(chan struct{})
// 	numbers := make(chan int)

// 	go func() {
// 		i := 1
// 		for {
// 			select {
// 			case <-done:
// 				fmt.Println("goroutine stopped")
// 				return
// 			case numbers <- i:
// 				i++
// 			}
// 		}

// 	}()

// 	for j := 0; j < 5; j++ {
// 		fmt.Println(<-numbers)
// 	}
// 	close(done)
// 	time.Sleep(100 * time.Millisecond)
// }

// practice6
// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(2)
// 	done := make(chan struct{})
// 	data := make(chan int)
// 	go func() {
// 		var i int
// 		defer wg.Done()
// 		for {
// 			select {
// 			case <-done:
// 				fmt.Println("producer stopped")
// 				return
// 			case data <- i:
// 				i++

// 			}
// 		}
// 	}()

// 	go func() {
// 		defer wg.Done()
// 		for {
// 			select {
// 			case <-done:
// 				fmt.Println("consumer stopped")
// 				return
// 			case value := <-data:
// 				fmt.Println(value)
// 			}
// 		}
// 	}()
// 	time.Sleep(1 * time.Second)
// 	close(done)
// 	wg.Wait()
// 	fmt.Println("main finished")
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	done := make(chan struct{})
// 	data := make(chan int, 3) // buffered

// 	// Producer
// 	go func() {
// 		i := 0

// 		for {
// 			select {
// 			case <-done:
// 				fmt.Println("producer stopped")
// 				return

// 			case data <- i:
// 				fmt.Println("produced:", i)
// 				i++
// 			}
// 		}
// 	}()

// 	// Consumer
// 	go func() {
// 		for i := 0; i < 3; i++ {
// 			value := <-data
// 			fmt.Println("consumed:", value)
// 		}

// 		fmt.Println("consumer stopped")
// 	}()

// 	time.Sleep(1 * time.Second)

// 	close(done)

// 	time.Sleep(1 * time.Second)

// 	fmt.Println("main finished")
// }

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(2)
// 	ch := make(chan int)

// 	go func() {
// 		defer wg.Done()

// 		for i := 0; i < 5; i++ {
// 			ch <- i
// 		}
// 		close(ch)
// 	}()
// 	//time.Sleep(1 * time.Second)
// 	//close(ch)
// 	go func() {
// 		defer wg.Done()
// 		for {
// 			select {
// 			case value, ok := <-ch:
// 				if ok == false {
// 					fmt.Println("channnel closed")
// 					return
// 				} else {
// 					fmt.Println(value)

// 				}
// 			}
// 		}
// 	}()

// 	wg.Wait()

// }
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	ch1 := make(chan int)
// 	ch2 := make(chan int)
// 	ch3 := make(chan int)
// 	var wg sync.WaitGroup
// 	wg.Add(4)
// 	go func() {
// 		defer wg.Done()
// 		ch1 <- 1
// 	}()
// 	go func() {
// 		defer wg.Done()
// 		ch2 <- 2
// 	}()
// 	go func() {
// 		defer wg.Done()
// 		ch3 <- 3
// 	}()

//		go func() {
//			defer wg.Done()
//			total := 0
//			for i := 0; i < 3; i++ {
//				select {
//				case value := <-ch1:
//					total += value
//				case value := <-ch2:
//					total += value
//				case value := <-ch3:
//					total += value
//				}
//			}
//			fmt.Println(total)
//		}()
//		wg.Wait()
//	}
package main

import "fmt"

func main() {
	highPriority := make(chan int)
	lowPriority := make(chan int)
	done := make(chan int)

	go func() {
		i := 0
		for {
			select {
			case <-done:
				fmt.Println("finished")
				return
			default:
				highPriority <- i
				i++
			}
			select {

			case highPriority <- i:
				i++
			case lowPriority <- i:
				i++

			}

		}
	}()

}
