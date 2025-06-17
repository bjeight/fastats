package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/bjeight/fastats/fasta"
)

func Test_length(t *testing.T) {
	out := new(bytes.Buffer)
	err := length([]string{}, arguments{
		file: false, counts: false, description: false, filenames: false,
	}, out)
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
	err := length([]string{}, arguments{
		file: true, counts: false, description: false, filenames: false,
	}, out)
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

func Test_lengthFilenames(t *testing.T) {
	out := new(bytes.Buffer)
	err := length([]string{}, arguments{
		file: false, counts: false, description: false, filenames: true,
	}, out)
	if err != nil {
		t.Error(err)
	}
	if out.String() != `file	record	length
` {
		fmt.Println(out.String())
		t.Errorf("problem in Test_lengthFilenames")
	}
}

func Test_lengthFileFilenames(t *testing.T) {
	out := new(bytes.Buffer)
	err := length([]string{}, arguments{
		file: true, counts: false, description: false, filenames: true,
	}, out)
	if err != nil {
		t.Error(err)
	}
	if out.String() != `file	length
stdin	0
` {
		fmt.Println(out.String())
		t.Errorf("problem in Test_lengthFileFilenames")
	}
}

func Test_lengthCounts(t *testing.T) {
	out := new(bytes.Buffer)
	err := length([]string{}, arguments{
		file: false, counts: true, description: false, filenames: false,
	}, out)
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
	err := length([]string{}, arguments{
		file: true, counts: true, description: false, filenames: false,
	}, out)
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
	r := fasta.NewReader(fastaR)
	out := new(bytes.Buffer)

	lengthRecords(
		"myfile.fasta",
		r,
		arguments{
			file:        false,
			counts:      false,
			description: false,
			pattern:     "",
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
	r := fasta.NewReader(fastaR)
	out := new(bytes.Buffer)

	lengthRecords(
		"myfile.fasta",
		r,
		arguments{
			file:        true,
			counts:      false,
			description: false,
			pattern:     "",
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
