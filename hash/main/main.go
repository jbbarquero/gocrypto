package main

import (
	"fmt"
	"os"

	hash "github.com/jbbarquero/gocrypto/hash"
)

func printUsage() {
	fmt.Printf("Usage: %s %s %s\n", os.Args[0], "<option: [HASH|HASH_PASS]>", "[<filename>|password]")
	fmt.Printf("Example: %s %s %s\n", os.Args[0], "HASH", "image.iso")
	fmt.Printf("Example: %s %s %s\n", os.Args[0], "HASH_PASS", "password")
}

func readOptionFromArgs() string {
	return os.Args[1]
}

func main() {
	if len(os.Args) < 3 {
		printUsage()
		os.Exit(1)
	}
	option := os.Args[1]
	fmt.Println(option)

	if option == "HASH" {
		filename := os.Args[2]
		fmt.Println("HASH ", filename)

		hash.HashFile(filename)
	} else if option == "HASH_PASS" {
		password := os.Args[2]
		fmt.Println("HASH PASSWORD", password)

		hash.HashPassword(password)
	} else {
		printUsage()
		os.Exit(1)
	}

}
