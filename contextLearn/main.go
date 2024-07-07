// @Author huzejun 2024/7/6 21:17:00
/*package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	//超时控制
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	urls := []string{
		"http://api.dream.com/users", // demo
		"http://api.dream.com/products",
		"http://api.dream.com/orders",
	}
	results := make(chan string, 3)
	for _, url := range urls {
		//3个协程 共用5秒的超时
		//go fetchAPI(ctx, url, results)
		//ctx := context.WithValue(ctx, "url", url)
		fetchAPI(ctx, url, results)
	}
	for range urls {
		fmt.Println(<-results)
	}

}

func fetchAPI(ctx context.Context, url string, results chan string) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		results <- fmt.Sprintf("create request error:%v", err)
		return
	}
	client := http.DefaultClient
	response, err := client.Do(req)
	if err != nil {
		results <- fmt.Sprintf("fetch api err:%v", err)
		return
	}
	defer response.Body.Close()
	results <- fmt.Sprintf("fetch api success:%s,status=%d", url, response.StatusCode)
}
*/

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.WithValue(context.Background(), "userId", 123)
	go performTask(ctx)
	ctx = context.WithValue(ctx, "userId1", 456)
	go performTask1(ctx)
	time.Sleep(time.Second)
}

func performTask(ctx context.Context) {
	value := ctx.Value("userId")
	fmt.Println("userId:", value)
	value1 := ctx.Value("userId1")
	fmt.Println("userId1:", value1)
}

func performTask1(ctx context.Context) {
	value := ctx.Value("userId")
	fmt.Println("performTask1 userId:", value)
	value1 := ctx.Value("userId1")
	fmt.Println("performTask1 userId1:", value1)
}
