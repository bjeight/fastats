package main

import (
	"fmt"
	"io"

	"github.com/bjeight/fastats/fasta"
)

type length struct {
	inputs            []string
	perFile           bool
	writeDescriptions bool
	writeFileNames    bool
	lenFormat         string
}

func (args length) writeHeader(w io.Writer) error {
	if args.perFile {
		_, err := w.Write([]byte("file\tlength"))
		if err != nil {
			return err
		}
	} else if !args.writeFileNames {
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

	switch args.lenFormat {
	case "kb", "mb", "gb":
		_, err := w.Write([]byte("_" + args.lenFormat + "\n"))
		if err != nil {
			return err
		}
	default:
		_, err := w.Write([]byte("\n"))
		if err != nil {
			return err
		}
	}
	return nil
}

func (args length) writeBody(w io.Writer) error {
	for _, input := range args.inputs {
		reader, file, err := getReaderFile(input)
		if err != nil {
			return err
		}
		defer file.Close()
		err = lengthRecords(input, reader, args, w)
		if err != nil {
			return err
		}
	}
	return nil
}

// lengthRecords does the work of fastats len for one fasta file at a time.
func lengthRecords(inputPath string, r *fasta.Reader, args length, w io.Writer) error {

	// initiate a count for the length of each record
	var l_total int64 = 0

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
		if args.perFile {
			l_total += int64(len(record.Seq))
		} else {
			if args.writeFileNames {
				_, err = w.Write([]byte(returnFileName(inputPath) + "\t"))
				if err != nil {
					return err
				}
			}
			s := fmt.Sprintf("%s\t%s\n", returnRecordName(record, args.writeDescriptions), returnLengthFormatted(int64(len(record.Seq)), args.lenFormat))
			_, err = w.Write([]byte(s))
			if err != nil {
				return err
			}
		}
	}

	// if the statistic is to be calculated per file, we print the total after all
	// the records have been processed
	if args.perFile {
		s := fmt.Sprintf("%s\t%s\n", returnFileName(inputPath), returnLengthFormatted(l_total, args.lenFormat))
		_, err := w.Write([]byte(s))
		if err != nil {
			return err
		}
	}

	return nil
}
