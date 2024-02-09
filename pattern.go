package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func pattern(infiles []string, pattern string, file bool, counts bool) error {

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

	err := template(patternRecords, infiles, pattern, file, counts)
	if err != nil {
		return err
	}

	return nil
}

func patternRecords(args arguments) error {

	var r *Reader
	if args.filepath == "stdin" {
		r = NewReader(os.Stdin)
	} else {
		f, err := os.Open(args.filepath)
		if err != nil {
			return err
		}
		defer f.Close()
		switch filepath.Ext(args.filepath) {
		case ".gz", ".bgz":
			r = NewZReader(f)
		default:
			r = NewReader(f)
		}
	}

	filename := filenameFromFullPath(args.filepath)
	pattern_array := parsePattern(args.pattern)
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

		if args.file {
			n_total += n
			d_total += len(record.Seq)
		} else {
			if args.counts {
				fmt.Printf("%s\t%d\n", record.ID, n)
			} else {
				proportion := float64(n) / float64(len(record.Seq))
				fmt.Printf("%s\t%f\n", record.ID, proportion)
			}
		}
	}
	if args.file {
		if args.counts {
			fmt.Printf("%s\t%d\n", filename, n_total)
		} else {
			proportion := float64(n_total) / float64(d_total)
			fmt.Printf("%s\t%f\n", filename, proportion)
		}
	}

	return nil
}

func parsePattern(pattern string) []byte {
	return []byte(pattern)
}
