package main

import (
	"bytes"
	"fmt"
	"time"
)

func testSlice1() {
	path := []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path, '/')
	dir1 := path[2:sepIndex:10]
	fmt.Println(len(dir1), cap(dir1))
}

type field struct {
	name string
}

func (p *field) print1() {
	fmt.Println(p.name)
}

func printTest() {
	data := []field{{"one"}, {"two"}, {"three"}}
	for _, v := range data {
		// 解决办法：添加如下语句
		// v := v
		go v.print1()
	}
	time.Sleep(1 * time.Second) // goroutines print: three, three, three

	data2 := []*field{{"one"}, {"two"}, {"three"}} // 注意data2是指针数组
	for _, v := range data2 {
		go v.print1() // go执行是函数，函数执行之前，函数的接受对象已经传过来
	}
	time.Sleep(1 * time.Second) // goroutines print: one, two, three
}

type field2 struct {
	num int
}

func (t *field2) print2(n int) {
	fmt.Println(t.num, n)
}
func testPrint2() {
	var i int = 1
	defer fmt.Println("result2 =>", func() int { return i * 2 }())
	i++

	v := field2{1}
	fmt.Printf("%p\n", &v)
	defer v.print2(func() int { return i * 2 }())
	v = field2{2}
	fmt.Printf("%p\n", &v)
	i++

	var v2 *field2 = &field2{0}
	fmt.Printf("%p\n", v2)
	v2 = &field2{1}
	fmt.Printf("%p\n", v2)
	defer v2.print2(func() int { return i * 2 }())
	*v2 = field2{2}
	fmt.Printf("%p\n", v2)
	i++

	// prints:
	// 2 4
	// result => 2 (not ok if you expected 4)
}

type data struct {
	name string
}

func test1222() {
	m := map[string]data{"x": data{"three"}}
	// m["x"].print() //error
	d2 := m["x"]
	d2.name = "3"
	fmt.Printf("%s, %s\n", d2.name, m["x"].name)
}

func test1233() {
	m := map[string][]int{}
	m["1"] = []int{1, 11}
	m["2"] = []int{2, 22}
	m["3"] = []int{2, 22}
	m["4"] = []int{2, 22}
	// m["x"].print() //error
	d2 := m["1"]
	d2[0] = 111
	m["2"][0] = 222

	d3 := m["3"]
	d3 = append(d3, []int{111, 111, 111, 111}...)
	m["4"] = append(m["4"], []int{222, 222, 222, 222}...)

	fmt.Printf("%d, %d\n", m["1"][0], m["2"][0])
	fmt.Printf("%d\n", m["4"][3])
	fmt.Printf("%d, %d\n", m["3"][3], m["4"][3])
}

func main() {
	test1233()
}
