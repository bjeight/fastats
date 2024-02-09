package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func pattern(infiles []string, pattern string, file bool, counts bool) error {

	pattern_array := parsePattern(pattern)

	switch {
	case file && counts:
		fmt.Println("file\t" + pattern + "_count")
	case file && !counts:
		fmt.Println("file\t" + pattern + "_prop")
	case !file && counts:
		fmt.Println("record\t" + pattern + "_count")
	case !file && !counts:
		fmt.Println("record\t" + pattern + "_prop")
	}

	if len(infiles) > 0 {
		for _, infile := range infiles {
			err := patternFile(infile, pattern_array, file, counts)
			if err != nil {
				return err
			}
		}
	} else {
		f := os.Stdin
		r := NewReader(f)
		filename := "stdin"
		err := patternRecords(r, pattern_array, filename, file, counts)
		if err != nil {
			return err
		}
	}

	return nil
}

func patternFile(infile string, pattern_array []byte, file bool, counts bool) error {
	f, err := os.Open(infile)
	if err != nil {
		return err
	}
	defer f.Close()

	var r *Reader
	switch filepath.Ext(infile) {
	case ".gz", ".bgz":
		r = NewZReader(f)
	default:
		r = NewReader(f)
	}

	filename := parseInfile(infile)

	err = patternRecords(r, pattern_array, filename, file, counts)
	if err != nil {
		return err
	}

	return nil
}

func patternRecords(r *Reader, pattern_array []byte, filename string, file bool, counts bool) error {

	n_total := 0
	d_total := 0

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		var lookup [256]int
		for _, nuc := range record.Seq {
			lookup[nuc] += 1
		}
		n := 0
		for _, b := range pattern_array {
			n += lookup[b]
		}

		if file {
			n_total += n
			d_total += len(record.Seq)
		} else {
			if counts {
				fmt.Printf("%s\t%d\n", record.ID, n)
			} else {
				stat := float64(n) / float64(len(record.Seq))
				fmt.Printf("%s\t%f\n", record.ID, stat)
			}
		}
	}

	if file {
		if counts {
			fmt.Printf("%s\t%d\n", filename, n_total)
		} else {
			stat := float64(n_total) / float64(d_total)
			fmt.Printf("%s\t%f\n", filename, stat)
		}
	}

	return nil
}

func parsePattern(pattern string) []byte {
	return []byte(pattern)
}

func parseInfile(infile string) string {
	sa := strings.Split(infile, "/")
	return sa[len(sa)-1]
}
