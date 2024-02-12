package main

import (
	"bytes"
	"fmt"
	"testing"
)

func Test_length(t *testing.T) {
	out := new(bytes.Buffer)
	err := length(out, []string{}, "", false, false)
	if err != nil {
		t.Error(err)
	}
	if out.String() != `record	length
` {
		fmt.Println(out.String())
		t.Errorf("problem in Test_length")
	}
}

func Test_lengthFile(t *testing.T) {
	out := new(bytes.Buffer)
	err := length(out, []string{}, "", true, false)
	if err != nil {
		t.Error(err)
	}
	if out.String() != `file	length
stdin	0
` {
		fmt.Println(out.String())
		t.Errorf("problem in Test_lengthFile")
	}
}

func Test_lengthCounts(t *testing.T) {
	out := new(bytes.Buffer)
	err := length(out, []string{}, "", false, true)
	if err != nil {
		t.Error(err)
	}
	if out.String() != `record	length
` {
		fmt.Println(out.String())
		t.Errorf("problem in Test_lengthCounts")
	}
}

func Test_lengthFileCounts(t *testing.T) {
	out := new(bytes.Buffer)
	err := length(out, []string{}, "", true, true)
	if err != nil {
		t.Error(err)
	}
	if out.String() != `file	length
stdin	0
` {
		fmt.Println(out.String())
		t.Errorf("problem in Test_lengthFileCounts")
	}
}

func Test_lengthRecords(t *testing.T) {
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

	lengthRecords(
		r,
		arguments{
			filepath: "/path/to/myfile.fasta",
			file:     false,
			counts:   false,
			pattern:  "",
		},
		out,
	)

	desiredResult := `seq1	3
seq2	19
seq3	4
`

	if out.String() != desiredResult {
		fmt.Println(out.String())
		t.Errorf("problem in Test_lengthRecords")
	}
}

func Test_lengthRecordsFile(t *testing.T) {
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

	lengthRecords(
		r,
		arguments{
			filepath: "/path/to/myfile.fasta",
			file:     true,
			counts:   false,
			pattern:  "",
		},
		out,
	)

	desiredResult := `myfile.fasta	26
`

	if out.String() != desiredResult {
		fmt.Println(out.String())
		t.Errorf("problem in Test_lengthRecordsFile")
	}
}
