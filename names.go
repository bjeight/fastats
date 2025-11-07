package main

import (
	"fmt"
	"io"

	"github.com/bjeight/fastats/fasta"
)

type names struct {
	inputs            []string
	writeDescriptions bool
}

func (args names) writeHeader(w io.Writer) error {
	var err error
	if args.writeDescriptions {
		_, err = w.Write([]byte("file\tdescription\n"))
	} else {
		_, err = w.Write([]byte("file\tid\n"))
	}
	return err
}

func (args names) writeBody(w io.Writer) error {
	for _, input := range args.inputs {
		reader, file, err := getReaderFile(input)
		if err != nil {
			return err
		}
		defer file.Close()
		err = namesRecords(input, reader, args, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func namesRecords(inputPath string, r *fasta.Reader, args names, w io.Writer) error {

	// iterate over every record in the fasta file
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		// if the statistic is to be calculated per file, add this record's length
		// to the total, else just write it.
		s := fmt.Sprintf("%s\t%s\n", returnFileName(inputPath), returnRecordName(record, args.writeDescriptions))
		_, err = w.Write([]byte(s))
		if err != nil {
			return err
		}
	}

	return nil
}
