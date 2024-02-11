package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// length() is fastats len in the cli. It prints the appropriate header, depending on the cli
// aguments, then passes lengthRecords() + the cli arguments to template, which processes the
// fasta file(s) from the command line or stdin, depending on what is provided by the user.
func length(filepaths []string, pattern string, file bool, counts bool) error {

	// print the correct header to stdout, depending on whether the statistics are
	// to be calculated per file or per record
	if file {
		fmt.Println("file\tlength")
	} else {
		fmt.Println("record\tlength")
	}

	// pass lengthRecords + the cli arguments to template() for processesing the fasta file(s)
	err := template(lengthRecords, filepaths, pattern, file, counts)
	if err != nil {
		return err
	}

	return nil
}

// lengthRecords does the work of fastats len for one fasta file at a time.
func lengthRecords(args arguments) error {

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

	// get the file name in case we need to print it to stdout
	filename := filenameFromFullPath(args.filepath)

	// initiate a count for the length of each record
	l_total := 0

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return (err)
		}
		// if the statistic is to be calculated per file, add this record's length
		// to the total, else just print it.
		if args.file {
			l_total += len(record.Seq)
		} else {
			fmt.Printf("%s\t%d\n", record.ID, len(record.Seq))
		}
	}

	// if the statistic is to be calculated per file, we print the total after all
	// the records have been processed
	if args.file {
		fmt.Printf("%s\t%d\n", filename, l_total)
	}

	return nil
}
