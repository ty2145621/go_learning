## golang 同步语法
#### 1. atomic.Value
使用atom.Value 来存储任意值，常用来共享map（map不支持并发访问，会panic）和其他一些有读有写的数据（如一个协程定时更新，其他协程获取数据）。共享map更常用下边的sync.map
```go
package concurrency

import (
	"fmt"
	"sync/atomic"
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
```

#### 2.sync.map
相比map+Mutex/RWMutex, 更适合读多写少的场景
注意LoadOrStore的用法，可以参考router中的用法
```go
type mapInterface interface {
	Load(interface{}) (interface{}, bool)
	Store(key, value interface{})
	LoadOrStore(key, value interface{}) (actual interface{}, loaded bool)
	LoadAndDelete(key interface{}) (value interface{}, loaded bool)
	Delete(interface{})
	Range(func(key, value interface{}) (shouldContinue bool))
}
```


#### 3. bigCache
__TODO：后续讲解__
支持并发的大规模数据（GB，百万以上）的map缓存，详情见文章
```https://pandaychen.github.io/2020/03/03/BIGCACHE-ANALYSIS/```

#### 4.同步，sync包
- 锁：sync.Mutex sync.RWMutex
- sync.Once 只执行一次，在程序开始服务前使用（http serve之前），或者懒加载在接口入口处使用
- sync.Cond 条件通知，一对多，可用`close(chan)`代替
- sync.Pool 对象池 __TODO:后续介绍实际案例__


#### 5.channel 协程间通信
- 同步数据（消息传递，结果汇总，事件订阅等），本质上就是生产者和消费者，一对一
```go
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
	len := 10
	res := make(chan string, len)
	wg := sync.WaitGroup{}
	for i := 0; i < len; i++ {
		wg.Add(1)
		i := i // 重要
		go func() {
			defer wg.Done()
			res <- GetNews(i)
		}()
	}
	wg.Wait()
	close(res)
	resArr := make([]string, 0, len)
	for v := range res {
		resArr = append(resArr, v)
	}
	fmt.Println(resArr)
}
```

- 控制并发
```go
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
```

```go

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
```

- 控制流程（close广播）
```go

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
```


- select channel 超时处理
- 只读chan`<-chan` 、只写chan`chan<-`
- 使用 `_, ok := <- ch` 判断chan是否关闭
```go 
if _, ok := <- ch; ok {
}
```

- channel 传指针并不一定比传结构体好，先了解下结论 __TODO:后续展开__


#### 6.singleflight
多请求并发获取时，只有一个实际发起请求，其他请求等待返回结果，详见db_proxy








