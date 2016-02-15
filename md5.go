package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

func main() {

	inputFile := os.Stdin
	fs, _ := inputFile.Stat()
	fm := fs.Mode()
	var fname string
	// Piped input
	if fm&os.ModeNamedPipe == os.ModeNamedPipe {
		md5Hash(inputFile)
	} else if fm&os.ModeCharDevice == os.ModeCharDevice { // unix character input
		if len(os.Args) > 1 {
			fname = os.Args[1]
		} else {
			fname = os.Args[0]
		}
		inputFile, _ = os.Open(fname)
		md5Hash(inputFile)
	} else { //stdin input
		md5Hash(inputFile)
	}
}

func md5Hash(file io.Reader) error {
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return err
	}
	fmt.Printf("%x\n", hash.Sum(nil))
	return nil
}
