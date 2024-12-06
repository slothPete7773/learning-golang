// Reference
// https://medium.com/programming-%E7%9F%A5%E8%AD%98%E5%AE%B6/goroutine-%E5%BF%85%E9%A0%88%E7%9F%A5%E9%81%93%E7%9A%84%E4%BA%94%E4%BB%B6%E4%BA%8B%E6%83%85-c1c6816899ee

// 1. Sync data/exit with Context

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

//範例: 兄弟執行緒間不求同生只求同死，使用context

var wg sync.WaitGroup //宣告計數器

func aRoutine(ctx context.Context) {
	defer wg.Done() //當該執行緒執行到最後計數器-1
	select {
	case <-time.After(1 * time.Second): // 1秒之後繼續執行
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // context deadline exceeded
	}

}

func main() {

	const shortDuration = 11 * time.Millisecond

	d := time.Now().Add(shortDuration)
	ctx, cancel := context.WithDeadline(context.Background(), d) //宣告一個context.WithDeadline並注入1.001秒之類為執行完的執行緒將發產出ctx.Err
	defer cancel()                                               // 程式最後執行WithDeadline失效
	go aRoutine(ctx)                                             // 啟動aRoutine執行序
	wg.Add(1)                                                    // 計數器+1
	wg.Wait()                                                    //等待計數器歸零
}
