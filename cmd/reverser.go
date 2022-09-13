package main

import (
	"bufio"
	"fmt"
	"os"
)

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func main() {
	if stdinFileInfo, err := os.Stdin.Stat(); err == nil && (stdinFileInfo.Mode()&os.ModeCharDevice) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			input := scanner.Text()
			fmt.Println(Reverse(input))
		}
	} else {
		panic("stdin")
	}
}
