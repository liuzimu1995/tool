package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	filename := flag.String("f", "", "Use -f fileName")
	if *filename == "" {
		fmt.Println("Use -f fileName")
	}
	flag.Parse()
	fileInfo, err := os.Stat(*filename)
	if err != nil {
		fmt.Println("File path error")
		return
	}
	// file size
	filesize := fileInfo.Size()
	fmt.Println(filesize)
	// file md5
	fileCont, inerr := os.Open(*filename)
	if inerr == nil {
		md5h := md5.New()
		io.Copy(md5h, fileCont)
		fmt.Printf("%x", md5h.Sum([]byte(""))) //md5
	}
}
