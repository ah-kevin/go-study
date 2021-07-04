package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main1() {
	exitChan := make(chan bool, 1)
	wg.Add(1)
	go worker(exitChan)
	// 如何优雅的实现结束自goroutine
	exitChan <- true
	wg.Wait()
	fmt.Println("over")
}

func worker(ch <-chan bool) {
	defer wg.Done()
Label:
	for {
		select {
		case <-ch:
			break Label
		default:
			fmt.Println("worker...")
			time.Sleep(time.Second)
		}

	}
}
