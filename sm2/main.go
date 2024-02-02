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
	scanner := bufio.NewScanner(os.Stdin)

loop:
	for {
		select {
		case <-chStop:
			break loop
		default:
			fmt.Printf("Type file path: ")

			scanner.Scan()
			if err := scanner.Err(); err != nil {
				fmt.Printf("failed to read value: %v\n", err)
				break
			}

			if _, err := os.Stat(scanner.Text()); err != nil {
				fmt.Printf("failed to find file: %v\n", err)
				break
			}

			b, err := os.ReadFile(scanner.Text())
			if err != nil {
				fmt.Printf("failed to read file: %v\n", err)
				break
			}

			stat, err := CountLetters(string(b))
			if err != nil {
				fmt.Printf("failed to count letters from file: %v", err)
				break
			}

			fmt.Println(stat)

			fmt.Printf("Type [y] to continue or [n] to stop: ")
			scanner.Scan()
			if scanner.Text() == "n" {
				break loop
			}
		}
	}
	fmt.Println("\napp was stopped")
}

func CountLetters(s string) (string, error) {
	var (
		ltrs    = make(map[int32]int)
		trimmed = strings.ReplaceAll(s, " ", "")
	)

	if len(trimmed) == 0 {
		return "", fmt.Errorf("provided invalid string")
	}

	for _, ltr := range strings.ToLower(trimmed) {
		if _, ok := ltrs[ltr]; ok {
			ltrs[ltr]++
			continue
		}
		ltrs[ltr] = 1
	}

	sb := strings.Builder{}
	for ltr, count := range ltrs {
		str := fmt.Sprintf("[%s] - %d%%\n", string(ltr), int32(float32(count)/float32(len(trimmed))*100))
		sb.WriteString(str)
	}

	return sb.String(), nil
}
