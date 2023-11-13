package main

import (
	"fmt"
	"net/http"
	"runtime"
)

func main() {
	/*m := map[int]string{0: "https://www.baidu.com", 1: "https://www.bilibili.com", 2: "https://www.weibo.com",
	3: "https://www.baidu.com", 4: "https://www.bilibili.com", 5: "https://www.weibo.com"}*/
	num := 6
	for index := 0; index < num; index++ {
		resp, _ := http.Get("https://www.baidu.com")
		fmt.Printf("%p\n", resp)
		// _, _ = ioutil.ReadAll(resp.Body)
	}
	fmt.Printf("此时goroutine个数= %d\n", runtime.NumGoroutine())
}
