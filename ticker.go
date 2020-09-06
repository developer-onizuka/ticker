package main

import (
	"fmt"
	"time"
)

func recv(name string, ch <-chan int, start <-chan bool, done <-chan bool) {
	for {
		select {
		case <-start:
			fmt.Printf("%v starts\n",name)
		case t := <-ch:
			fmt.Printf("%v: %v\n",t,name)
		case <-done:
			fmt.Printf("%v is done\n",name)
		}
	}
}

func main() {
	ticker := time.NewTicker(time.Second)

	//for {
	//	fmt.Printf("hello\n")
	//	<-ticker.C
	//}

	ch := make(chan int, 8)
	done1 := make(chan bool)
	done2 := make(chan bool)
	done3 := make(chan bool)
	start1 := make(chan bool)
	start2 := make(chan bool)
	start3 := make(chan bool)

	go recv("1st buff", ch, start1, done1)
	go recv("2nd buff", ch, start2, done2)
	go recv("3rd buff", ch, start3, done3)

	start1 <- true
	start2 <- true
	start3 <- true
	<-ticker.C

	i := 0
	for i < 10 {
		ch <- i
		i++
		<-ticker.C
	}

	done1 <- true
	done2 <- true
	done3 <- true
	<-ticker.C
}
