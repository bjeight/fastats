package main

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/bjeight/fastats/fasta"
)

type arguments struct {
	filepath    string
	file        bool
	counts      bool
	description bool
	filenames   bool
	pattern     string
	lenFormat   string
}

type fastatsFunction func(*fasta.Reader, arguments, io.Writer) error

// For every file provided on the command line, collectCommandLine applies the correct functionality based on the cli arguments.
// If no files are provided, it signals that we should try to read an uncompressed fasta file from stdin.
func collectCommandLine(w io.Writer, fn fastatsFunction, filepaths []string, pattern string, file bool, count bool, description bool, filenames bool, lenFormat string) error {

	// for every file provided on the command line...
	for _, fp := range filepaths {
		// wrap the arguments up in a struct
		a := arguments{
			filepath:    fp,
			file:        file,
			counts:      count,
			description: description,
			filenames:   filenames,
			pattern:     pattern,
			lenFormat:   lenFormat,
		}
		// and pass them to the correct function (defined when collectCommandLine is called)
		err := applyFastatsFunction(fn, a, w)
		if err != nil {
			return err
		}
	}

	// If there were no files provided on the command line, attempt to read from stdin
	if len(filepaths) == 0 {
		a := arguments{
			filepath:    "stdin",
			file:        file,
			counts:      count,
			description: description,
			filenames:   filenames,
			pattern:     pattern,
			lenFormat:   lenFormat,
		}
		err := applyFastatsFunction(fn, a, w)
		if err != nil {
			return err
		}
	}

	return nil
}

// Given a function to apply to the fasta file and the other command line arguments,
// open the correct file, create an appropriate reader, and apply the function
func applyFastatsFunction(fn fastatsFunction, args arguments, w io.Writer) error {
	// open stdin or a file
	var r *fasta.Reader
	if args.filepath == "stdin" {
		r = fasta.NewReader(os.Stdin)
	} else {
		f, err := os.Open(args.filepath)
		if err != nil {
			return err
		}
		defer f.Close()
		// depending on whether the fasta file is compressed or not, provide the correct reader
		switch filepath.Ext(args.filepath) {
		case ".gz", ".bgz":
			r = fasta.NewZReader(f)
		default:
			r = fasta.NewReader(f)
		}
	}

	err := fn(r, args, w)
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
