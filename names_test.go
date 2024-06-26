package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/bjeight/fastats/fasta"
)

func Test_names(t *testing.T) {
	fastaData := []byte(
		`>seq1 is cool
ATG
>seq2 is cool too
ATG-ATG-
ATGCATGC
ATG
>seq3
ATGN
`)
	fastaR := bytes.NewReader(fastaData)
	r := fasta.NewReader(fastaR)
	out := new(bytes.Buffer)

	namesRecords(
		r,
		arguments{
			filepath:    "/path/to/myfile.fasta",
			file:        false,
			counts:      false,
			description: false,
			pattern:     "",
		},
		out,
	)

	desiredResult := `myfile.fasta	seq1
myfile.fasta	seq2
myfile.fasta	seq3
`

	if out.String() != desiredResult {
		fmt.Println(out.String())
		t.Errorf("problem in Test_names (id)")
	}

	fastaR = bytes.NewReader(fastaData)
	r = fasta.NewReader(fastaR)
	out = new(bytes.Buffer)

	namesRecords(
		r,
		arguments{
			filepath:    "/path/to/myfile.fasta",
			file:        false,
			counts:      false,
			description: true,
			pattern:     "",
		},
		out,
	)

	desiredResult = `myfile.fasta	seq1 is cool
myfile.fasta	seq2 is cool too
myfile.fasta	seq3
`

	if out.String() != desiredResult {
		fmt.Println(out.String())
		t.Errorf("problem in Test_names (description)")
	}
}
