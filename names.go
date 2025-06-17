package main

import (
	"fmt"
	"io"

	"github.com/bjeight/fastats/fasta"
)

// names() is fastats names in the cli. It writes the IDs / descriptions of the records
func names(filepaths []string, args arguments, w io.Writer) error {

	// Write the correct header for the output
	var err error
	if args.description {
		_, err = w.Write([]byte("file\tdescription\n"))
		if err != nil {
			return err
		}
	} else {
		_, err = w.Write([]byte("file\tid\n"))
		if err != nil {
			return err
		}
	}

	// pass namesRecords + the cli arguments to collectCommandLine() for processing the fasta file(s)
	err = applyFastatsFunction(filepaths, namesRecords, args, w)
	if err != nil {
		return err
	}

	return nil
}

// namesRecords does the work of fastats names for one fasta file at a time.
func namesRecords(filename string, r *fasta.Reader, args arguments, w io.Writer) error {

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
		s := fmt.Sprintf("%s\t%s\n", filename, returnRecordName(record, args.description))
		_, err = w.Write([]byte(s))
		if err != nil {
			return err
		}
	}

	return nil
}
