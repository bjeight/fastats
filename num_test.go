package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/bjeight/fastats/fasta"
)

func Test_num(t *testing.T) {
	out := new(bytes.Buffer)
	err := num([]string{"stdin"}, arguments{
		file: false, counts: false, description: false, filenames: false,
	}, out)
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
	err := num([]string{"stdin"}, arguments{
		file: true, counts: false, description: false, filenames: false,
	}, out)
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

func Test_numFilenames(t *testing.T) {
	out := new(bytes.Buffer)
	err := num([]string{}, arguments{
		file: false, counts: false, description: false, filenames: true,
	}, out)
	if err != nil {
		t.Error(err)
	}
	if out.String() != `file	n_records
stdin	0
` {
		fmt.Println(out.String())
		t.Errorf("problem in Test_numFilenames")
	}
}

func Test_numFileFilenames(t *testing.T) {
	out := new(bytes.Buffer)
	err := num([]string{}, arguments{
		file: true, counts: false, description: false, filenames: true,
	}, out)
	if err != nil {
		t.Error(err)
	}
	if out.String() != `file	n_records
stdin	0
` {
		fmt.Println(out.String())
		t.Errorf("problem in Test_numFileFilenames")
	}
}

func Test_numCounts(t *testing.T) {
	out := new(bytes.Buffer)
	err := num([]string{}, arguments{
		file: false, counts: true, description: false, filenames: false,
	}, out)
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
	err := num([]string{}, arguments{
		file: true, counts: true, description: false, filenames: false,
	}, out)
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
	r := fasta.NewReader(fastaR)
	out := new(bytes.Buffer)

	numRecords(
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

	desiredResult := `myfile.fasta	3
`

	if out.String() != desiredResult {
		fmt.Println(out.String())
		t.Errorf("problem in Test_numRecords")
	}
}
