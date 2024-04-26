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
	pattern     string
}

type fastatsFunction func(*fasta.Reader, arguments, io.Writer) error

// For every file provided on the command line, template applies the correct functionality based on the cli arguments.
// If no files are provided, it signals that we should try to read an uncompressed fasta file from stdin.
func collectCommandLine(w io.Writer, fn fastatsFunction, filepaths []string, pattern string, file bool, count bool, description bool) error {

	// for every file provided on the command line...
	for _, fp := range filepaths {
		// wrap the arguments up in s struct
		a := arguments{
			filepath:    fp,
			file:        file,
			counts:      count,
			description: description,
			pattern:     pattern,
		}
		// and pass them to the correct function (defined when template is called)
		err := applyFastatsFunction(fn, a, w)
		if err != nil {
			return err
		}
	}

	// if there were no files provided on the command line, attemp to read from stdin
	if len(filepaths) == 0 {
		a := arguments{
			filepath:    "stdin",
			file:        file,
			counts:      count,
			pattern:     pattern,
			description: description,
		}
		err := applyFastatsFunction(fn, a, w)
		if err != nil {
			return err
		}
	}

	return nil
}

// Given a function to apply to the fasta file and the other command line arguments,
// open the correct file, create an appropriate reader, and apply the function (passing
// it the command line arguments and the writer from the scope above)
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
		// depending on whwether the fasta file is compressed or not, provide the correct reader
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

	return err
}

// Get just the filename from the path + filename
func filenameFromFullPath(filepath string) string {
	sa := strings.Split(filepath, "/")
	return sa[len(sa)-1]
}

func return_record_name(fr fasta.FastaRecord, description bool) string {
	if description {
		return fr.Description
	} else {
		return fr.ID
	}
}
