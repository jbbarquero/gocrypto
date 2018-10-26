package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

const minSizeLargeFile = 10000 //1048576

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
	fmt.Printf("HASH LARGE file %s\n\n", filename)

	// Open file for reading
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create new hasher, which is a writer interface
	hasher := md5.New()

	// Default buffer size for copying is 32*1024 or 32kb per copy
	// Use io.CopyBuffer() if you want to specify the buffer to use
	// It will write 32kb at a time to the digest/hash until EOF
	// The hasher implements a Write() function making it satisfy
	// the writer interface. The Write() function performs the digest
	// at the time the data is copied/written to it. It digests
	// and processes the hash one chunk at a time as it is received.
	_, err = io.Copy(hasher, file)
	if err != nil {
		log.Fatal(err)
	}

	// Now get the final sum or checksum.
	// We pass nil to the Sum() function because
	// we already copied the bytes via the Copy to the
	// writer interface and don't need to pass any new bytes
	checksum := hasher.Sum(nil)
	fmt.Printf("Md5 checksum: %x\n", checksum)
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
