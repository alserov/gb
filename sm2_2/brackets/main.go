package main

import "fmt"

// 4 не делал,ничего сильно не поменяется, тем более пакет уже устаревший как я знаю
// The ioutil package in Go has been deprecated and may be removed in future versions.

// Task 5
func main() {
	fmt.Println(generateBrackets(3))
}

func generateBrackets(n int) []string {
	var result []string
	backtrack(&result, "", 0, 0, n)
	return result
}

func backtrack(result *[]string, current string, open, close, max int) {
	if len(current) == max*2 {
		*result = append(*result, current)
		return
	}

	if open < max {
		backtrack(result, current+"(", open+1, close, max)
	}
	if close < open {
		backtrack(result, current+")", open, close+1, max)
	}
}
