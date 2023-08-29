package main

import (
	"fmt"
	"io"
	"os"
)

func length(infiles []string, file bool, counts bool) error {

	fmt.Println("record\tlength")

	for _, infile := range infiles {

		f, err := os.Open(infile)
		if err != nil {
			return (err)
		}
		defer f.Close()

		r := NewReader(f)

		l_total := 0

		for {
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				return (err)
			}
			if file {
				l_total += len(record.Seq)
			} else {
				fmt.Printf("%s\t%d\n", record.ID, len(record.Seq))
			}
		}

		if file {
			fmt.Printf("%s\t%d\n", parseInfile(infile), l_total)
		}
	}

	return nil
}
