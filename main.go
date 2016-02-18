// used to output the md5hash of the input file
/* format supported-
  1) ./main data/1.txt
  2) ./main < data/1.txt
  3)	cat data/1.txt | ./main
 	4)	./main */
package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)
const (
	Success = iota
	Failed
)
func main() {

	inputFile := os.Stdin
	defer inputFile.Close()
	fs, _ := inputFile.Stat()
	fm := fs.Mode()
	var fname string
	// bit and operation with file mode tag to identify the type of file -piped , stdin , unix chharacter
	// Piped input
		if fm&os.ModeNamedPipe == os.ModeNamedPipe {
		md5Hash(inputFile)
	} else if fm&os.ModeCharDevice == os.ModeCharDevice { // unix character input
		if len(os.Args) > 1 {
			fname = os.Args[1]
		} else {
			fname = os.Args[0]
		}
		inputFile, err := os.Open(fname)
		if err!=nil{
			fmt.Println("Error in reading file")
			os.Exit(Failed)
		}
		md5Hash(inputFile)
	} else { //stdin input
		md5Hash(inputFile)
	}
}
// used for getting the md5 checksum of the file
func md5Hash(file io.Reader) error {
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return err
	}
	fmt.Printf("%x\n", hash.Sum(nil))
	return nil
}
