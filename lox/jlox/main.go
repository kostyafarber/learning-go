package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args
	if len(args) > 1 {
		fmt.Println("usage: jlox [script]")
		os.Exit(1)
	} else if len(args) == 2 {
		runFile(args[0])
	} else {
		runPrompt()
	}
}

func runFile(path string) {
	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("error reading file")
		os.Exit(1)
	}
	run(string(content))
}

func runPrompt() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("error reading line")
			os.Exit(1)
		}
		if err == io.EOF {
			break
		}
		run(line)
	}
}

func run(source string) {
	fmt.Println(source)
}
