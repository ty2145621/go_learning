package concurrency

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

type item struct {
	ID    int64
	Value string
}

type MySyncMap atomic.Value // map[string]*item

func NewSyncMap() MySyncMap {
	m := map[string]*item{}
	newAtomValue := atomic.Value{}
	newAtomValue.Store(m)
	return MySyncMap(newAtomValue)
}

func (m MySyncMap) Get(key string) (value *item, ok bool) {
	am := atomic.Value(m)
	oldMap, ok := am.Load().(map[string]*item)
	if !ok {
		return nil, ok
	}
	value, ok = oldMap[key]
	return value, ok
}

// Set true-成功 false-失败
func (m MySyncMap) Set(key string, value *item) bool {
	am := atomic.Value(m)
	oldMap, ok := am.Load().(map[string]*item)
	if !ok {
		return ok
	}
	oldMap[key] = value
	return true
}

func syncMap() {
	mySyncMap := NewSyncMap()
	fmt.Println(mySyncMap.Get("test"))
	fmt.Println(mySyncMap.Set("test", &item{
		ID:    1,
		Value: "2",
	}))
	fmt.Println(mySyncMap.Get("test"))
}

func GetNews(num int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	millisec := r.Intn(1000)
	time.Sleep(time.Duration(millisec) * time.Millisecond)
	return strconv.Itoa(num)
}

// 同步数据
func channel1_1() {
	defer func(start time.Time) {
		fmt.Printf("cost:%d ms", time.Since(start).Milliseconds())
	}(time.Now())
	length := 10
	res := make(chan string, length)
	wg := sync.WaitGroup{}
	for i := 0; i < length; i++ {
		wg.Add(1)
		i := i // 重要
		go func() {
			defer wg.Done()
			res <- GetNews(i)
		}()
	}
	wg.Wait()
	close(res)
	resArr := make([]string, 0, length)
	for v := range res {
		resArr = append(resArr, v)
	}
	fmt.Println(resArr)
}

func channel2_1() {
	defer func(start time.Time) {
		fmt.Printf("cost:%d ms", time.Since(start).Milliseconds())
	}(time.Now())
	length := 10
	maxConcurrentNum := 3
	conCh := make(chan struct{}, maxConcurrentNum)
	res := make(chan string, length)
	wg := sync.WaitGroup{}
	for i := 0; i < length; i++ {
		wg.Add(1)
		i := i // 重要
		go func() {
			conCh <- struct{}{} // 也可以放在 go func 外边，体会下不同之处
			defer func() { <-conCh }()
			defer wg.Done()
			res <- GetNews(i)
		}()
	}
	wg.Wait()
	close(res)
	resArr := make([]string, 0, length)
	for v := range res {
		resArr = append(resArr, v)
	}
	fmt.Println(resArr)
}

type Glimit struct {
	n int
	c chan struct{}
}

// initialization Glimit struct
func New(n int) *Glimit {
	return &Glimit{
		n: n,
		c: make(chan struct{}, n),
	}
}

// Run f in a new goroutine but with limit.
func (g *Glimit) Run(f func()) {
	g.c <- struct{}{}
	go func() {
		f()
		<-g.c
	}()
}

var wg = sync.WaitGroup{}

func channel2_2() {
	length := 10
	res := make(chan string, length)

	g := New(2)
	for i := 0; i < length; i++ {
		i := i
		wg.Add(1)
		goFunc := func() {
			defer wg.Done()
			res <- GetNews(i)
		}
		g.Run(goFunc)
	}
	wg.Wait()
	close(res)
	resArr := make([]string, 0, length)
	for v := range res {
		resArr = append(resArr, v)
	}
	fmt.Println(resArr)
}

func loop(num int, stop <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	t := time.NewTicker(200 * time.Millisecond)
	for {
		select {
		case <-t.C:
			fmt.Printf("%d, %s\n", num, time.Now())
		case <-stop:
			fmt.Printf("stop chan %d\n", num)
			return
		}
	}
}
func waitSomeTimeAndStop(stop chan<- struct{}) {
	time.Sleep(1 * time.Second)
	close(stop)
}
func channel3_1() {
	stop := make(chan struct{})
	wg := &sync.WaitGroup{} // 注意要用引用
	wg.Add(2)
	go loop(1, stop, wg)
	go loop(2, stop, wg)
	waitSomeTimeAndStop(stop)
	wg.Wait()
	fmt.Println("stop success")
}

func channel4_1() {
	rspCh := make(chan string, 10)
	t := time.NewTicker(100 * time.Millisecond)
	for {
		t.Reset(100 * time.Millisecond) // 此处有隐患，t.C可能未清空,rsp处理最好是另起协程
		select {
		case rsp := <-rspCh:
			fmt.Printf("%s\n", rsp)
		case <-t.C:
			fmt.Printf("time exceeds\n")
		}
	}
}

func testTimerReset() {
	t := time.NewTicker(100 * time.Millisecond)
	start := time.Now()
	time.Sleep(200 * time.Millisecond)
	<-t.C
	fmt.Println(time.Since(start).Milliseconds()) // 200
	time.Sleep(200 * time.Millisecond)
	t.Reset(200 * time.Millisecond) // reset 不会清除已有的t.C
	fmt.Println(time.Since(start).Milliseconds()) // 400
	<-t.C
	fmt.Println(time.Since(start).Milliseconds()) // 400
	<-t.C
	fmt.Println(time.Since(start).Milliseconds()) // 600
}
