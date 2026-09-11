// //exercise 1

// package main

// import(
// 	"fmt"
// 	"sync"
// )
// type My struct{
// 	test  map[string]int
// 	mu sync.RWMutex
// }
// func (my *My)Set(key string,value int){
	
// 	my.mu.Lock()
// 	my.test[key]=value
// 	my.mu.Unlock()
// }
// func (my *My)Get(key string)(int,bool){
// 	my.mu.RLock()
// 	value,ok:= my.test[key]
// 	defer my.mu.RUnlock()
// 	return value,ok
	
// }
// func main(){
// 	var wg sync.WaitGroup
// 	wg.Add(7)

// 	x:= My{test: map[string]int {"ali":27}}
// 	//fmt.Println(x)
// 	go func(){
// 		defer wg.Done()
// 		x.Set("naser",13)
// 	}()
// 	go func(){
// 		defer wg.Done()
// 		x.Set("salem",16)
// 	}()
// 	go func(){
// 		defer wg.Done()
// 		x.Set("mmahmoud",15)
// 	}()
// 	go func(){
// 		defer wg.Done()
// 		x.Set("kamran",18)
// 	}()
// 	go func(){
// 		defer wg.Done()
// 		x.Set("babak",19)
// 	}()
// 	go func(){
// 		defer wg.Done()
// 		value,ok:=x.Get("ali")
// 		fmt.Println(value,ok)
// 	}()
// 	go func(){
// 		defer wg.Done()
// 		value,ok:=x.Get("saleh")
// 		fmt.Println(value,ok)
// 	}()
// 	wg.Wait()
// 	fmt.Println(x.test)
// }
func main(){

	
}

//end of program

