//practice 1

// package main

// import(
// //	"sync"
// 	"fmt"
// )


// func main(){
// //	var wg sync.WaitGroup
// 	ch:=make(chan int)
// //	wg.Add(1)
// 	go func(){
// 		//defer wg.Done()
// 		ch<-42
// 	}()
// 	fmt.Println(<-ch)
// 	//wg.Wait()
	
	
// }

//practice2
// package main
// import("fmt")

// func main(){
// 	ch:=make(chan string)
// 	go func(){
// 		ch<-"Go"
// 		ch<-"is"
// 		ch<-"awesome"
// 		close(ch)
// 	}()
// 	//fmt.Println(<-ch)
// 	// fmt.Println(<-ch)
// 	// fmt.Println(<-ch)
// 	for value:=range ch{
// 		fmt.Println(value)
// 	}
// }

//practice 3
package main

import(
	"fmt"
	"time"
)
func main(){
	ch:=make(chan string)
	go func(){
		time.Sleep(2* time.Second)
		ch<-"Hello"
	}()
	fmt.Println(<-ch)

}