package main

import (
	"bufio"
	"fmt"
	"strings"
)

func createScanner(s string, sfunc bufio.SplitFunc) *bufio.Scanner {
	scanner := strings.NewReader(s)
	scanWords := bufio.NewScanner(scanner)
	scanWords.Split(sfunc)

	return scanWords
}

type WordCounter int

func (w *WordCounter) Write(p []byte) (int, error) {
	scanner := createScanner(string(p), bufio.ScanWords)
	words := 0
	for scanner.Scan() {
		words += 1
	}

	*w += WordCounter(words)
	return words, nil
}

type LineCounter int

func (l *LineCounter) Write(p []byte) (int, error) {
	scanner := createScanner(string(p), bufio.ScanLines)
	lines := 0
	for scanner.Scan() {
		lines += 1
	}

	*l += LineCounter(lines)
	return lines, nil
}

func main() {
	var wc WordCounter
	s1 := "Is it me you are looking for?"
	fmt.Printf("Writing: %s to WordCounter\n", s1)
	wc.Write([]byte(s1))
	fmt.Printf("Printing WordCounter: %d\n", wc)

	wc = 0
	s2 := "No, it isn't "
	fmt.Printf("Writing string: %sKostya, instead using Fprintf\n", s2)
	fmt.Fprintf(&wc, "%s Kostya\n", s2)
	fmt.Printf("After setting to zero and writing again: %d\n\n", wc)

	s3 := `This is a multiline string.
It servers as an example for the line counter.
Another line.
`

	fmt.Printf("Count lines in string: \n\n%s", s3)
	var lc LineCounter
	lc.Write([]byte(s3))

	fmt.Printf("\nLines are: %d\n", lc)
}
