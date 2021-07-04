package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg2 sync.WaitGroup

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg2.Add(1)
	go worker1(ctx)
	// 如何优雅的实现结束自goroutine
	time.Sleep(time.Second * 5)
	cancel()
	wg2.Wait()
	fmt.Println("over")
}

func worker1(ctx context.Context) {
	defer wg2.Done()
Label:
	for {
		fmt.Println("worker...")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			break Label
		default:

		}

	}
}
