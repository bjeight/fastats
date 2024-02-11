package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// pattern() is fastats at, gc, atgc etc. in the cli. It prints the appropriate header,
// depending on the cli aguments, then passes patternRecords() + the cli arguments to template,
// which processes the fasta file(s) from the command line or stdin, depending on what
//
//	is provided by the user.
func pattern(filepaths []string, pattern string, file bool, counts bool) error {

	switch {
	case file && counts:
		fmt.Println("file\t" + pattern + "_count")
	case file && !counts:
		fmt.Println("file\t" + pattern + "_prop")
	case !file && counts:
		fmt.Println("record\t" + pattern + "_count")
	case !file && !counts:
		fmt.Println("record\t" + pattern + "_prop")
	}

	err := template(patternRecords, filepaths, pattern, file, counts)
	if err != nil {
		return err
	}

	return nil
}

// lengthRecords does the work of fastats at, gc, etc. for one fasta file at a time.
func patternRecords(args arguments) error {

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

	// we need the pattern to be counted as a slice of bytes so that we can perform
	// the array lookup in the next step
	pattern_slice := []byte(args.pattern)

	// initiate counts for the number of occurences of the specified patterh, and
	// the length of each record
	n_total := 0
	l_total := 0

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		// initiate a table of counts
		var lookup [256]int
		// for every nucleotide in the sequence, +1 its cell in the lookup table
		for _, nuc := range record.Seq {
			lookup[nuc] += 1
		}
		// for every nucleotide to be looked up, add its count from the lookup table
		// to the total
		n := 0
		for _, b := range pattern_slice {
			n += lookup[b]
		}

		// if the statistic is to be calculated per file, add this record's pattern count
		// and length to the total, else print this records statistic.
		if args.file {
			n_total += n
			l_total += len(record.Seq)
		} else {
			// print a count or a proportion
			if args.counts {
				fmt.Printf("%s\t%d\n", record.ID, n)
			} else {
				proportion := float64(n) / float64(len(record.Seq))
				fmt.Printf("%s\t%f\n", record.ID, proportion)
			}
		}
	}

	// if the statistic is to be calculated per file, we print the statistic after all
	// the records have been processed
	if args.file {
		if args.counts {
			fmt.Printf("%s\t%d\n", filename, n_total)
		} else {
			proportion := float64(n_total) / float64(l_total)
			fmt.Printf("%s\t%f\n", filename, proportion)
		}
	}

	return nil
}
