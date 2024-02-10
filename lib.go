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

func template(function func(arguments) error, filepaths []string, pattern string, file bool, count bool) error {

	for _, fp := range filepaths {
		a := arguments{
			filepath: fp,
			file:     file,
			counts:   count,
			pattern:  pattern,
		}
		err := function(a)
		if err != nil {
			return err
		}
	}

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

func filenameFromFullPath(filepath string) string {
	sa := strings.Split(filepath, "/")
	return sa[len(sa)-1]
}
