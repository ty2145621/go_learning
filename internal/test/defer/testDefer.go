package main

import (
	"fmt"
)

func testDefer1() int {
	i := 0
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(2)
	}()
	defer func() {
		fmt.Println("i =", i)
		fmt.Println(3)
		i = 4
	}()
	i = 5

	return i
}

func testDefer2() (i int) {
	i = 0
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(2)
	}()
	defer func() {
		fmt.Println("i =", i)
		fmt.Println(3)
		i = 4
	}()
	i = 5

	return i
}

func testDefer3() (j int) {
	i := 0
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(2)
	}()
	defer func() {
		fmt.Println("i =", i)
		fmt.Println("j =", j)
		fmt.Println(3)
		i = 4
	}()
	i = 5

	return i
}

func main() {
	// fmt.Println(testDefer1())
	fmt.Println(testDefer2())
	// fmt.Println(testDefer3())
}
