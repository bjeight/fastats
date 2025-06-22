package main

import (
	"fmt"
	"io"

	"github.com/bjeight/fastats/fasta"
)

type pattern struct {
	inputs            []string // the list of files provided on the command line
	perFile           bool     // calculate stats per file
	writeCounts       bool     // calculate counts (else calculate proportions)
	writeDescriptions bool     // write record descriptions (else write record ids)
	writeFileNames    bool     // write a column with filename
	bases             string   // arbitrary base content to apply the pattern functionality to
}

func (args pattern) writeHeader(w io.Writer) error {
	if args.perFile || args.writeFileNames {
		_, err := w.Write([]byte("file\t"))
		if err != nil {
			return err
		}
	}
	if !args.perFile {
		_, err := w.Write([]byte("record\t"))
		if err != nil {
			return err
		}
	}
	if args.writeCounts {
		_, err := w.Write([]byte(args.bases + "_count\n"))
		if err != nil {
			return err
		}
	} else {
		_, err := w.Write([]byte(args.bases + "_prop\n"))
		if err != nil {
			return err
		}
	}
	return nil
}

func (args pattern) writeBody(w io.Writer) error {
	for _, input := range args.inputs {
		reader, file, err := getReaderFile(input)
		if err != nil {
			return err
		}
		defer file.Close()
		err = patternRecords(input, reader, args, w)
		if err != nil {
			return err
		}
	}
	return nil
}

// patternRecords does the work of fastats at, gc, etc. for one fasta file at a time.
func patternRecords(inputPath string, r *fasta.Reader, args pattern, w io.Writer) error {

	// we need the pattern to be counted as a slice of bytes so that we can perform
	// the array lookup in the next step
	bases_slice := []byte(args.bases)

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
		for _, b := range bases_slice {
			n += lookup[b]
		}

		// if the statistic is to be calculated per file, add this record's pattern count
		// and length to the total, else write this records statistic.
		if args.perFile {
			n_total += n
			l_total += len(record.Seq)
		} else {
			if args.writeFileNames {
				// output = output + returnFileName(inputPath) + "\t"
				_, err = w.Write([]byte(returnFileName(inputPath) + "\t"))
				if err != nil {
					return err
				}
			}
			// print a count or a proportion
			if args.writeCounts {
				s := fmt.Sprintf("%s\t%d\n", returnRecordName(record, args.writeDescriptions), n)
				_, err = w.Write([]byte(s))
				if err != nil {
					return err
				}
			} else {
				proportion := float64(n) / float64(len(record.Seq))
				s := fmt.Sprintf("%s\t%f\n", returnRecordName(record, args.writeDescriptions), proportion)
				_, err = w.Write([]byte(s))
				if err != nil {
					return err
				}
			}
		}
	}

	// if the statistic is to be calculated per file, we print the statistic after all
	// the records have been processed
	if args.perFile {
		if args.writeCounts {
			s := fmt.Sprintf("%s\t%d\n", returnFileName(inputPath), n_total)
			_, err := w.Write([]byte(s))
			if err != nil {
				return err
			}
		} else {
			proportion := float64(n_total) / float64(l_total)
			s := fmt.Sprintf("%s\t%f\n", returnFileName(inputPath), proportion)
			_, err := w.Write([]byte(s))
			if err != nil {
				return err
			}
		}
	}

	return nil
}
