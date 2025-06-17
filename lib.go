package main

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/bjeight/fastats/fasta"
)

type arguments struct {
	file        bool
	counts      bool
	description bool
	filenames   bool
	pattern     string
	lenFormat   string
}

type fastatsFunction func(string, *fasta.Reader, arguments, io.Writer) error

// For every file provided on the command line, collectCommandLine applies the correct functionality based on the cli arguments.
// If no files are provided, it signals that we should try to read an uncompressed fasta file from stdin.
func applyFastatsFunction(filepaths []string, fn fastatsFunction, a arguments, w io.Writer) error {

	// for every file provided on the command line...
	for _, fp := range filepaths {
		err := applyFastatsFunctionFile(fp, fn, a, w)
		if err != nil {
			return err
		}
	}

	// If there were no files provided on the command line, attempt to read from stdin
	if len(filepaths) == 0 {
		err := applyFastatsFunctionFile("stdin", fn, a, w)
		if err != nil {
			return err
		}
	}

	return nil
}

// Given a function to apply to the fasta file and the other command line arguments,
// open the correct file, create an appropriate reader, and apply the function
func applyFastatsFunctionFile(fp string, fn fastatsFunction, args arguments, w io.Writer) error {
	// get the file name in case we need to print it to stdout
	filename := filenameFromFullPath(fp)

	// open stdin or a file
	var r *fasta.Reader
	if fp == "stdin" {
		r = fasta.NewReader(os.Stdin)
	} else {
		f, err := os.Open(fp)
		if err != nil {
			return err
		}
		defer f.Close()
		// depending on whether the fasta file is compressed or not, provide the correct reader
		switch filepath.Ext(fp) {
		case ".gz", ".bgz":
			r = fasta.NewZReader(f)
		default:
			r = fasta.NewReader(f)
		}
	}

	err := fn(filename, r, args, w)
	if err != nil {
		return err
	}

	return nil
}

// Get just the filename from path + filename
func filenameFromFullPath(filepath string) string {
	sa := strings.Split(filepath, "/")
	return sa[len(sa)-1]
}

// Return either fasta record ID or its (full) description
func returnRecordName(fr fasta.FastaRecord, description bool) string {
	if description {
		return fr.Description
	} else {
		return fr.ID
	}
}
