package jargo

import (
	"archive/zip"
	"bufio"
	"bytes"
	"io"
	"log"
	"strings"
)

type Manifest map[string]string

type JarInfo struct {
	*Manifest
	Files []string
}

const MANIFEST_FULL_NAME = "META-INF/MANIFEST.MF"

// GetManifest extracts the manifest info from a Java JAR file
// It returns a pointer to a Manifest (map[string]string) which is the key:values pairs from the META-INF/MANIFEST.MF file
func GetManifest(filename string) (error, *Manifest) {
	err, jar := jmake(filename, false)
	if err != nil {
		return err, nil
	}
	return nil, jar.Manifest
}

// GetJarInfo extracts various info from a Java JAR file
// It extracts the Manifest (like GetManifest)
// It extracts an array of the filenames in the JAR file
// It returns a pointer to a JarInfo struct
func GetJarInfo(filename string) (error, *JarInfo) {
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
	jar.Manifest = makeManifestMap(lines)
	return nil, jar
}

func makeManifestMap(lines []string) (error *Manifest) {
	manifestMap := make(Manifest)

	for _, line := range lines {
		i := strings.Index(line, ":")
		if i == -1 {
			log.Fatal(line)
		}
		key := strings.TrimSpace(line[0:i])
		value := strings.TrimSpace(line[i+1:])
		manifestMap[key] = value
	}
	return &manifestMap
}
