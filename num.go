package main

import (
	"fmt"
	"io"

	"github.com/bjeight/fastats/fasta"
)

type num struct {
	inputs []string
}

func (args num) writeHeader(w io.Writer) error {
	_, err := w.Write([]byte("file\tn_records\n"))
	return err
}

func (args num) writeBody(w io.Writer) error {
	for _, input := range args.inputs {
		reader, file, err := getReaderFile(input)
		if err != nil {
			return err
		}
		defer file.Close()
		s, err := numRecords(input, reader, args)
		if err != nil {
			return err
		}
		_, err = w.Write([]byte(s))
		if err != nil {
			return err
		}
	}
	return nil
}

// numRecords does the work of fastats num for one fasta file at a time.
func numRecords(inputPath string, r *fasta.Reader, args num) (string, error) {
	// initiate a count for the number of records
	c_total := 0

	// iterate over every record in the fasta file
	for {
		_, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}
		// for every record, +1 the count
		c_total += 1
	}
	s := fmt.Sprintf("%s\t%d\n", returnFileName(inputPath), c_total)
	return s, nil
}
