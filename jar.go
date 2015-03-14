package jargo

import (
	"archive/zip"
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"strings"
)

type Manifest struct {
	ManifestMap map[string]string
}

type JarInfo struct {
	Manifest
	Files []string
}

const MANIFEST_FULL_NAME = "META-INF/MANIFEST.MF"

func MakeManifest(filename string) (error, *Manifest) {
	err, jar := jmake(filename, false)
	if err != nil {
		return err, nil
	}
	return nil, &jar.Manifest
}

func MakeJarInfo(filename string) (error, *JarInfo) {
	return jmake(filename, true)
}

func jmake(filename string, fullJar bool) (error, *JarInfo) {
	r, err := zip.OpenReader(filename)
	if err != nil {
		log.Println(err)
		return err, nil
	}
	defer r.Close()

	var (
		part   []byte
		prefix bool
		lines  []string
	)

	jar := new(JarInfo)
	if fullJar {
		lines = make([]string, 0)
	}
	lineNumber := -1
	for _, f := range r.File {
		if fullJar {
			jar.Files = append(jar.Files, f.Name)
		}
		if f.Name == MANIFEST_FULL_NAME {
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
				if len(part) == 0 {
					continue
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
					buffer.Reset()
				}
			}
			if err == io.EOF {
				err = nil
			}
			rc.Close()
		}
	}
	jar.ManifestMap = makeManifestMap(lines)

	fmt.Println("NumFiles")
	return nil, jar
}

func makeManifestMap(lines []string) (error map[string]string) {
	manifestMap := make(map[string]string)

	for _, line := range lines {
		i := strings.Index(line, ":")
		if i == -1 {
			log.Fatal(line)
		}
		key := strings.TrimSpace(line[0:i])
		value := strings.TrimSpace(line[i+1:])
		manifestMap[key] = value
	}
	return manifestMap
}
