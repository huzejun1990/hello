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

/*package main

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
*/

/*package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	ctx = context.WithoutCancel(ctx)
	go proformTask(ctx)
	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}

func proformTask(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			//去做一些资深释放的事情
			fmt.Println("task canceled")
			return
		default:
			fmt.Println("task running")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
*/

/*package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func main() {
	//创建一个带有超时的上下文
	ctx, cancle := context.WithTimeoutCause(context.Background(), 2*time.Second, errors.New("超时了"))
	//ctx, cancle := context.WithTimeout(context.Background(), 2*time.Second)
	//ctx, cancle = context.WithTimeout(ctx, 4*time.Second)
	defer cancle() //在程序结束前调用 cancel 函数释放资源
	//启动一个任务
	go task(ctx)
	//等待一段时间，超时后任务会被取消
	time.Sleep(3 * time.Second)
	fmt.Println("Main goroutine: Done")
}

func task(ctx context.Context) {
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Task: Finished")
	case <-ctx.Done():
		fmt.Printf("Task: Context cancelled or timed out:%v, cause=%v \n", ctx.Err(), context.Cause(ctx))
	}
}*/

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() //确保在函数结束时取消context
	//设置一个2秒后执行的函数
	stop := context.AfterFunc(ctx, func() {
		fmt.Println("AfterFunc executed")
	})
	fmt.Println("stop", stop())
	go performTask(ctx, stop)
	//阻塞主goroutine,防止程序立即退出
	time.Sleep(3 * time.Second)
}

func performTask(ctx context.Context, stop func() bool) {
	select {
	case <-ctx.Done():
		fmt.Println("Content canceled:", ctx.Err())
		//fmt.Println("stop", stop())

	}
}
