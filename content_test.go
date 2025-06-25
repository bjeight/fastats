package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/bjeight/fastats/fasta"
)

func TestPatternWriteHeader1(t *testing.T) {
	c := content{
		perFile:           false,
		writeDescriptions: false,
		writeFileNames:    false,
		patterns: []pattern{
			{
				stat:  "prop",
				bases: "ATat",
			},
		},
	}
	out := bytes.NewBuffer(make([]byte, 0))
	desiredResult := "record\tATat_prop\n"

	err := c.writeHeader(out)
	if err != nil {
		t.Error(err)
	}

	if out.String() != desiredResult {
		fmt.Print(out.String())
		t.Fail()
	}
}

func TestPatternWriteHeader2(t *testing.T) {
	c := content{
		perFile:           true,
		writeDescriptions: false,
		writeFileNames:    false,
		patterns: []pattern{
			{
				stat:  "prop",
				bases: "ATat",
			},
		},
	}
	out := bytes.NewBuffer(make([]byte, 0))
	desiredResult := "file\tATat_prop\n"

	err := c.writeHeader(out)
	if err != nil {
		t.Error(err)
	}

	if out.String() != desiredResult {
		fmt.Print(out.String())
		t.Fail()
	}
}

func TestPatternWriteHeader3(t *testing.T) {
	c := content{
		perFile:           false,
		writeDescriptions: false,
		writeFileNames:    false,
		patterns: []pattern{
			{
				stat:  "count",
				bases: "ATat",
			},
		},
	}
	out := bytes.NewBuffer(make([]byte, 0))
	desiredResult := "record\tATat_count\n"

	err := c.writeHeader(out)
	if err != nil {
		t.Error(err)
	}

	if out.String() != desiredResult {
		fmt.Print(out.String())
		t.Fail()
	}
}

func TestPatternWriteHeader4(t *testing.T) {
	c := content{
		perFile:           false,
		writeDescriptions: false,
		writeFileNames:    true,
		patterns: []pattern{
			{
				stat:  "count",
				bases: "ATat",
			},
		}}
	out := bytes.NewBuffer(make([]byte, 0))
	desiredResult := "file\trecord\tATat_count\n"

	err := c.writeHeader(out)
	if err != nil {
		t.Error(err)
	}

	if out.String() != desiredResult {
		fmt.Print(out.String())
		t.Fail()
	}
}

func TestPatternWriteHeader6(t *testing.T) {
	c := content{
		perFile:           true,
		writeDescriptions: false,
		writeFileNames:    true,
		patterns: []pattern{
			{
				stat:  "prop",
				bases: "ATat",
			},
			{
				stat:  "count",
				bases: "ATat",
			},
		}}
	out := bytes.NewBuffer(make([]byte, 0))
	desiredResult := "file\tATat_prop\tATat_count\n"

	err := c.writeHeader(out)
	if err != nil {
		t.Error(err)
	}

	if out.String() != desiredResult {
		fmt.Print(out.String())
		t.Fail()
	}
}

func TestPatternRecords1(t *testing.T) {
	c := content{
		perFile:           false,
		writeDescriptions: false,
		writeFileNames:    false,
		patterns: []pattern{
			{
				stat:  "prop",
				bases: "ATat",
			},
		},
	}
	fastaFile := []byte(`>Seq1
ATGATG
>Seq2
ATTAT-
`)
	reader := fasta.NewReader(bytes.NewReader(fastaFile))
	desiredResult := `Seq1	0.666667
Seq2	0.833333
`
	out := bytes.NewBuffer(make([]byte, 0))

	err := contentRecords("stdin", reader, c, out)
	if err != nil {
		t.Error(err)
	}

	if out.String() != desiredResult {
		fmt.Print(out.String())
		t.Fail()
	}
}

func TestPatternRecords2(t *testing.T) {
	c := content{
		perFile:           false,
		writeDescriptions: false,
		writeFileNames:    false,
		patterns: []pattern{
			{
				stat:  "prop",
				bases: "GCgc",
			},
		},
	}
	fastaFile := []byte(`>Seq1
ATGATG
>Seq2
ATTAT-
`)
	reader := fasta.NewReader(bytes.NewReader(fastaFile))
	desiredResult := `Seq1	0.333333
Seq2	0.000000
`
	out := bytes.NewBuffer(make([]byte, 0))

	err := contentRecords("stdin", reader, c, out)
	if err != nil {
		t.Error(err)
	}

	if out.String() != desiredResult {
		fmt.Print(out.String())
		t.Fail()
	}
}

func TestPatternRecords3(t *testing.T) {
	c := content{
		perFile:           true,
		writeDescriptions: false,
		writeFileNames:    false,
		patterns: []pattern{
			{
				stat:  "prop",
				bases: "ATat",
			},
		},
	}
	fastaFile := []byte(`>Seq1
ATGATG
>Seq2
ATTAT-
`)
	reader := fasta.NewReader(bytes.NewReader(fastaFile))
	desiredResult := `stdin	0.750000
`
	out := bytes.NewBuffer(make([]byte, 0))

	err := contentRecords("stdin", reader, c, out)
	if err != nil {
		t.Error(err)
	}

	if out.String() != desiredResult {
		fmt.Print(out.String())
		t.Fail()
	}
}

func TestPatternRecords4(t *testing.T) {
	c := content{
		perFile:           false,
		writeDescriptions: false,
		writeFileNames:    false,
		patterns: []pattern{
			{
				stat:  "count",
				bases: "ATat",
			},
		},
	}
	fastaFile := []byte(`>Seq1
ATGATG
>Seq2
ATTAT-
`)
	reader := fasta.NewReader(bytes.NewReader(fastaFile))
	desiredResult := `Seq1	4
Seq2	5
`
	out := bytes.NewBuffer(make([]byte, 0))

	err := contentRecords("stdin", reader, c, out)
	if err != nil {
		t.Error(err)
	}

	if out.String() != desiredResult {
		fmt.Print(out.String())
		t.Fail()
	}
}

func TestPatternRecords5(t *testing.T) {
	c := content{
		perFile:           true,
		writeDescriptions: false,
		writeFileNames:    false,
		patterns: []pattern{
			{
				stat:  "count",
				bases: "ATat",
			},
		},
	}
	fastaFile := []byte(`>Seq1
ATGATG
>Seq2
ATTAT-
`)
	reader := fasta.NewReader(bytes.NewReader(fastaFile))
	desiredResult := `stdin	9
`
	out := bytes.NewBuffer(make([]byte, 0))

	err := contentRecords("stdin", reader, c, out)
	if err != nil {
		t.Error(err)
	}

	if out.String() != desiredResult {
		fmt.Print(out.String())
		t.Fail()
	}
}

func TestPatternRecords6(t *testing.T) {
	c := content{
		perFile:           false,
		writeDescriptions: false,
		writeFileNames:    true,
		patterns: []pattern{
			{
				stat:  "count",
				bases: "ATat",
			},
		}}
	fastaFile := []byte(`>Seq1
ATGATG
>Seq2
ATTAT-
`)
	reader := fasta.NewReader(bytes.NewReader(fastaFile))
	desiredResult := `my.fasta	Seq1	4
my.fasta	Seq2	5
`
	out := bytes.NewBuffer(make([]byte, 0))

	err := contentRecords("/path/to/my.fasta", reader, c, out)
	if err != nil {
		t.Error(err)
	}

	if out.String() != desiredResult {
		fmt.Print(out.String())
		t.Fail()
	}
}

func TestPatternRecords7(t *testing.T) {
	c := content{
		perFile:           false,
		writeDescriptions: true,
		writeFileNames:    false,
		patterns: []pattern{
			{
				stat:  "count",
				bases: "ATat",
			},
		},
	}
	fastaFile := []byte(`>Seq1 Homo_sapiens
ATGATG
>Seq2 Danio_rerio
ATTAT-
`)
	reader := fasta.NewReader(bytes.NewReader(fastaFile))
	desiredResult := `Seq1 Homo_sapiens	4
Seq2 Danio_rerio	5
`
	out := bytes.NewBuffer(make([]byte, 0))

	err := contentRecords("/path/to/my.fasta", reader, c, out)
	if err != nil {
		t.Error(err)
	}

	if out.String() != desiredResult {
		fmt.Print(out.String())
		t.Fail()
	}
}

func TestPatternRecords8(t *testing.T) {
	c := content{
		perFile:           false,
		writeDescriptions: true,
		writeFileNames:    true,
		patterns: []pattern{
			{
				stat:  "count",
				bases: "ATatGCgc",
			},
		}}
	fastaFile := []byte(`>Seq1 Homo_sapiens
ATGATG
>Seq2 Danio_rerio
ATTAT-
`)
	reader := fasta.NewReader(bytes.NewReader(fastaFile))
	desiredResult := `my.fasta	Seq1 Homo_sapiens	6
my.fasta	Seq2 Danio_rerio	5
`
	out := bytes.NewBuffer(make([]byte, 0))

	err := contentRecords("/path/to/my.fasta", reader, c, out)
	if err != nil {
		t.Error(err)
	}

	if out.String() != desiredResult {
		fmt.Print(out.String())
		t.Fail()
	}
}

func TestPatternRecords9(t *testing.T) {
	c := content{
		perFile:           false,
		writeDescriptions: false,
		writeFileNames:    false,
		patterns: []pattern{
			{
				stat:  "prop",
				bases: "ATatGCgc",
			},
			{
				stat:  "prop",
				bases: "GCgc",
			},
			{
				stat:  "count",
				bases: "GCgc",
			},
		},
	}
	fastaFile := []byte(`>Seq1 Homo_sapiens
ATGATG
>Seq2 Danio_rerio
ATTAT-
`)
	reader := fasta.NewReader(bytes.NewReader(fastaFile))
	desiredResult := `Seq1	1.000000	0.333333	2
Seq2	0.833333	0.000000	0
`
	out := bytes.NewBuffer(make([]byte, 0))

	err := contentRecords("/path/to/my.fasta", reader, c, out)
	if err != nil {
		t.Error(err)
	}

	if out.String() != desiredResult {
		fmt.Print(out.String())
		t.Fail()
	}
}

func TestPatternRecords10(t *testing.T) {
	c := content{
		perFile:           true,
		writeDescriptions: false,
		writeFileNames:    true,
		patterns: []pattern{
			{
				stat:  "count",
				bases: "ATatGCgc",
			},
		}}
	fastaFile := []byte(`>Seq1 Homo_sapiens
ATGATG
>Seq2 Danio_rerio
ATTAT-
`)
	reader := fasta.NewReader(bytes.NewReader(fastaFile))
	desiredResult := `my.fasta	11
`
	out := bytes.NewBuffer(make([]byte, 0))

	err := contentRecords("/path/to/my.fasta", reader, c, out)
	if err != nil {
		t.Error(err)
	}

	if out.String() != desiredResult {
		fmt.Print(out.String())
		t.Fail()
	}
}
