package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func num(infiles []string) error {

	fmt.Println("file\tn_records")

	if len(infiles) > 0 {
		for _, infile := range infiles {
			n, err := numFile(infile)
			if err != nil {
				return err
			}
			fmt.Printf("%s\t%d\n", parseInfile(infile), n)
		}
	} else {
		f := os.Stdin
		r := NewReader(f)
		n, err := numRecords(r)
		if err != nil {
			return err
		}
		fmt.Printf("%s\t%d\n", "stdin", n)
	}

	return nil
}

func numFile(infile string) (int, error) {
	f, err := os.Open(infile)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	var r *Reader
	switch filepath.Ext(infile) {
	case ".gz", ".bgz":
		r = NewZReader(f)
	default:
		r = NewReader(f)
	}

	c_total, err := numRecords(r)

	return c_total, err
}

func numRecords(r *Reader) (int, error) {
	c_total := 0

	for {
		_, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return c_total, err
		}

		c_total += 1
	}

	return c_total, nil
}
