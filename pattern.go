package main

import (
	"fmt"
	"io"
)

// pattern() is fastats at, gc, gaps etc. in the cli. It writes the appropriate header (which
// depends on the cli arguments), then passes lengthRecords() + the cli arguments + the writer to
// collectCommandLine which processes the fasta file(s) from the command line or stdin, depending
// on what is provided by the user.
func pattern(w io.Writer, filepaths []string, pattern string, file bool, counts bool) error {

	switch {
	case file && counts:
		_, err := w.Write([]byte("file\t" + pattern + "_count\n"))
		if err != nil {
			return err
		}
	case file && !counts:
		_, err := w.Write([]byte("file\t" + pattern + "_prop\n"))
		if err != nil {
			return err
		}
	case !file && counts:
		_, err := w.Write([]byte("record\t" + pattern + "_count\n"))
		if err != nil {
			return err
		}
	case !file && !counts:
		_, err := w.Write([]byte("record\t" + pattern + "_prop\n"))
		if err != nil {
			return err
		}
	}

	err := collectCommandLine(w, patternRecords, filepaths, pattern, file, counts)
	if err != nil {
		return err
	}

	return nil
}

// lengthRecords does the work of fastats at, gc, etc. for one fasta file at a time.
func patternRecords(r *Reader, args arguments, w io.Writer) error {

	// get the file name in case we need to print it to stdout
	filename := filenameFromFullPath(args.filepath)

	// we need the pattern to be counted as a slice of bytes so that we can perform
	// the array lookup in the next step
	pattern_slice := []byte(args.pattern)

	// initiate counts for the number of occurences of the specified patterh, and
	// the length of each record
	n_total := 0
	l_total := 0

	// iterate over every record in the fasta file
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
		// and length to the total, else write this records statistic.
		if args.file {
			n_total += n
			l_total += len(record.Seq)
		} else {
			// print a count or a proportion
			if args.counts {
				s := fmt.Sprintf("%s\t%d\n", record.ID, n)
				_, err := w.Write([]byte(s))
				if err != nil {
					return err
				}
			} else {
				proportion := float64(n) / float64(len(record.Seq))
				s := fmt.Sprintf("%s\t%f\n", record.ID, proportion)
				_, err := w.Write([]byte(s))
				if err != nil {
					return err
				}
			}
		}
	}

	// if the statistic is to be calculated per file, we print the statistic after all
	// the records have been processed
	if args.file {
		if args.counts {
			s := fmt.Sprintf("%s\t%d\n", filename, n_total)
			_, err := w.Write([]byte(s))
			if err != nil {
				return err
			}
		} else {
			proportion := float64(n_total) / float64(l_total)
			s := fmt.Sprintf("%s\t%f\n", filename, proportion)
			_, err := w.Write([]byte(s))
			if err != nil {
				return err
			}
		}
	}

	return nil
}
