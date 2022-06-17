package tarutil

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func checkForNestedDir(name string) {
	contents := strings.Split(name, "/")
	path, _ := strings.Join(contents[:len(contents)-1], "/"), contents[len(contents)-1:]

	// log.Println("checking for path", path, file)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			log.Fatalf("unable to create intermediate directories %s: %v", path, err)
		}

		// fmt.Println("successfully created directory", path)
	}
}

func ExtractTarGz(name string, stream io.Reader) (string, error) {
	uncompressedStream, err := gzip.NewReader(stream)
	if err != nil {
		panic(err)
	}

	tarReader := tar.NewReader(uncompressedStream)

	rootDir, err := os.MkdirTemp("", name)
	if err != nil {
		log.Printf("cannot create root directory: %v\n", err)
		return "", err
	}

	rootDir += "/"

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			return rootDir, nil
		}

		checkForNestedDir(filepath.Join(rootDir, header.Name))

		outFile, err := os.Create(filepath.Join(rootDir, header.Name))
		if err != nil {
			log.Printf("error creating file %#v: %v", header, err)
			return "", err
		}

		if _, err = io.Copy(outFile, tarReader); err != nil {
			outFile.Close()
			log.Printf("error copying file %s: %v", filepath.Join(rootDir, header.Name), err)
			return "", err
		}
		outFile.Close()

	}
}
