package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/bjeight/fastats/fasta"
)

func Test_pattern(t *testing.T) {
	out := new(bytes.Buffer)
	err := pattern(out, []string{}, "ATat", false, false, false)
	if err != nil {
		t.Error(err)
	}
	if out.String() != `record	ATat_prop
` {
		fmt.Println(out.String())
		t.Errorf("problem in Test_pattern")
	}
}

func Test_patternFile(t *testing.T) {
	out := new(bytes.Buffer)
	err := pattern(out, []string{}, "ATat", true, false, false)
	if err != nil {
		t.Error(err)
	}
	if out.String() != `file	ATat_prop
stdin	NaN
` {
		fmt.Println(out.String())
		t.Errorf("problem in Test_patternFile")
	}
}

func Test_patternCounts(t *testing.T) {
	out := new(bytes.Buffer)
	err := pattern(out, []string{}, "ATat", false, true, false)
	if err != nil {
		t.Error(err)
	}
	if out.String() != `record	ATat_count
` {
		fmt.Println(out.String())
		t.Errorf("problem in Test_patternCounts")
	}
}

func Test_patternFileCounts(t *testing.T) {
	out := new(bytes.Buffer)
	err := pattern(out, []string{}, "ATat", true, true, false)
	if err != nil {
		t.Error(err)
	}
	if out.String() != `file	ATat_count
stdin	0
` {
		fmt.Println(out.String())
		t.Errorf("problem in Test_patternFileCounts")
	}
}

func Test_patternRecordsAT(t *testing.T) {
	fastaData := []byte(
		`>seq1
ATGC
>seq2
ATG-ATG-
ATGCATGC
ATGC
>seq3
AT
`)
	fastaR := bytes.NewReader(fastaData)
	r := fasta.NewReader(fastaR)
	out := new(bytes.Buffer)

	patternRecords(
		r,
		arguments{
			filepath:    "/path/to/myfile.fasta",
			file:        false,
			counts:      false,
			description: false,
			pattern:     "ATat",
		},
		out,
	)

	desiredResult := `seq1	0.500000
seq2	0.500000
seq3	1.000000
`

	if out.String() != desiredResult {
		fmt.Println(out.String())
		t.Errorf("problem in Test_patternRecordsAT")
	}
}

func Test_patternRecordsGC(t *testing.T) {
	fastaData := []byte(
		`>seq1
ATGC
>seq2
ATG-ATG-
ATGCATGC
ATGC
>seq3
AT
`)
	fastaR := bytes.NewReader(fastaData)
	r := fasta.NewReader(fastaR)
	out := new(bytes.Buffer)

	patternRecords(
		r,
		arguments{
			filepath:    "/path/to/myfile.fasta",
			file:        false,
			counts:      false,
			description: false,
			pattern:     "GCgc",
		},
		out,
	)

	desiredResult := `seq1	0.500000
seq2	0.400000
seq3	0.000000
`

	if out.String() != desiredResult {
		fmt.Println(out.String())
		t.Errorf("problem in Test_patternRecordsGC")
	}
}

func Test_patternRecordsATFile(t *testing.T) {
	fastaData := []byte(
		`>seq1
ATGC
>seq2
ATG-ATG-
ATGCATGC
ATGC
>seq3
ATGC
`)
	fastaR := bytes.NewReader(fastaData)
	r := fasta.NewReader(fastaR)
	out := new(bytes.Buffer)

	patternRecords(
		r,
		arguments{
			filepath:    "/path/to/myfile.fasta",
			file:        true,
			counts:      false,
			description: false,
			pattern:     "ATat",
		},
		out,
	)

	desiredResult := `myfile.fasta	0.500000
`

	if out.String() != desiredResult {
		fmt.Println(out.String())
		t.Errorf("problem in Test_patternRecordsATFile")
	}
}

func Test_patternRecordsATCount(t *testing.T) {
	fastaData := []byte(
		`>seq1
ATGC
>seq2
ATG-ATG-
ATGCATGC
ATGC
>seq3
AT
`)
	fastaR := bytes.NewReader(fastaData)
	r := fasta.NewReader(fastaR)
	out := new(bytes.Buffer)

	patternRecords(
		r,
		arguments{
			filepath:    "/path/to/myfile.fasta",
			file:        false,
			counts:      true,
			description: false,
			pattern:     "ATat",
		},
		out,
	)

	desiredResult := `seq1	2
seq2	10
seq3	2
`

	if out.String() != desiredResult {
		fmt.Println(out.String())
		t.Errorf("problem in Test_patternRecordsATCount")
	}
}

func Test_patternRecordsATFileCount(t *testing.T) {
	fastaData := []byte(
		`>seq1
ATGC
>seq2
ATG-ATG-
ATGCATGC
ATGC
>seq3
AT
`)
	fastaR := bytes.NewReader(fastaData)
	r := fasta.NewReader(fastaR)
	out := new(bytes.Buffer)

	patternRecords(
		r,
		arguments{
			filepath:    "/path/to/myfile.fasta",
			file:        true,
			counts:      true,
			description: false,
			pattern:     "ATat",
		},
		out,
	)

	desiredResult := `myfile.fasta	14
`

	if out.String() != desiredResult {
		fmt.Println(out.String())
		t.Errorf("problem in Test_patternRecordsATFileCount")
	}
}
