//practice1
// package main

// import (
// 	"context"
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	i := 0
// 	ctx, cancel := context.WithCancel(context.Background())
// 	go func(ctx context.Context) {
// 		defer wg.Done()
// 		for {
// 			select {
// 			case <-ctx.Done():
// 				fmt.Println("stopped")
// 				return
// 			default:
// 				time.Sleep(500 * time.Millisecond)
// 				fmt.Println(i)
// 				i++

// 			}
// 		}

//		}(ctx)
//		time.Sleep(3 * time.Second)
//		cancel()
//		wg.Wait()
//	}
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func doWork(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("stopped,Timeout")
		//ctx.Err()
	case <-time.After(3 * time.Second):
		fmt.Println("work finished")
	}
}
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	go func() {
		defer wg.Done()
		doWork(ctx)

	}()
	defer cancel()
	wg.Wait()
}
