package main

import (
	"fmt"
	"io"
	"os"
)

func gap(infile string) error {
	f, err := os.Open(infile)
	if err != nil {
		return (err)
	}
	defer f.Close()

	r := NewReader(f)
	for {
		var lookup [256]int
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return (err)
		}
		for _, nuc := range record.Seq {
			lookup[nuc] += 1
		}
		d := 0
		for _, v := range lookup {
			d += v
		}
		n := lookup['-']
		f := float64(n) / float64(d)

		fmt.Printf("%s gap content: %f\n", record.ID, f)
	}

	return nil
}
