package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	d := time.Now().Add(1 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// Even though ctx will be expired, it is good practice to call its
	// cancelation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
	defer cancel()

	for i := 0; i < 1000; i++ {
		select {
		//case <-time.After(2 * time.Second):
		//	fmt.Println("oversleep")
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		default:
			time.Sleep(100000 * time.Microsecond)
			fmt.Println("run .... ", i)
		}
	}
}
