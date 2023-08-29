package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func pattern(infiles []string, pattern string, file bool, counts bool) error {

	ba := parsePattern(pattern)

	if counts {
		fmt.Println("record\t" + pattern + "_count")
	} else {
		fmt.Println("record\t" + pattern + "_prop")
	}

	for _, infile := range infiles {

		f, err := os.Open(infile)
		if err != nil {
			return (err)
		}
		defer f.Close()

		r := NewReader(f)

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

			if file {
				n_total += n
				d_total += len(record.Seq)
			} else {
				if counts {
					fmt.Printf("%s\t%d\n", record.ID, n)
				} else {
					stat := float64(n) / float64(len(record.Seq))
					fmt.Printf("%s\t%f\n", record.ID, stat)
				}
			}
		}

		if file {
			if counts {
				fmt.Printf("%s\t%d\n", parseInfile(infile), n_total)
			} else {
				stat := float64(n_total) / float64(d_total)
				fmt.Printf("%s\t%f\n", parseInfile(infile), stat)
			}
		}
	}

	return nil
}

func parsePattern(pattern string) []byte {
	return []byte(pattern)
}

func parseInfile(infile string) string {
	sa := strings.Split(infile, "/")
	return sa[len(sa)-1]
}
