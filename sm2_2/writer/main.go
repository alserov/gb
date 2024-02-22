package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Task 1
func main() {
	chStop := make(chan os.Signal, 1)
	signal.Notify(chStop, syscall.SIGINT, syscall.SIGTERM)

	f, _ := os.OpenFile("out.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	w := NewWriter(f)
	defer f.Close()

	for {
		select {
		case <-chStop:
			return
		default:
			fmt.Println("Enter record:")

			var text string
			_, err := fmt.Scanln(&text)
			if err != nil {
				if errors.Is(err, io.EOF) {
					return
				}
				fmt.Println("Error: " + err.Error())
				break
			}
			if text == "exit" {
				return
			}

			if err = w.Write(text); err != nil {
				fmt.Printf("failed to create record: %v\n", err)
			}
			fmt.Println("Successfully written")
		}
	}
}

type Writer interface {
	Write(val string) error
}

func NewWriter(out io.Writer) Writer {
	return &writer{
		out: out,
	}
}

type writer struct {
	len uint
	out io.Writer
}

func (w *writer) Write(val string) error {
	if _, err := w.out.Write([]byte(fmt.Sprintf("%d.\t%v\t\t%s\n", w.len, time.Now().Format("2006-01-02 15:04:05"), val))); err != nil {
		return fmt.Errorf("failed to write: %w", err)
	}

	w.len++

	return nil
}
