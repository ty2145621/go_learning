package main

import (
	"fmt"
	"time"
)

func testDebug(c chan error) error {
	time.Sleep(3 * time.Second)
	c <- fmt.Errorf("sleep too long")
	return nil
}

func main() {
	start := time.Now()
	errC := make(chan error, 2)
	go testDebug(errC)
	select {
	case <-time.After(1000 * time.Millisecond):
		fmt.Printf("callee timeout:%+v\n", 1000)
	case err := <-errC:
		fmt.Errorf("testDebug run success")
		fmt.Println(err)
	}
	fmt.Println("cost:", time.Since(start).Milliseconds())

}
