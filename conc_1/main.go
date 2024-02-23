package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os/signal"
	"strconv"
	"syscall"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	count(ctx)
}

func count(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			break
		default:
			var numb string
			if _, err := fmt.Scanln(&numb); err != nil {
				if errors.Is(err, io.EOF) {
					return
				}
			}

			ch := make(chan int)

			go func() {
				n, err := strconv.Atoi(numb)
				if err != nil {
					fmt.Println("invalid data")
					return
				}
				n *= n
				ch <- n
				fmt.Printf("Squared: %d\n", n)
			}()

			go func() {
				n := <-ch
				n = n * 2
				fmt.Printf("Doubled: %d\n", n)
			}()
		}
	}
}
