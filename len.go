package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func length(infiles []string, file bool) error {

	if file {
		fmt.Println("file\tlength")
	} else {
		fmt.Println("record\tlength")
	}

	if len(infiles) > 0 {
		for _, infile := range infiles {
			err := lenFile(infile, file)
			if err != nil {
				return err
			}
		}
	} else {
		f := os.Stdin
		r := NewReader(f)
		filename := "stdin"
		err := lenRecords(r, filename, file)
		if err != nil {
			return err
		}
	}

	return nil
}

func lenFile(infile string, file bool) error {
	f, err := os.Open(infile)
	if err != nil {
		return (err)
	}
	defer f.Close()

	filename := parseInfile(infile)

	var r *Reader
	switch filepath.Ext(infile) {
	case ".gz", ".bgz":
		r = NewZReader(f)
	default:
		r = NewReader(f)
	}

	err = lenRecords(r, filename, file)
	if err != nil {
		return err
	}

	return nil
}

func lenRecords(r *Reader, filename string, file bool) error {
	l_total := 0

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return (err)
		}
		if file {
			l_total += len(record.Seq)
		} else {
			fmt.Printf("%s\t%d\n", record.ID, len(record.Seq))
		}
	}

	if file {
		fmt.Printf("%s\t%d\n", filename, l_total)
	}

	return nil
}
