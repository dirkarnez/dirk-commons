package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

// ReadFile read a file
func ReadFile(path string, onFileRead func(*os.File) error) error {
	file, err := os.Open(path)

	if err != nil {
		return err
	}

	defer file.Close()

	return onFileRead(file)
}

// ReadFileAsString read file as string
func ReadFileAsString(path string) (string, error) {
	buf := new(bytes.Buffer)

	err := ReadFile(path, func(file *os.File) error {
		_, err := buf.ReadFrom(file)
		return err
	})

	if err != nil {
		return "", err
	} else {
		return buf.String(), nil
	}
}

// ReadFileAsLines Read file as lines
func ReadFileAsLines(path string, onEachLine func(string) error) error {
	return ReadFile(path, func(file *os.File) error {
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			err := onEachLine(scanner.Text())
			if err != nil {
				return err
			}
		}

		return scanner.Err()
	})
}

// CreateFile create file
func CreateFile(path string, onFileCreate func(*os.File) error) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	return onFileCreate(file)
}

// WriteLinesToFile write lines to File
func WriteLinesToFile(path string, lines []string) error {
	return CreateFile(path, func(file *os.File) error {
		w := bufio.NewWriter(file)

		for _, line := range lines {
			fmt.Fprintln(w, line)
		}

		return w.Flush()
	})
}

// io.Reader includes io.ReadCloser
// WriteIOReaderToFile create file
func WriteIOReaderToFile(path string, reader io.Reader) error {
	return CreateFile(path, func(file *os.File) error {
		_, err = io.Copy(outFile, reader)
		return err
	})
}
