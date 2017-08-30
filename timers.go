package main

import "time"
import "fmt"

func main() {
	timer1 := time.NewTimer(time.Second * 2)

	<-timer1.C
	fmt.Println("Timer 1 expired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		// 返回的是当前时间
		a := <-timer2.C
		fmt.Println(a)
		fmt.Println("Timer 2 expired")
	}()

	time.Sleep(time.Second * 5)
	// cancel the timer before it expires
	// stop2 := timer2.Stop()
	// if stop2 {
	// 	fmt.Println("Timer 2 stopped")
	// }
}
