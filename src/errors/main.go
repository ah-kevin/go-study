package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	err := bar()
	//if err == sql.ErrNoRows {
	//	fmt.Printf("data not found, %+v\n", err)
	//	return
	//}
	//if errors.Cause(err) == sql.ErrNoRows {
	if errors.Is(err, sql.ErrNoRows) {
		//fmt.Printf("data not found, %v\n", err)
		fmt.Printf("%+v\n", err)
		return
	}
}

//注意：无论是 Wrap ， WithMessage 还是 WithStack
//，当传入的 err 参数为 nil 时， 都会返回nil，
//这意味着我们在调用此方法之前无需作 nil 判断，保持了代码简洁
func foo() error {
	//return errors.WithStack(sql.ErrNoRows)
	return errors.Wrap(sql.ErrNoRows, "foo failed")
	//return fmt.Errorf("foo err, %v", sql.ErrNoRows)
	//return sql.ErrNoRows
}
func bar() error {
	return errors.WithMessage(foo(), "bar failed")
	//return foo()
}
