package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func pattern(infile, pattern string, records bool, counts bool) error {

	ba := parsePattern(pattern)

	f, err := os.Open(infile)
	if err != nil {
		return (err)
	}
	defer f.Close()

	r := NewReader(f)

	if counts {
		fmt.Println("record\t" + pattern + "_count")
	} else {
		fmt.Println("record\t" + pattern + "_prop")
	}

	n_total := 0
	d_total := 0

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return (err)
		}
		var lookup [256]int
		for _, nuc := range record.Seq {
			lookup[nuc] += 1
		}
		n := 0
		for _, b := range ba {
			n += lookup[b]
		}

		if records {
			if counts {
				fmt.Printf("%s\t%d\n", record.ID, n)
			} else {
				stat := float64(n) / float64(len(record.Seq))
				fmt.Printf("%s\t%f\n", record.ID, stat)
			}
		} else {
			n_total += n
			d_total += len(record.Seq)
		}
	}

	if !records {
		if counts {
			fmt.Printf("%s\t%d\n", parseInfile(infile), n_total)
		} else {
			stat := float64(n_total) / float64(d_total)
			fmt.Printf("%s\t%f\n", parseInfile(infile), stat)
		}

	}

	return nil
}

func parsePattern(pattern string) []byte {
	ba := make([]byte, 0)
	for i := range pattern {
		ba = append(ba, pattern[i])
	}
	return ba
}

func parseInfile(infile string) string {
	sa := strings.Split(infile, "/")
	s := sa[(len(sa) - 1)]
	return s
}
