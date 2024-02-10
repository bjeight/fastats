package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func num(filepaths []string, pattern string, file bool, counts bool) error {

	fmt.Println("file\tn_records")

	err := template(numRecords, filepaths, pattern, file, counts)
	if err != nil {
		return err
	}

	return nil
}

func numRecords(args arguments) error {

	var r *Reader
	if args.filepath == "stdin" {
		r = NewReader(os.Stdin)
	} else {
		f, err := os.Open(args.filepath)
		if err != nil {
			return err
		}
		defer f.Close()
		switch filepath.Ext(args.filepath) {
		case ".gz", ".bgz":
			r = NewZReader(f)
		default:
			r = NewReader(f)
		}
	}

	filename := filenameFromFullPath(args.filepath)

	c_total := 0

	for {
		_, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		c_total += 1
	}
	fmt.Printf("%s\t%d\n", filename, c_total)

	return nil
}
