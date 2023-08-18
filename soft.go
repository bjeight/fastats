package main

import (
	"fmt"
	"io"
	"os"
)

func soft(infile string) error {
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
		n := lookup['a'] + lookup['t'] + lookup['g'] + lookup['c'] + lookup['n']
		f := float64(n) / float64(d)

		fmt.Printf("%s soft-masked: %f\n", record.ID, f)
	}

	return nil
}
