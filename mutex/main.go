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

// package main
// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// type RateLimiter struct{
// 	mu sync.Mutex
// 	count int
// }

// func (r *RateLimiter)Allow()bool{
// 	defer r.mu.Unlock()
// 	r.mu.Lock()
// 	if r.count < 5{
// 		r.count++
// 		return true

// 	}else{
// 		fmt.Println("request is more than 5/persecond")
// 		return false
// 	}

// }

// func (r *RateLimiter)StartReset(){
// 	ticker := time.NewTicker(time.Second)
// 	for{
// 		<-ticker.C
// 		r.mu.Lock()
// 		r.count=0
// 		r.mu.Unlock()
// 	}
// }
// func main(){
// 	var wg sync.WaitGroup
// 	wg.Add(11)
// 	r:=RateLimiter{count: 0,}

// 	for i:=0;i<10;i++{
// 		go func(){
// 			defer wg.Done()
// 			b:=r.Allow()
// 			if b == true{
// 				fmt.Println("accepted")
// 				return
// 			}else{
// 				fmt.Println("access denied")
// 				return
// 			}
// 		}()	
// 	}
// 	go func(){
// 		defer wg.Done()
// 		go r.StartReset()
// 	}()
// 	wg.Wait()
// }

// package main

// import(
// 	"fmt"
// 	"sync"
// )

// type Product struct{
// 	mu sync.Mutex
// 	Name string
// 	Stock int
// }
// func (p *Product)Buying(){
// 	p.mu.Lock()
// 	defer p.mu.Unlock()
// 	if p.Stock <=0{
// 		fmt.Println("inventory is not enough")
// 		return
// 	}else{
		
// 		p.Stock--
		
// 		fmt.Println("the product ",p.Name," added to your Shopping Cart")
// 	} 
// }



// func main(){
// 	p:=Product{Name:"scarf",Stock:1}
// 	var wg sync.WaitGroup
// 	wg.Add(7)
// 	for i:=0;i<7;i++{
// 		go func(){
// 				defer wg.Done()
// 				p.Buying()
// 		}()
// 	}
// 	wg.Wait()
// }
// package main
// import (
// 	"sync"
// )

// func main(){
// 	var mu1 sync.Mutex
// 	var mu2 sync.Mutex
// 	x:=10
// 	go	func(){
// 		mu1.Lock()
// 		mu2.Lock()
// 		x++
// 		mu2.Unlock()
// 		mu1.Unlock()
		
// 	}
// 	go func(){
// 		mu1.Lock()
// 		mu2.Lock()
// 		x++
// 		mu2.Unlock()
// 		mu1.Unlock()		
// 	}

// }

package main


import(
	"fmt"
	"sync"
)

type Counter struct {
    mu    sync.Mutex
    value int
}
func (c *Counter) Increment() {
    c.mu.Lock()
	c.value++
	c.mu.Unlock()
}
func main(){
	c:=Counter{value:0}
	var wg sync.WaitGroup
	wg.Add(100)
	for i:=0;i<100;i++{
		go func(){
			defer wg.Done()
			for j:=0;j<1000;j++{
				c.Increment()
			}
		}()
	}
	wg.Wait()
	fmt.Println(c.value)
}


//end of program

