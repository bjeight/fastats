package main

import (
	"fmt"
	"io"
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
		err = numRecords(input, reader, args, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func numRecords(inputPath string, r *Reader, args num, w io.Writer) error {
	// initiate a count for the number of records
	var c_total int64 = 0

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
	s := fmt.Sprintf("%s\t%d\n", returnFileName(inputPath), c_total)
	_, err := w.Write([]byte(s))
	return err
}
