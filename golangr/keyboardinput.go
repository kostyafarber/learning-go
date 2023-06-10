package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a number between 0 and 10: ")
	num, _ := reader.ReadString('\n')
	num = strings.TrimSuffix(num, "\n")
	num_int, err := strconv.Atoi(num)

	if err != nil {
		panic(err)
	}

	if num_int >= 0 && num_int <= 10 {
		fmt.Printf("Your number was between 0 and 10 and it is: %s\n", num)
	} else {
		fmt.Println("Not between 10")
	}
}
