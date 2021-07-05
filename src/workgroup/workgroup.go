package workgroup

// RunFunc 是一个在自己的 goroutine 中与其他相关函数一起执行的函数。
// 传递给 RunFunc 的channel关闭应该触发返回。
type RunFunc func(<-chan struct{}) error

// Group 是一组相关的 goroutine。
// Group 的零值无需初始化即可完全使用。
type Group struct {
	fns []RunFunc
}

// Add 向 Group 添加一个函数。
//当 Run 被调用时，该函数将在自己的 goroutine 中执行。
//添加必须在运行之前调用。
func (g *Group) Add(fn RunFunc) {
	g.fns = append(g.fns, fn)
}

// Run 在自己的 goroutine 中执行通过 Add 注册的每个函数。
// Run 直到所有函数都返回，然后从它们返回第一个非零错误（如果有）。
// 要返回的第一个函数将触发传递给每个函数的通道的关闭，而每个函数又应该返回。
func (g *Group) Run() error {
	if len(g.fns) == 0 {
		return nil
	}
	stop := make(chan struct{})
	done := make(chan error, len(g.fns))
	defer close(done)
	for _, fn := range g.fns {
		go func(fn RunFunc) {
			done <- fn(stop)
		}(fn)
	}
	var err error
	for i := 0; i < cap(done); i++ {
		if err == nil {
			err = <-done
		} else {
			<-done
		}
		if i == 0 {
			close(stop)
		}
	}
	return err
}
