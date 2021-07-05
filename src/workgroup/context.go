package workgroup

import "context"

// Context 上下文创建使用上下文取消执行的函数。
func Context(ctx context.Context) RunFunc {
	return func(stop <-chan struct{}) error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-stop:
			return nil
		}
	}
}
