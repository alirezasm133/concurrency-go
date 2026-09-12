package main

import(
//	"sync"
	"fmt"
)


func main(){
//	var wg sync.WaitGroup
	ch:=make(chan int)
//	wg.Add(1)
	go func(){
		//defer wg.Done()
		ch<-42
	}()
	fmt.Println(<-ch)
	//wg.Wait()
	
	
}