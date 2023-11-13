package code

import (
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"sync"
	"time"
	"unsafe"
)

type student struct {
	Name string
}

// test1
// golang中有规定，switch type的case T1，类型列表只有一个，那么v := m.(type)中的v的类型就是T1类型。
//
// 如果是case T1, T2，类型列表中有多个，那v的类型还是多对应接口的类型，也就是m的类型。
//
// 所以这里msg的类型还是interface{}，所以他没有Name这个字段，编译阶段就会报错。具体解释见： https://golang.org/ref/spec#Type_switches
func test1(v interface{}) string {
	switch msg := v.(type) {
	case *student, student:
		// return msg.Name
		return msg.(student).Name
	}
	return ""
}

type People struct {
	Name string
}

// String
// 在golang中String() string 方法实际上是实现了String的接口的，该接口定义在fmt/print.go 中：
//
//	type Stringer interface {
//		String() string
//	}
//
// 在使用 fmt 包中的打印方法时，如果类型实现了这个接口，会直接调用。而题目中打印 p 的时候会直接调用 p 实现的 String() 方法，然后就产生了循环调用。
func (p *People) String() string {
	// fmt.Sprintf format %v with arg p causes recursive (*go_learning/interview/code.People).String method call
	// return fmt.Sprintf("print: %v", p)
	return ""
}

func test2() {
	// map的value本身是不可寻址的，因为map中的值会在内存中移动，并且旧的指针地址在map改变时会变得无效。
	// 故如果需要修改map值，可以将map中的非指针类型value，修改为指针类型，比如使用map[string]*Student.
	m := map[string]student{"people": {"zhoujielun"}}
	// Cannot assign to m["people"].Name
	// m["people"].Name = "myName"
	s := m["people"]
	s.Name = "myName"
	m["people"] = s

	// 或者 m := map[string]*student
}

func test3() {
	var i byte
	// Condition 'i <= 255' is always 'true'
	for i = 0; i <= 255; i++ {
	}
}

// 正在被执行的 goroutine 发生以下情况时让出当前 goroutine 的执行权，并调度后面的 goroutine 执行：
//
// IO 操作
// Channel 阻塞
// system call
// 运行较长时间
// 如果一个 goroutine 执行时间太长，scheduler 会在其 G 对象上打上一个标志（ preempt），
// 当这个 goroutine 内部发生函数调用的时候，会先主动检查这个标志，如果为 true 则会让出执行权。

func test4() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("i: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

// test5 golang 语言中没有继承概念，只有组合，也没有虚方法，更没有重载。因此，*Teacher 的 ShowB 不会覆写被组合的 People 的方法。
func test5() {
	t := Teacher{}
	t.ShowA()
	fmt.Println("-----")
	t.ShowB()
	fmt.Println("-----")
	t.People.ShowA()
	fmt.Println("-----")
	t.People.ShowB()
	fmt.Println("-----")

}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

// test6 defer 在定义的时候会计算好调用函数的参数，所以会优先输出10、20 两个参数。然后根据定义的顺序倒序执行。
func test6() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}

type Talk interface {
	Speak(string) string
}

type Student struct {
	Name string
}

func (stu *Student) Speak(think string) (talk string) {
	if think == "good" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func test7() {
	// Cannot use 'Student{}' (type Student) as the type Talk
	// Type does not implement 'Talk' as the 'Speak' method has a pointer receiver
	// var talk Talk = Student{}
	var talk Talk = &Student{}
	think := "bitch"
	fmt.Println(talk.Speak(think))
}

func test8() {
	var stu *Student
	var stuTalk Talk = stu
	var talk Talk
	var talkPtr *Talk
	fmt.Println("stu == nil", stu == nil)             // true
	fmt.Println("stuTalk == nil", stuTalk == nil)     // false
	fmt.Printf("stuTalk：%p, %+v\n", stuTalk, stuTalk) // 0x0, <nil>
	fmt.Println("talk == nil", talk == nil)           // true
	fmt.Println("talkPtr == nil", talkPtr == nil)     // true
}

func test9() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(time.Millisecond)
		wg.Done()
		wg.Add(1) // panic: 在wg.wait之后，不能调用add
	}()
	wg.Wait()
}

var lock sync.Mutex

type lazysingleton struct {
}

var instance *lazysingleton
var initInstance int

func GetInstance() *lazysingleton {
	if initInstance == 1 {
		return instance
	}
	lock.Lock()
	defer lock.Unlock()
	if initInstance == 0 {
		instance = &lazysingleton{} // unnecessary locking if instance already created
		initInstance = 1
	}

	return instance
}

func test10() {
	var ch chan int
	go func() {
		ch = make(chan int, 1)
		ch <- 1
	}()
	go func(ch chan int) {
		time.Sleep(time.Second)
		t := <-ch // 卡住，因为ch是传值的地址，一直为0x0
		fmt.Println(t)
	}(ch)
	c := time.Tick(1 * time.Second)
	for range c {
		fmt.Printf("#goroutines: %d\n", runtime.NumGoroutine())
	}
}

func test11() {
	var ch chan int
	// ch <- 1
	// close(ch)  // panic: close of nil channel
	<-ch
}

func test12() {
	// var x string = nil  // string 不能赋值为nil
	var x string
	// if x == nil {  // string 不能与nil比较
	if x == "" {
		x = "default"
	}
	fmt.Println(x)
}

const (
	a = iota
	b = iota
)
const (
	name = "menglu"
	c    = iota
	_    = iota
	d    = iota
)

func test13() {
	str1 := []string{"a", "b", "c"}
	str2 := str1[1:]
	str2[1] = "new"
	fmt.Println(str1)
	str2 = append(str2, "z", "x", "y")
	fmt.Println(str1)
	fmt.Println(str2)
}

func test14() {
	fmt.Println(&Student{Name: "menglu"} == &Student{Name: "menglu"}) // false
	fmt.Println(Student{Name: "menglu"} == Student{Name: "menglu"})   // true

	fmt.Println([...]string{"1"} == [...]string{"1"}) // 数组可以比较
	// fmt.Println([]string{"1"} == []string{"1"}) // 切片不能等于 Invalid operation: the operator == is not defined on []string
}

func test15() {
	kv := map[string]Student{"menglu": {Name: "21"}}
	// kv["menglu"].Name = "23"  // 不能更改map中的value值

	kv2 := map[string]*Student{"menglu": {Name: "21"}}
	kv2["menglu"].Name = "23"
	s := []Student{{Name: "23"}}
	s[0].Name = "22"
	fmt.Println(kv, s)
}

func test16() {
	ch := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go product16(wg, ch)
	go consumer16(wg, ch)
	wg.Wait()
}

func product16(wg *sync.WaitGroup, in chan<- int) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		in <- rand.Intn(5)
	}
	close(in)
}

func consumer16(wg *sync.WaitGroup, out <-chan int) {
	defer wg.Done()
	for i := range out {
		fmt.Println(i)
	}
}

func test17() {
	var ch chan int
	fmt.Println(ch == nil)

	ch = make(chan int)
	fmt.Println(ch == nil)

	close(ch)
	fmt.Println(ch == nil)

	ch <- 1
}

func test18() {
	go func() {
		t := time.NewTicker(1 * time.Second)
		for {
			select {
			case <-t.C:
				fmt.Println("tick")
			}
		}
	}()

	select {}
}

// test19 WaitTimeout，还是用test20好一些
func test19(wg *sync.WaitGroup, timeout time.Duration) bool {
	// 要求sync.WaitGroup支持timeout功能
	// 如果timeout到了超时时间返回true
	// 如果WaitGroup自然结束返回false
	ch := make(chan bool, 1)

	go time.AfterFunc(timeout, func() {
		ch <- true // 这里也有问题，ch不能close
	})

	go func() {
		wg.Wait()   // 这里有协程内存泄漏吧
		ch <- false // 这里也有问题，ch不能close
	}()

	return <-ch
}

// test20 WaitTimeout
func test20(wg *sync.WaitGroup, timeout time.Duration) bool {
	c := make(chan struct{})
	go func() {
		defer close(c)
		wg.Wait()
	}()
	select {
	case <-c:
		return false // completed normally
	case <-time.After(timeout):
		return true // timed out
	}
}

func test21() {
	arr := make([]int, 0)
	for i := 0; i < 10000; i++ {
		if len(arr) == cap(arr) {
			fmt.Println("len 为", len(arr), "cap 为", cap(arr))
		}
		arr = append(arr, i)
	}
}

func test22() {
	a := make([]int, 3, 4)
	a[0] = 1
	changeSlice(a)
	fmt.Println(a)
}

func changeSlice(a []int) {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&a))
	sh.Len = 4
}
