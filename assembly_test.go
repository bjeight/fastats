package main

import (
	"bytes"
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
	output := bytes.NewBuffer(make([]byte, 0))
	expected := "file\tN50\n"

	err := a.writeHeader(output)
	if err != nil {
		t.Error(err)
	}

	if output.String() != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
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
	output := bytes.NewBuffer(make([]byte, 0))
	expected := "file\tN90\n"

	err := a.writeHeader(output)
	if err != nil {
		t.Error(err)
	}

	if output.String() != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
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
	output := bytes.NewBuffer(make([]byte, 0))
	expected := "file\tN10\n"

	err := a.writeHeader(output)
	if err != nil {
		t.Error(err)
	}

	if output.String() != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
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
	output := bytes.NewBuffer(make([]byte, 0))
	expected := "file\tNG50\n"

	err := a.writeHeader(output)
	if err != nil {
		t.Error(err)
	}

	if output.String() != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
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
	output := bytes.NewBuffer(make([]byte, 0))
	expected := "file\tNG90\n"

	err := a.writeHeader(output)
	if err != nil {
		t.Error(err)
	}

	if output.String() != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
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
	output := bytes.NewBuffer(make([]byte, 0))
	expected := "file\tN50\tL50\tNG50\n"

	err := a.writeHeader(output)
	if err != nil {
		t.Error(err)
	}

	if output.String() != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
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
	expected := `stdin	8
`
	output := bytes.NewBuffer(make([]byte, 0))

	err := assemblyRecords("stdin", reader, a, output)
	if err != nil {
		t.Error(err)
	}

	if output.String() != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
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
	expected := `stdin	4
` // len -f fastaFile = 54, so n90 = 48.6 bases, 10 + 9 + 8 + 7 + 6 + 5 + 4 = 49, so n90 = 4
	output := bytes.NewBuffer(make([]byte, 0))

	err := assemblyRecords("stdin", reader, a, output)
	if err != nil {
		t.Error(err)
	}

	if output.String() != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
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
	expected := `stdin	8
`
	output := bytes.NewBuffer(make([]byte, 0))

	err := assemblyRecords("stdin", reader, a, output)
	if err != nil {
		t.Error(err)
	}

	if output.String() != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
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
	expected := `stdin	6
` // len -f fastaFile = 54, so n50 = 40 bases, 10 + 9 + 8 + 7 + 6 = 40, so n50 = 6
	output := bytes.NewBuffer(make([]byte, 0))

	err := assemblyRecords("stdin", reader, a, output)
	if err != nil {
		t.Error(err)
	}

	if output.String() != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
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
	expected := `stdin	8	3	8
` // len -f fastaFile = 54, so n50 = 40 bases, 10 + 9 + 8 + 7 + 6 = 40, so n50 = 6
	output := bytes.NewBuffer(make([]byte, 0))

	err := assemblyRecords("stdin", reader, a, output)
	if err != nil {
		t.Error(err)
	}

	if output.String() != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
	}
}
