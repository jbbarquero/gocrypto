package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const minSizeLargeFile = 1048576

func hashSmallFile(filename string) {
	fmt.Printf("HASH small file %s\n\n", filename)

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	// Hash the file and output results
	fmt.Printf("Md5: %x\n\n", md5.Sum(data))
	fmt.Printf("Sha1: %x\n\n", sha1.Sum(data))
	fmt.Printf("Sha256: %x\n\n", sha256.Sum256(data))
	fmt.Printf("Sha512: %x\n\n", sha512.Sum512(data))
}

func hashLargeFile(filename string) {
	fmt.Println("HASH LARGE file", filename)
}

func HashFile(filename string) {

	info, err := os.Stat(filename)
	if err != nil {
		fmt.Println("Error accesing the file", filename)
		os.Exit(2)
	}

	size := info.Size()
	fmt.Printf("Filename %s has %d bytes\n", filename, size)

	if size > minSizeLargeFile {
		hashLargeFile(filename)
	} else {
		hashSmallFile(filename)
	}

}
