package main

import (
	"fmt"
	"io"

	"github.com/bjeight/fastats/fasta"
)

// names() is fastats names in the cli. It writes the ids / descriptions of the records
func names(w io.Writer, filepaths []string, pattern string, file bool, counts bool, description bool) error {

	var err error
	if description {
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

	err = collectCommandLine(w, namesRecords, filepaths, pattern, file, counts, description)
	if err != nil {
		return err
	}

	return nil
}

func namesRecords(r *fasta.Reader, args arguments, w io.Writer) error {

	// get the file name in case we need to print it to stdout
	filename := filenameFromFullPath(args.filepath)

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
		s := fmt.Sprintf("%s\t%s\n", filename, return_record_name(record, args.description))
		_, err = w.Write([]byte(s))
		if err != nil {
			return err
		}
	}

	return nil
}
