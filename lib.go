package main

import (
	"strings"
)

type arguments struct {
	filepath string
	file     bool
	counts   bool
	pattern  string
}

// For every file provided on the command line, template applies the correct functionality based on the cli arguments.
// If no files are provided, it attempts to read an uncompressed fasta file from stdin.
func template(function func(arguments) error, filepaths []string, pattern string, file bool, count bool) error {

	// for every file provided on the command line...
	for _, fp := range filepaths {
		// wrap the arguments up in s struct
		a := arguments{
			filepath: fp,
			file:     file,
			counts:   count,
			pattern:  pattern,
		}
		// and pass them to the correct function (defined when template is called)
		err := function(a)
		if err != nil {
			return err
		}
	}

	// if there were no files provided on the command line, attemp to read from stdin
	if len(filepaths) == 0 {
		a := arguments{
			filepath: "stdin",
			file:     file,
			counts:   count,
			pattern:  pattern,
		}
		err := function(a)
		if err != nil {
			return err
		}
	}

	return nil
}

// Get just the filename from the path + filename
func filenameFromFullPath(filepath string) string {
	sa := strings.Split(filepath, "/")
	return sa[len(sa)-1]
}
