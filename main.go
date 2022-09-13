package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
	cmd := exec.Command("go", "run", "cmd/reverser.go")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Panicln("stdin", err)
	}
	defer stdin.Close()

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		log.Panicln("start", err)
	}

	for {
		input := time.Now().String()
		io.WriteString(stdin, input+"\n")
		fmt.Println(input)

		time.Sleep(time.Second)
	}
}
