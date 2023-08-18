package main

import (
	"fmt"
	"io"
	"os"
)

func length(infile string, records bool, counts bool) error {
	f, err := os.Open(infile)
	if err != nil {
		return (err)
	}
	defer f.Close()

	r := NewReader(f)

	fmt.Println("record\tlength")

	l_total := 0

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return (err)
		}
		if records {
			fmt.Printf("%s\t%d\n", record.ID, len(record.Seq))
		} else {
			l_total += len(record.Seq)
		}
	}

	if !records {
		fmt.Printf("%s\t%d\n", parseInfile(infile), l_total)
	}

	return nil
}
