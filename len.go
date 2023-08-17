package main

import (
	"fmt"
	"io"
	"os"
)

func length(infile string) error {
	f, err := os.Open(infile)
	if err != nil {
		return (err)
	}
	defer f.Close()

	r := NewReader(f)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return (err)
		}

		fmt.Printf("%s length: %d\n", record.ID, len(record.Seq))
	}

	return nil
}
