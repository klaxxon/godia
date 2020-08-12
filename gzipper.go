package main

import (
	"compress/gzip"
	"io/ioutil"
	"log"
	"os"
)

func CreateGzipFile(fn string, b []byte) {
	f, err := os.Create(fn)
	if err != nil {
		log.Fatal(err)
	}
	w := gzip.NewWriter(f)

	_, err = w.Write(b)
	if err != nil {
		log.Fatal(err)
	}
	if err := w.Close(); err != nil {
		log.Fatal(err)
	}
}

func ReadGzipFile(fn string) ([]byte, error) {
	f, err := os.Open(fn)
	if err != nil {
		return nil, err
	}

	r, err := gzip.NewReader(f)
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	if err := r.Close(); err != nil {
		return nil, err
	}
	return content, nil
}
