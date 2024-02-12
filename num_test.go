package main

import (
	"bytes"
	"fmt"
	"testing"
)

func Test_num(t *testing.T) {
	out := new(bytes.Buffer)
	err := num(out, []string{}, "", false, false)
	if err != nil {
		t.Error(err)
	}
	if out.String() != `file	n_records
stdin	0
` {
		fmt.Println(out.String())
		t.Errorf("problem in Test_num")
	}
}

func Test_numFile(t *testing.T) {
	out := new(bytes.Buffer)
	err := num(out, []string{}, "", true, false)
	if err != nil {
		t.Error(err)
	}
	if out.String() != `file	n_records
stdin	0
` {
		fmt.Println(out.String())
		t.Errorf("problem in Test_numFile")
	}
}

func Test_numCounts(t *testing.T) {
	out := new(bytes.Buffer)
	err := num(out, []string{}, "", false, true)
	if err != nil {
		t.Error(err)
	}
	if out.String() != `file	n_records
stdin	0
` {
		fmt.Println(out.String())
		t.Errorf("problem in Test_numCounts")
	}
}

func Test_numFileCounts(t *testing.T) {
	out := new(bytes.Buffer)
	err := num(out, []string{}, "", true, true)
	if err != nil {
		t.Error(err)
	}
	if out.String() != `file	n_records
stdin	0
` {
		fmt.Println(out.String())
		t.Errorf("problem in Test_numFileCounts")
	}
}

func Test_numRecords(t *testing.T) {
	fastaData := []byte(
		`>seq1
ATG
>seq2
ATG-ATG-
ATGCATGC
ATG
>seq3
ATGN
`)
	fastaR := bytes.NewReader(fastaData)
	r := NewReader(fastaR)
	out := new(bytes.Buffer)

	numRecords(
		r,
		arguments{
			filepath: "/path/to/myfile.fasta",
			file:     false,
			counts:   false,
			pattern:  "",
		},
		out,
	)

	desiredResult := `myfile.fasta	3
`

	if out.String() != desiredResult {
		fmt.Println(out.String())
		t.Errorf("problem in Test_numRecords")
	}
}
