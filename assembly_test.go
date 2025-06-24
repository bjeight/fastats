package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/bjeight/fastats/fasta"
)

var fastaFile = []byte(`>s5
AAAAAA
>s6
AAAAAAA
>s9
AAAAAAAAAA
>s2
AAA
>s3
AAAA
>s4
AAAAA
>s7
AAAAAAAA
>s1
AA
>s8
AAAAAAAAA
`)

func TestAssemblyWriteHeader1(t *testing.T) {
	a := assembly{
		stats: []assemblyStatistic{
			{
				sType:  "N",
				sValue: 50,
			},
		},
	}
	out := bytes.NewBuffer(make([]byte, 0))
	desiredResult := "file\tN50\n"

	err := a.writeHeader(out)
	if err != nil {
		t.Error(err)
	}

	if out.String() != desiredResult {
		fmt.Print(out.String())
		t.Fail()
	}
}

func TestAssemblyWriteHeader2(t *testing.T) {
	a := assembly{
		stats: []assemblyStatistic{
			{
				sType:  "N",
				sValue: 90,
			},
		},
	}
	out := bytes.NewBuffer(make([]byte, 0))
	desiredResult := "file\tN90\n"

	err := a.writeHeader(out)
	if err != nil {
		t.Error(err)
	}

	if out.String() != desiredResult {
		fmt.Print(out.String())
		t.Fail()
	}
}

func TestAssemblyWriteHeader3(t *testing.T) {
	a := assembly{
		stats: []assemblyStatistic{
			{
				sType:  "N",
				sValue: 10,
			},
		},
	}
	out := bytes.NewBuffer(make([]byte, 0))
	desiredResult := "file\tN10\n"

	err := a.writeHeader(out)
	if err != nil {
		t.Error(err)
	}

	if out.String() != desiredResult {
		fmt.Print(out.String())
		t.Fail()
	}
}

func TestN50WriteHeader4(t *testing.T) {
	a := assembly{
		stats: []assemblyStatistic{
			{
				sType:  "NG",
				sValue: 50,
			},
		},
	}
	out := bytes.NewBuffer(make([]byte, 0))
	desiredResult := "file\tNG50\n"

	err := a.writeHeader(out)
	if err != nil {
		t.Error(err)
	}

	if out.String() != desiredResult {
		fmt.Print(out.String())
		t.Fail()
	}
}

func TestAssemblyWriteHeader5(t *testing.T) {
	a := assembly{
		stats: []assemblyStatistic{
			{
				sType:  "NG",
				sValue: 90,
			},
		},
	}
	out := bytes.NewBuffer(make([]byte, 0))
	desiredResult := "file\tNG90\n"

	err := a.writeHeader(out)
	if err != nil {
		t.Error(err)
	}

	if out.String() != desiredResult {
		fmt.Print(out.String())
		t.Fail()
	}
}

func TestAssemblyWriteHeader6(t *testing.T) {
	a := assembly{
		stats: []assemblyStatistic{
			{
				sType:  "N",
				sValue: 50,
			},
			{
				sType:  "L",
				sValue: 50,
			},
			{
				sType:  "NG",
				sValue: 50,
			},
		},
	}
	out := bytes.NewBuffer(make([]byte, 0))
	desiredResult := "file\tN50\tL50\tNG50\n"

	err := a.writeHeader(out)
	if err != nil {
		t.Error(err)
	}

	if out.String() != desiredResult {
		fmt.Print(out.String())
		t.Fail()
	}
}

func TestAssemblyWriteBody1(t *testing.T) {
	a := assembly{
		inputs: []string{"stdin"},
		stats: []assemblyStatistic{
			{
				sType:  "N",
				sValue: 50,
			},
		},
	}
	reader := fasta.NewReader(bytes.NewReader(fastaFile))
	desiredResult := `stdin	8
`
	out := bytes.NewBuffer(make([]byte, 0))

	err := assemblyRecords("stdin", reader, a, out)
	if err != nil {
		t.Error(err)
	}

	if out.String() != desiredResult {
		fmt.Print(out.String())
		t.Fail()
	}
}

func TestAssemblyWriteBody2(t *testing.T) {
	a := assembly{
		inputs: []string{"stdin"},
		stats: []assemblyStatistic{
			{
				sType:  "N",
				sValue: 90,
			},
		},
	}
	reader := fasta.NewReader(bytes.NewReader(fastaFile))
	desiredResult := `stdin	4
` // len -f fastaFile = 54, so n90 = 48.6 bases, 10 + 9 + 8 + 7 + 6 + 5 + 4 = 49, so n90 = 4
	out := bytes.NewBuffer(make([]byte, 0))

	err := assemblyRecords("stdin", reader, a, out)
	if err != nil {
		t.Error(err)
	}

	if out.String() != desiredResult {
		fmt.Print(out.String())
		t.Fail()
	}
}

func TestN50WriteBody3(t *testing.T) {
	a := assembly{
		inputs: []string{"stdin"},
		stats: []assemblyStatistic{
			{
				sType:  "NG",
				sValue: 50,
			},
		},
		genomeSize: 54,
	}
	reader := fasta.NewReader(bytes.NewReader(fastaFile))
	desiredResult := `stdin	8
`
	out := bytes.NewBuffer(make([]byte, 0))

	err := assemblyRecords("stdin", reader, a, out)
	if err != nil {
		t.Error(err)
	}

	if out.String() != desiredResult {
		fmt.Print(out.String())
		t.Fail()
	}
}

func TestN50WriteBody4(t *testing.T) {
	a := assembly{
		inputs: []string{"stdin"},
		stats: []assemblyStatistic{
			{
				sType:  "NG",
				sValue: 50,
			},
		},
		genomeSize: 80,
	}
	reader := fasta.NewReader(bytes.NewReader(fastaFile))
	desiredResult := `stdin	6
` // len -f fastaFile = 54, so n50 = 40 bases, 10 + 9 + 8 + 7 + 6 = 40, so n50 = 6
	out := bytes.NewBuffer(make([]byte, 0))

	err := assemblyRecords("stdin", reader, a, out)
	if err != nil {
		t.Error(err)
	}

	if out.String() != desiredResult {
		fmt.Print(out.String())
		t.Fail()
	}
}

func TestN50WriteBody5(t *testing.T) {
	a := assembly{
		inputs: []string{"stdin"},
		stats: []assemblyStatistic{
			{
				sType:  "N",
				sValue: 50,
			},
			{
				sType:  "L",
				sValue: 50,
			},
			{
				sType:  "NG",
				sValue: 50,
			},
		},
		genomeSize: 54,
	}
	reader := fasta.NewReader(bytes.NewReader(fastaFile))
	desiredResult := `stdin	8	3	8
` // len -f fastaFile = 54, so n50 = 40 bases, 10 + 9 + 8 + 7 + 6 = 40, so n50 = 6
	out := bytes.NewBuffer(make([]byte, 0))

	err := assemblyRecords("stdin", reader, a, out)
	if err != nil {
		t.Error(err)
	}

	if out.String() != desiredResult {
		fmt.Print(out.String())
		t.Fail()
	}
}
