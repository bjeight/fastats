package main

import (
	"fmt"
	"io"
	"strconv"

	"github.com/bjeight/fastats/fasta"
)

// length() is fastats len in the cli. It writes the appropriate header (which depends on the cli
// arguments), then passes lengthRecords() + the cli arguments + the writer to collectCommandLine,
// which processes the fasta file(s) from the command line or stdin, depending on what is provided
// by the user.
func length(filepaths []string, args arguments, w io.Writer) error {

	// write the correct header, depending on whether the statistics are
	// to be calculated per file or per record...
	if args.file {
		_, err := w.Write([]byte("file\tlength"))
		if err != nil {
			return err
		}
	} else if !args.filenames {
		_, err := w.Write([]byte("record\tlength"))
		if err != nil {
			return err
		}
	} else {
		_, err := w.Write([]byte("file\trecord\tlength"))
		if err != nil {
			return err
		}
	}

	// ...and whether we are printing length in units other than bases
	switch args.lenFormat {
	case "kb":
		_, err := w.Write([]byte("_kb\n"))
		if err != nil {
			return err
		}
	case "mb":
		_, err := w.Write([]byte("_mb\n"))
		if err != nil {
			return err
		}
	case "gb":
		_, err := w.Write([]byte("_gb\n"))
		if err != nil {
			return err
		}
	default:
		_, err := w.Write([]byte("\n"))
		if err != nil {
			return err
		}
	}

	err := applyFastatsFunction(filepaths, lengthRecords, args, w)
	if err != nil {
		return err
	}

	return nil
}

// lengthRecords does the work of fastats len for one fasta file at a time.
func lengthRecords(filename string, r *fasta.Reader, args arguments, w io.Writer) error {

	// initiate a count for the length of each record
	l_total := 0

	// iterate over every record in the fasta file
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return (err)
		}
		// if the statistic is to be calculated per file, add this record's length
		// to the total, else just write it.
		if args.file {
			l_total += len(record.Seq)
		} else {
			if args.filenames {
				w.Write([]byte(filename + "\t"))
			}
			s := fmt.Sprintf("%s\t%s\n", returnRecordName(record, args.description), returnRecordLength(len(record.Seq), args.lenFormat))
			_, err := w.Write([]byte(s))
			if err != nil {
				return err
			}
		}
	}

	// if the statistic is to be calculated per file, we print the total after all
	// the records have been processed
	if args.file {
		s := fmt.Sprintf("%s\t%s\n", filename, returnRecordLength(l_total, args.lenFormat))
		_, err := w.Write([]byte(s))
		if err != nil {
			return err
		}
	}

	return nil
}

// returnRecordLength (potentially) converts bases to kb, mb, gb.
func returnRecordLength(l int, unit string) string {
	var s string
	switch unit {
	case "kb":
		s = strconv.FormatFloat(float64(l)/float64(1000), 'f', 3, 64)
	case "mb":
		s = strconv.FormatFloat(float64(l)/float64(1000000), 'f', 6, 64)
	case "gb":
		s = strconv.FormatFloat(float64(l)/float64(1000000000), 'f', 9, 64)
	default:
		s = strconv.Itoa(l)
	}
	return s
}
