package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	flag "github.com/spf13/pflag"
	"os"
)

var (
	SHA384Flag bool
	SHA512Flag bool
)

func main() {
	args := os.Args
	bytesArg := []byte(args[1])

	flag.BoolVar(&SHA384Flag, "384", false, "Create SHA384 hash")
	flag.BoolVar(&SHA512Flag, "512", false, "Create SHA512 hash")
	flag.Parse()

	switch {
	case SHA384Flag:
		hash := sha512.Sum384(bytesArg)
		fmt.Printf("Creating SHA384 hash %x\n", hash)

	case SHA512Flag:
		hash := sha512.Sum512(bytesArg)
		fmt.Printf("Creating SHA512 hash %x\n", hash)

	default:
		hash := sha256.Sum256(bytesArg)
		fmt.Printf("Creating SHA256 hash: %x\n", hash)
	}

}
