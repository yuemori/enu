package enu

import (
	"bufio"
	"os"
)

func FromFile(path string) *Enumerable[string] {
	return New[string](NewFileReader(path))
}

type FileReader struct {
	path    string
	f       *os.File
	err     error
	scanner *bufio.Scanner
}

func NewFileReader(path string) *FileReader {
	return &FileReader{path: path}
}

func (r *FileReader) Err() error {
	return r.err
}

func (r *FileReader) Dispose() {
	if r.f != nil {
		err := r.f.Close()
		if err != nil {
			r.err = err
		}
	}
	r.scanner = nil
}

func (r *FileReader) Next() (string, bool) {
	if r.scanner == nil {
		f, err := os.Open(r.path)
		if err != nil {
			r.err = err
			return "", false
		}
		r.f = f
		r.scanner = bufio.NewScanner(f)
	}

	ok := r.scanner.Scan()
	if !ok {
		return "", false
	}
	return r.scanner.Text(), true
}
