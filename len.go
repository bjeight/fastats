package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func length(infiles []string, pattern string, file bool, counts bool) error {

	if file {
		fmt.Println("file\tlength")
	} else {
		fmt.Println("record\tlength")
	}

	err := template(lengthRecords, infiles, pattern, file, counts)
	if err != nil {
		return err
	}

	return nil
}

func lengthRecords(args arguments) error {

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
		if err != nil {
			return err
		}
	}

	filename := filenameFromFullPath(args.filepath)
	l_total := 0

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return (err)
		}
		if args.file {
			l_total += len(record.Seq)
		} else {
			fmt.Printf("%s\t%d\n", record.ID, len(record.Seq))
		}
	}
	if args.file {
		fmt.Printf("%s\t%d\n", filename, l_total)
	}

	return nil
}
