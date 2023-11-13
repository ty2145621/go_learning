## golang

### goroutine
+ 生产者消费者

  交替输出数字和字母

+ 协程调度
  - 何时发生调度：
  
  ```text
  正在被执行的 goroutine 发生以下情况时让出当前 goroutine 的执行权，并调度后面的 goroutine 执行：
  - IO 操作
  - Channel 阻塞
  - system call
  - 运行较长时间
  如果一个 goroutine 执行时间太长，scheduler 会在其 G 对象上打上一个标志（ preempt），
  当这个 goroutine 内部发生函数调用的时候，会先主动检查这个标志，如果为 true 则会让出执行权。
  ```

  - 如果只有一个调度器，即 runtime.GOMAXPROCS(1) ，以下代码输出什么。输出5,0,1,2,3,4
  ```go
  func main() {
    runtime.GOMAXPROCS(1)
    wg := sync.WaitGroup{}
    wg.Add(10)
    for i := 0; i < 6; i++ {
        go func(i int) {
            fmt.Println("i: ", i)
            wg.Done()
        }(i)
    }
    wg.Wait()
  }
  ```
  
+ slice 扩容
  ```text
  func growslice; <256双倍扩容，>256之后逐渐缩小至1.25倍，具体是 newcap += newcap/4 + 192
  ```
  
+ map 扩容和查找机制
  ```text  
    
  ```