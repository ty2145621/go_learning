package interview

import (
	"fmt"
	"sync"
)

// 两个协程交替打印奇偶数

func printOdd() {
	oddsC := make(chan bool, 1)
	endC := make(chan bool, 1)
	i := 1
	limit := 100

	stopC := make(chan bool)
	oddsC <- true
	// 打印奇数
	go func() {
		for {
			select {
			case <-oddsC:
				fmt.Println(i)
				i++
				if i > limit {
					close(stopC)
				}
				endC <- true
			case <-stopC:
				return
			}
		}
	}()

	// 打印偶数
	go func() {
		for {
			select {
			case <-endC:
				fmt.Println(i)
				i++
				if i > limit {
					close(stopC)
				}
				oddsC <- true
			case <-stopC:
				return
			}
		}
	}()

	<-stopC
}

// 生产这消费者，一个协程生产 1-100 之内的整数。 两个协程消费，输出总的累加和
func product() {
	productChan := make(chan int, 10)
	resultChan := make(chan int, 10)

	wg := sync.WaitGroup{}
	wg.Add(3)
	// 生产者
	go func() {
		defer wg.Done()
		for i := 0; i <= 100; i++ {
			productChan <- i
		}
		close(productChan)
	}()

	// 消费者 2个
	go func() {
		defer wg.Done()
		sum := 0
		for {
			c, ok := <-productChan
			if !ok {
				resultChan <- sum
				return
			}
			sum += c
		}
	}()

	go func() {
		defer wg.Done()
		sum := 0
		for {
			c, ok := <-productChan
			if !ok {
				resultChan <- sum
				return
			}
			sum += c
		}
	}()

	wg.Wait()
	close(resultChan)
	sum := 0
	for {
		c, ok := <-resultChan
		if !ok {
			break
		}
		sum += c
	}

	fmt.Println("sum:", sum)
}
