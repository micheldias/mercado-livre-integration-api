package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()
	select {
	case <-time.After(time.Second * 15): // This task takes 15 seconds
		fmt.Println("Finished the task")
	case <-ctx.Done():
		t, ok := ctx.Deadline()
		fmt.Println(t)
		fmt.Println(ok)
		fmt.Println("We've waited too long, let's move on!") // We only wait for 10 seconds
	}
}
