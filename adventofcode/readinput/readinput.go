package readinput

import (
	"bufio"
	"log"
	"os"
)

func ReadInput(filePath string) *bufio.Scanner {
	f, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	return scanner
}
