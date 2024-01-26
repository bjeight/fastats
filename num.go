package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func num(infiles []string) error {

	fmt.Println("file\tn_records")

	for _, infile := range infiles {

		f, err := os.Open(infile)
		if err != nil {
			return (err)
		}
		defer f.Close()

		var r *Reader
		switch filepath.Ext(infile) {
		case ".gz", ".bgz":
			r = NewZReader(f)
		default:
			r = NewReader(f)
		}

		c_total := 0

		for {
			_, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				return (err)
			}

			c_total += 1
		}

		fmt.Printf("%s\t%d\n", parseInfile(infile), c_total)
	}

	return nil
}
