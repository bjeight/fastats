package main

import (
	"fmt"
	"io"

	"github.com/bjeight/fastats/fasta"
)

// length() is fastats len in the cli. It writes the appropriate header (which depends on the cli
// arguments), then passes lengthRecords() + the cli arguments + the writer to collectCommandLine,
// which processes the fasta file(s) from the command line or stdin, depending on what is provided
// by the user.
func length(w io.Writer, filepaths []string, pattern string, file bool, counts bool, description bool) error {

	// print the correct header to stdout, depending on whether the statistics are
	// to be calculated per file or per record
	if file {
		_, err := w.Write([]byte("file\tlength\n"))
		if err != nil {
			return err
		}
	} else {
		_, err := w.Write([]byte("record\tlength\n"))
		if err != nil {
			return err
		}
	}

	// pass lengthRecords + the cli arguments to template() for processesing the fasta file(s)
	err := collectCommandLine(w, lengthRecords, filepaths, pattern, file, counts, description)
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
			s := fmt.Sprintf("%s\t%d\n", return_record_name(record, args.description), len(record.Seq))
			_, err := w.Write([]byte(s))
			if err != nil {
				return err
			}
		}
	}

	// if the statistic is to be calculated per file, we print the total after all
	// the records have been processed
	if args.file {
		s := fmt.Sprintf("%s\t%d\n", filename, l_total)
		_, err := w.Write([]byte(s))
		if err != nil {
			return err
		}
	}

	return nil
}
