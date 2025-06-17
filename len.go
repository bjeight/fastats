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
func length(w io.Writer, filepaths []string, pattern string, file bool, counts bool, description bool, filenames bool, lenFormat string) error {

	// write the correct header, depending on whether the statistics are
	// to be calculated per file or per record...
	if file {
		_, err := w.Write([]byte("file\tlength"))
		if err != nil {
			return err
		}
	} else if !filenames {
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
	switch lenFormat {
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

	// pass lengthRecords + the cli arguments to collectCommandLine() for processing the fasta file(s)
	err := collectCommandLine(w, lengthRecords, filepaths, pattern, file, counts, description, filenames, lenFormat)
	if err != nil {
		return err
	}

	return nil
}

// lengthRecords does the work of fastats len for one fasta file at a time.
func lengthRecords(r *fasta.Reader, args arguments, w io.Writer) error {

	// get the file name in case we need to print it to stdout
	filename := filenameFromFullPath(args.filepath)

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
