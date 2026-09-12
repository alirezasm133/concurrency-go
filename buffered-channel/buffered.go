package main


import(
	"fmt"
)



func main(){
	ch:=make(chan int,3)
	// go func(){

	// }()
	ch<-10
	ch<-20
	ch<-30
	//ch<-40
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)		
}