package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/bjeight/fastats/fasta"
)

func getReaderFile(inputPath string) (*fasta.Reader, *os.File, error) {
	// open stdin or a file
	var r *fasta.Reader
	var f *os.File
	if inputPath == "stdin" {
		f := os.Stdin
		r = fasta.NewReader(f)
	} else {
		f, err := os.Open(inputPath)
		if err != nil {
			return r, f, err
		}
		// depending on whether the fasta file is compressed or not, provide the correct reader
		switch filepath.Ext(inputPath) {
		case ".gz", ".bgz":
			r = fasta.NewZReader(f)
		default:
			r = fasta.NewReader(f)
		}
	}

	return r, f, nil
}

// Get just the filename from path + filename or just return the full path
func returnFileName(filepath string) string {
	sa := strings.Split(filepath, "/")
	return sa[len(sa)-1]
}

// Return either fasta record ID or its (full) description
func returnRecordName(record fasta.Record, description bool) string {
	if description {
		return record.Description
	}
	return record.ID
}
