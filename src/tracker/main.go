package main

import (
	"context"
	"fmt"
	"time"
)

// 1. 搞清楚把并发给调用者
// 2. 搞清楚goroutine什么时候退出,管控生命周期
// 3. 能够控制这个goroutine什么时候退出(channel, context都可以)
func main() {
	tracker := NewTracker()
	go tracker.Run()
	_ = tracker.Event(context.Background(), "test")
	_ = tracker.Event(context.Background(), "test")
	_ = tracker.Event(context.Background(), "test")
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
	defer cancel()
	tracker.Shutdown(ctx)
}

type Tracker struct {
	ch   chan string
	stop chan struct{}
}

func NewTracker() *Tracker {
	return &Tracker{
		ch: make(chan string, 10),
	}
}

func (t *Tracker) Event(ctx context.Context, data string) error {
	select {
	case t.ch <- data:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (t *Tracker) Run() {
	for data := range t.ch {
		time.Sleep(1 * time.Second)
		fmt.Println(data)
	}
	t.stop <- struct{}{}
}
func (t *Tracker) Shutdown(ctx context.Context) {
	close(t.ch)
	select {
	case <-t.stop:
	case <-ctx.Done():
	}

}
