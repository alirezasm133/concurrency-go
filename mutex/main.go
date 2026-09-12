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

package main
import (
	"fmt"
	"sync"
)

type RateLimiter struct{
	mu sync.Mutex
	count int
}

func (r *RateLimiter)Allow()bool{
	defer r.mu.Unlock()
	r.mu.Lock()
	if r.count < 5{
		r.count++
		return true

	}else{
		fmt.Println("request is more than 5/persecond")
		return false
	}

}

func main(){
	var wg sync.WaitGroup
	wg.Add(10)
	r:=RateLimiter{count: 0,}

	go func(){
		defer wg.Done()
		b:=r.Allow()
		if b== true{
			fmt.Println("accepted")
			return
		}else{
			fmt.Println("access denied")
			return
		}
	}()

	go func(){
		defer wg.Done()
		b:=r.Allow()
		if b== true{
			fmt.Println("accepted")
			return
		}else{
			fmt.Println("access denied")
			return
		}
	}()
	
	go func(){
		defer wg.Done()
		b:=r.Allow()
		if b== true{
			fmt.Println("accepted")
			return
		}else{
			fmt.Println("access denied")
			return
		}
	}()

	go func(){
		defer wg.Done()
		b:=r.Allow()
		if b== true{
			fmt.Println("accepted")
			return
		}else{
			fmt.Println("access denied")
			return
		}
	}()

	go func(){
		defer wg.Done()
		b:=r.Allow()
		if b== true{
			fmt.Println("accepted")
			return
		}else{
			fmt.Println("access denied")
			return
		}
	}()

	go func(){
		defer wg.Done()
		b:=r.Allow()
		if b== true{
			fmt.Println("accepted")
			return
		}else{
			fmt.Println("access denied")
			return
		}
	}()

	go func(){
		defer wg.Done()
		b:=r.Allow()
		if b== true{
			fmt.Println("accepted")
			return
		}else{
			fmt.Println("access denied")
			return
		}
	}()

	go func(){
		defer wg.Done()
		b:=r.Allow()
		if b== true{
			fmt.Println("accepted")
			return
		}else{
			fmt.Println("access denied")
			return
		}
	}()	


	go func(){
		defer wg.Done()
		b:=r.Allow()
		if b== true{
			fmt.Println("accepted")
			return
		}else{
			fmt.Println("access denied")
			return
		}
	}()


	go func(){
		defer wg.Done()
		b:=r.Allow()
		if b== true{
			fmt.Println("accepted")
			return
		}else{
			fmt.Println("access denied")
			return
		}
	}()
	wg.Wait()
}

//end of program

