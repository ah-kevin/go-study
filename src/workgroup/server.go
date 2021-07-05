package workgroup

// Server 创建取消执行的函数。
//  传入的第一个函数应该阻塞。
//  传入的第二个函数应该解除第一个函数的阻塞。
func Server(serve func() error, shutdown func() error) RunFunc {
	return func(stop <-chan struct{}) error {
		done := make(chan error)
		defer close(done)

		go func() {
			done <- serve()
		}()

		select {
		case err := <-done:
			return err
		case <-stop:
			err := shutdown()
			if err == nil {
				err = <-done
			} else {
				<-done
			}
			return err
		}
	}
}
