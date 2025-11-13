package main

import (
	"bytes"
	"slices"
	"testing"
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

func TestAssemblyNStat1(t *testing.T) {
	contigLengths := []int64{6, 7, 10, 3, 4, 5, 8, 2, 9}
	slices.Sort(contigLengths)
	NX := nStat(contigLengths, 54, 50)
	if NX != 8 {
		t.Errorf("expected:\n 8\n, got:\n %d", NX)
	}
}

func TestAssemblyNStat2(t *testing.T) {
	contigLengths := []int64{6, 7, 10, 3, 4, 5, 8, 2, 9}
	slices.Sort(contigLengths)
	NX := nStat(contigLengths, 54, 90)
	if NX != 4 {
		t.Errorf("expected:\n 4,\n got:\n %d", NX)
	}
}

func TestAssemblyNStat3(t *testing.T) {
	contigLengths := []int64{6, 7, 10, 3, 4, 5, 8, 2, 9}
	slices.Sort(contigLengths)
	NGX := nStat(contigLengths, 80, 50)
	if NGX != 6 {
		t.Errorf("expected:\n 6,\n got:\n %d", NGX)
	}
}

func TestAssemblyLStat1(t *testing.T) {
	contigLengths := []int64{6, 7, 10, 3, 4, 5, 8, 2, 9}
	slices.Sort(contigLengths)
	LX := lStat(contigLengths, 54, 50)
	if LX != 3 {
		t.Errorf("expected:\n 3,\n got:\n %d", LX)
	}
}

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
	expected := "file\tn_records\tlength\tN50\n"

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
	expected := "file\tn_records\tlength\tN90\n"

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
	expected := "file\tn_records\tlength\tN10\n"

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
	expected := "file\tn_records\tlength\tNG50\n"

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
	expected := "file\tn_records\tlength\tNG90\n"

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
	expected := "file\tn_records\tlength\tN50\tL50\tNG50\n"

	err := a.writeHeader(output)
	if err != nil {
		t.Error(err)
	}

	if output.String() != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
	}
}

func TestAssemblyWriteHeader7(t *testing.T) {
	a := assembly{
		stats: []assemblyStatistic{
			{
				sType:  "N",
				sValue: 50,
			},
		},
		lenFormat: "kb",
	}
	output := bytes.NewBuffer(make([]byte, 0))
	expected := "file\tn_records\tlength_kb\tN50_kb\n"

	err := a.writeHeader(output)
	if err != nil {
		t.Error(err)
	}

	if output.String() != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
	}
}

func TestAssemblyWriteHeader8(t *testing.T) {
	a := assembly{
		stats: []assemblyStatistic{
			{
				sType:  "N",
				sValue: 50,
			},
		},
		lenFormat: "mb",
	}
	output := bytes.NewBuffer(make([]byte, 0))
	expected := "file\tn_records\tlength_mb\tN50_mb\n"

	err := a.writeHeader(output)
	if err != nil {
		t.Error(err)
	}

	if output.String() != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
	}
}

func TestAssemblyWriteHeader9(t *testing.T) {
	a := assembly{
		stats: []assemblyStatistic{
			{
				sType:  "N",
				sValue: 50,
			},
		},
		lenFormat: "gb",
	}
	output := bytes.NewBuffer(make([]byte, 0))
	expected := "file\tn_records\tlength_gb\tN50_gb\n"

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
	reader := NewReader(bytes.NewReader(fastaFile))
	expected := `stdin	9	54	8
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
	reader := NewReader(bytes.NewReader(fastaFile))
	expected := `stdin	9	54	4
` // len -f fastaFile = 54, so n90 = 48.6 bases, 10 + 9 + 8 + 7 + 6 + 5 + 4 = 49, so N90 = 4
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
	reader := NewReader(bytes.NewReader(fastaFile))
	expected := `stdin	9	54	8
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
	reader := NewReader(bytes.NewReader(fastaFile))
	expected := `stdin	9	54	6
`
	// genome size = 80 bases, so N50 = 40 bases, 10 + 9 + 8 + 7 + 6 = 40, so N50 = 6
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
	reader := NewReader(bytes.NewReader(fastaFile))
	expected := `stdin	9	54	8	3	8
`
	// len -f fastaFile = 54, so n50 = 27 bases, 10 + 9 + 8 = 27, so N50 = 8, and L50 = 3
	// genome size = 54 bases, so n50 = 27 bases, 10 + 9 + 8 = 27, so NG50 = 8
	output := bytes.NewBuffer(make([]byte, 0))

	err := assemblyRecords("stdin", reader, a, output)
	if err != nil {
		t.Error(err)
	}

	if output.String() != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
	}
}
