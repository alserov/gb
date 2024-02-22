package main

import (
	"os"
)

// Task 3
func main() {
	f, err := os.OpenFile("file.txt", os.O_RDONLY|os.O_CREATE, 0777)
	if err != nil {
		panic(err)
	}
	f.Close()

	defer func() {
		err = os.Remove("file.txt")
		if err != nil {
			panic(err)
		}
	}()

	f, _ = os.Open("file.txt")
	if _, err = f.Write([]byte("text")); err == nil {
		panic("expected error")
	}
	f.Close()
}
