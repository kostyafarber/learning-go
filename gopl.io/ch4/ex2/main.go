package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	flag "github.com/spf13/pflag"
	"log"
	"os"
)

var (
	SHA384Flag bool
	SHA512Flag bool
)

type shaPrinter struct{}
type shas string

func (s shaPrinter) sha512Printer(length shas, bytesArg []byte) {
	const (
		t84 shas = "384"
		f12      = "512"
	)

	switch {
	case length == t84:
		hash := sha512.Sum384(bytesArg)
		fmt.Printf("Creating SHA384 hash %x\n", hash)

	case length == f12:
		hash := sha512.Sum512(bytesArg)
		fmt.Printf("Creating SHA512 hash %x\n", hash)
	}
}
func main() {
	args := os.Args
	if len(args) == 1 || len(args) > 3 {
		log.Fatalf("usage: %s s1 --384|512", args[0])
		os.Exit(1)
	}

	bytesArg := []byte(args[1])

	flag.BoolVar(&SHA384Flag, "384", false, "Create SHA384 hash")
	flag.BoolVar(&SHA512Flag, "512", false, "Create SHA512 hash")
	flag.Parse()

	s := shaPrinter{}
	switch {
	case SHA384Flag:
		s.sha512Printer("384", bytesArg)
	case SHA512Flag:
		s.sha512Printer("512", bytesArg)
	default:
		hash := sha256.Sum256(bytesArg)
		fmt.Printf("Creating SHA256 hash: %x\n", hash)
	}

}
