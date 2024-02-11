package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// num() is fastats num in the cli. It prints the appropriate header, then passes
// numRecords() + the cli arguments to template, which processes the fasta file(s)
// from the command line or stdin, depending on what is provided by the user.
func num(filepaths []string, pattern string, file bool, counts bool) error {

	fmt.Println("file\tn_records")

	err := template(numRecords, filepaths, pattern, file, counts)
	if err != nil {
		return err
	}

	return nil
}

// numRecords does the work of fastats len for one fasta file at a time.
func numRecords(args arguments) error {

	// open stdin or a file
	var r *Reader
	if args.filepath == "stdin" {
		r = NewReader(os.Stdin)
	} else {
		f, err := os.Open(args.filepath)
		if err != nil {
			return err
		}
		defer f.Close()
		// depending on whwether the fasta file is compressed or not, provide the correct reader
		switch filepath.Ext(args.filepath) {
		case ".gz", ".bgz":
			r = NewZReader(f)
		default:
			r = NewReader(f)
		}
	}

	// get the file name for when we need to print it
	filename := filenameFromFullPath(args.filepath)

	// initiate a count for the number of records
	c_total := 0

	for {
		_, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		// for every record, += the count
		c_total += 1
	}
	// print the count
	fmt.Printf("%s\t%d\n", filename, c_total)

	return nil
}
