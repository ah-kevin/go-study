package main

import (
	"fmt"
	"time"
)

func main() {
	process("2222")
}

func search(term string) (string, error) {
	time.Sleep(200 * time.Millisecond)
	return "some value", nil
}

func process(term string) error {
	record, err := search(term)
	if err != nil {
		return err
	}
	fmt.Println("Received:", record)
	return nil
}
