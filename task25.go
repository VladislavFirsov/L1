package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println("count:", i)
		mySleep(1)
	}

}

func mySleep(seconds time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), seconds*time.Second)
	defer cancel()
	<-ctx.Done()

}
