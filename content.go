package main

import (
	"errors"
	"io"
	"strconv"

	"github.com/bjeight/fastats/fasta"
)

type content struct {
	inputs            []string  // the list of files provided on the command line
	perFile           bool      // calculate stats per file
	writeDescriptions bool      // write record descriptions (else write record ids)
	writeFileNames    bool      // write a column with filename
	patterns          []pattern // arbitrary base content to apply the content functionality to
}

type pattern struct {
	stat         string // {prop, count} - proportions or counts
	bases        string // arbitrary base content to apply the content functionality to
	headerPrefix string // header prefix for the output column
	inverse      bool   // count bases that are NOT the given content
}

func (p *pattern) header() string {
	s := ""
	if p.inverse {
		s += "!"
	}
	if len(p.headerPrefix) == 0 {
		s += p.bases + "_" + p.stat
	} else {
		s += p.headerPrefix + "_" + p.stat
	}
	return s
}

func (args content) writeHeader(w io.Writer) error {
	if args.perFile || args.writeFileNames {
		_, err := w.Write([]byte("file"))
		if err != nil {
			return err
		}
	}
	if !args.perFile {
		if args.writeFileNames {
			_, err := w.Write([]byte("\t"))
			if err != nil {
				return err
			}
		}
		_, err := w.Write([]byte("record"))
		if err != nil {
			return err
		}
	}
	for _, p := range args.patterns {
		_, err := w.Write([]byte("\t" + p.header()))
		if err != nil {
			return err
		}
	}
	_, err := w.Write([]byte("\n"))
	if err != nil {
		return err
	}

	return nil
}

func (args content) writeBody(w io.Writer) error {
	for _, input := range args.inputs {
		reader, file, err := getReaderFile(input)
		if err != nil {
			return err
		}
		defer file.Close()
		err = contentRecords(input, reader, args, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func contentRecords(inputPath string, r *fasta.Reader, args content, w io.Writer) error {

	// initiate a count for the total length of the file
	l_total := 0

	// initiate counts for each pattern
	n_totals := make([]int, len(args.patterns))
	for i := range args.patterns {
		n_totals[i] = 0
	}

	// iterate over every record in the fasta file
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		l_total += len(record.Seq)
		// initiate a table of counts
		var lookup [256]int
		// for every nucleotide in the sequence, +1 its cell in the lookup table
		for _, nuc := range record.Seq {
			lookup[nuc] += 1
		}
		if args.writeFileNames && !args.perFile {
			_, err = w.Write([]byte(returnFileName(inputPath) + "\t"))
			if err != nil {
				return err
			}
		}
		if !args.perFile {
			_, err = w.Write([]byte(returnRecordName(record, args.writeDescriptions)))
			if err != nil {
				return err
			}
		}

		// iterate over all the patterns to be counted
		for i, p := range args.patterns {
			n := 0
			for _, b := range []byte(p.bases) {
				n += lookup[b]
			}
			if p.inverse {
				n = len(record.Seq) - n
			}

			// if the statistic is to be calculated per file, add this record's content count
			// and length to the total, else write this records statistic.
			if args.perFile {
				n_totals[i] += n
			} else {
				// print a count or a proportion
				switch p.stat {
				case "count":
					_, err = w.Write([]byte("\t" + strconv.Itoa(n)))
					if err != nil {
						return err
					}
				case "prop":
					proportion := float64(n) / float64(len(record.Seq))
					_, err = w.Write([]byte("\t" + strconv.FormatFloat(proportion, 'f', 6, 64)))
					if err != nil {
						return err
					}
				default:
					return errors.New("unknown content statistic")
				}
			}
		}
		if !args.perFile {
			_, err = w.Write([]byte("\n"))
			if err != nil {
				return err
			}
		}
	}

	if args.perFile {
		_, err := w.Write([]byte(returnFileName(inputPath)))
		if err != nil {
			return err
		}
		for i, p := range args.patterns {
			switch p.stat {
			case "count":
				_, err := w.Write([]byte("\t" + strconv.Itoa(n_totals[i])))
				if err != nil {
					return err
				}
			case "prop":
				proportion := float64(n_totals[i]) / float64(l_total)
				_, err := w.Write([]byte("\t" + strconv.FormatFloat(proportion, 'f', 6, 64)))
				if err != nil {
					return err
				}
			default:
				return errors.New("unknown content statistic")
			}
		}
		_, err = w.Write([]byte("\n"))
		if err != nil {
			return err
		}
	}

	return nil
}
