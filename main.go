package jargo

import (
	"archive/zip"
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	// Open a zip archive for reading.
	if len(os.Args) != 2 {
		log.Fatal("Needs JAR argument")
	}
	r, err := zip.OpenReader(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	// Iterate through the files in the archive,
	// printing some of their contents.

	var (
		part   []byte
		prefix bool
	)

	lines := make([]string, 0)
	lineNumber := -1
	for _, f := range r.File {
		//fmt.Printf("Contents of %s:\n", f.Name)
		if f.Name == "META-INF/MANIFEST.MF" {
			fmt.Println("***************")
			rc, err := f.Open()
			if err != nil {
				log.Fatal(err)
			}
			reader := bufio.NewReader(rc)
			buffer := bytes.NewBuffer(make([]byte, 0))

			for {
				if part, prefix, err = reader.ReadLine(); err != nil {
					break
				}
				buffer.Write(part)
				if !prefix {
					//lines = append(lines, buffer.String())
					line := buffer.String()
					if line[0] == ' ' {
						lines[lineNumber] = lines[lineNumber] + line
					} else {
						lines = append(lines, line)
						lineNumber = lineNumber + 1
					}
					fmt.Println(buffer.String())
					buffer.Reset()
				}
			}
			if err == io.EOF {
				err = nil
			}

			rc.Close()
			fmt.Println("***************")
		}
	}
	fmt.Println(lines)
	fmt.Println(len(lines))
}
