package main

import (
	"fmt"
	"io"

	"github.com/bjeight/fastats/fasta"
)

// num() is fastats num in the cli. It writes the header, then passes numRecords() + the
// cli arguments + the writer to collectCommandLine, which processes the fasta file(s)
// from the command line or stdin, depending on what is provided by the user.
func num(filepaths []string, args arguments, w io.Writer) error {

	// Write the correct header for the output
	_, err := w.Write([]byte("file\tn_records\n"))
	if err != nil {
		return err
	}

	// pass numRecords + the cli arguments to collectCommandLine() for processing the fasta file(s)
	err = applyFastatsFunction(filepaths, numRecords, args, w)
	if err != nil {
		return err
	}

	return nil
}

// numRecords does the work of fastats num for one fasta file at a time.
func numRecords(filename string, r *fasta.Reader, args arguments, w io.Writer) error {

	// initiate a count for the number of records
	c_total := 0

	// iterate over every record in the fasta file
	for {
		_, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		// for every record, +1 the count
		c_total += 1
	}
	// print the count
	s := fmt.Sprintf("%s\t%d\n", filename, c_total)
	_, err := w.Write([]byte(s))
	if err != nil {
		return err
	}

	return nil
}
