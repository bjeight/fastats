package main

import (
	"fmt"
	"io"

	"github.com/bjeight/fastats/fasta"
)

// pattern() is fastats at, gc, gaps etc. in the cli. It writes the appropriate header (which
// depends on the cli arguments), then passes patternRecords() + the cli arguments + the writer to
// collectCommandLine which processes the fasta file(s) from the command line or stdin, depending
// on what is provided by the user.
func pattern(filepaths []string, args arguments, w io.Writer) error {

	// write the correct header, depending on whether the statistics are
	// to be calculated per record or per file, and whether they are counts
	// or proportions
	if args.file || args.filenames {
		_, err := w.Write([]byte("file\t"))
		if err != nil {
			return err
		}
	}
	if !args.file {
		_, err := w.Write([]byte("record\t"))
		if err != nil {
			return err
		}
	}
	if args.counts {
		_, err := w.Write([]byte(args.pattern + "_count\n"))
		if err != nil {
			return err
		}
	} else {
		_, err := w.Write([]byte(args.pattern + "_prop\n"))
		if err != nil {
			return err
		}
	}

	err := applyFastatsFunction(filepaths, patternRecords, args, w)
	if err != nil {
		return err
	}

	return nil
}

// patternRecords does the work of fastats at, gc, etc. for one fasta file at a time.
func patternRecords(filename string, r *fasta.Reader, args arguments, w io.Writer) error {

	// we need the pattern to be counted as a slice of bytes so that we can perform
	// the array lookup in the next step
	pattern_slice := []byte(args.pattern)

	// initiate counts for the number of occurrences of the specified pattern, and
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
			if args.filenames {
				w.Write([]byte(filename + "\t"))
			}
			// print a count or a proportion
			if args.counts {
				s := fmt.Sprintf("%s\t%d\n", returnRecordName(record, args.description), n)
				_, err := w.Write([]byte(s))
				if err != nil {
					return err
				}
			} else {
				proportion := float64(n) / float64(len(record.Seq))
				s := fmt.Sprintf("%s\t%f\n", returnRecordName(record, args.description), proportion)
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
