package main

import (
	"fmt"
	"time"
)

// 处理队列
func handle() {

	c := make(chan int, 10)

	c <- 1
	c <- 2
	c <- 3
	timer1 := time.NewTicker(1 * time.Second)
	timer2 := time.NewTicker(5 * time.Second)
	for i := 0; i < 15; i++ {
		select {
		case data, ok := <-c:
			if i%3 == 0 {
				time.Sleep(10 * time.Second)
			}
			fmt.Println("data OK", data, ok)
			timer2.Reset(5 * time.Second)
		case <-timer1.C:
			fmt.Println("timer1 OK")
		case <-timer2.C:
			fmt.Println("timer2 OK")
		}

		if i == 6 {
			c <- 6
		}
	}
}

/*func main() {
	handle()
}
*/
