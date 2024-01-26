package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	chStop := make(chan os.Signal, 1)
	signal.Notify(chStop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			fmt.Printf("Type file path: ")

			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			if err := scanner.Err(); err != nil {
				fmt.Printf("failed to read value: %v", err)
				continue
			}

			if err := CheckFile(scanner.Text()); err != nil {
				fmt.Println(err)
			}

			fmt.Printf("Type [y] to continue or [n] to stop: ")
			scanner.Scan()
			if scanner.Text() == "n" {
				chStop <- syscall.SIGTERM
				break
			}
		}
	}()

	<-chStop
	fmt.Println("program exited")
}

func CheckFile(p string) error {
	st, err := os.Stat(strings.Trim(p, " "))
	if err != nil {
		return fmt.Errorf("failed to find file or dir: %v \n", err)
	}

	parts := strings.Split(st.Name(), ".")

	// fName := path.Base(p)
	// ext := path.Base(p)

	fmt.Printf("filename: %s \next: %s \n", parts[0], parts[1])
	return nil
}
