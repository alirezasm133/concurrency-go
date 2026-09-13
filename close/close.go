// package main

// import(
// 	"fmt"
// 	//"sync"
// )


// func main(){
// 	ch :=make(chan int,3)
// 	ch<-10
// 	//ch<-20
// 	//ch<-30
// 	//close(ch)
// 	//ch<-30
// 	value,ok:=<-ch
// 	fmt.Println(value,ok)
// 	value1,ok1:=<-ch
// 	fmt.Println(value1,ok1)
// 	//fmt.Println(<-ch)
// 	//fmt.Println(<-ch)
// }

// package main 

// import(
// 	"fmt"
// )

// func main(){
// 	ch:=make(chan int)
// 	go func(){
// 		defer close(ch)
// 		for i:=1;i<=5;i++{
// 			ch<-i
// 		}

// 	}()

// 	for value:=range ch {
// 		fmt.Println(value)
// 	}
// }


package main

import(
	"fmt"
)

func main(){
	ch:=make(chan int)
	go func(){
		defer close(ch)
		for i:=1;i<=5;i++{
			ch<-i
		}
	}()
	for value:=range ch{
		fmt.Println(value)
	}
}