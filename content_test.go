package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/bjeight/fastats/fasta"
)

func TestContentWriteHeader1(t *testing.T) {
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

func TestContentWriteHeader2(t *testing.T) {
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

func TestContentWriteHeader3(t *testing.T) {
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

func TestContentWriteHeader4(t *testing.T) {
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

func TestContentWriteHeader5(t *testing.T) {
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

func TestContentWriteHeader6(t *testing.T) {
	c := content{
		perFile:           false,
		writeDescriptions: false,
		writeFileNames:    false,
		patterns: []pattern{
			{
				stat:    "prop",
				bases:   "ATat",
				inverse: true,
			},
			{
				stat:    "count",
				bases:   "ATat",
				inverse: true,
			},
		}}
	out := bytes.NewBuffer(make([]byte, 0))
	desiredResult := "record\t!ATat_prop\t!ATat_count\n"

	err := c.writeHeader(out)
	if err != nil {
		t.Error(err)
	}

	if out.String() != desiredResult {
		fmt.Print(out.String())
		t.Fail()
	}
}

func TestContentWriteHeader7(t *testing.T) {
	c := content{
		perFile:           false,
		writeDescriptions: false,
		writeFileNames:    false,
		patterns: []pattern{
			{
				stat:         "prop",
				bases:        "ATat",
				headerPrefix: "ATcontent",
			},
			{
				stat:  "count",
				bases: "ATat",
			},
		}}

	expected := "record\tATcontent_prop\tATat_count\n"
	output := bytes.NewBuffer(make([]byte, 0))
	err := c.writeHeader(output)
	if err != nil {
		t.Error(err)
	}

	if output.String() != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
	}
}

func TestContentRecords1(t *testing.T) {
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

	expected := `Seq1	0.666667
Seq2	0.833333
`
	output := bytes.NewBuffer(make([]byte, 0))

	err := contentRecords("stdin", reader, c, output)
	if err != nil {
		t.Error(err)
	}

	if output.String() != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
	}
}

func TestContentRecords2(t *testing.T) {
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
	expected := `Seq1	0.333333
Seq2	0.000000
`
	output := bytes.NewBuffer(make([]byte, 0))

	err := contentRecords("stdin", reader, c, output)
	if err != nil {
		t.Error(err)
	}

	if output.String() != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
	}
}

func TestContentRecords3(t *testing.T) {
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
	expected := `stdin	0.750000
`
	output := bytes.NewBuffer(make([]byte, 0))

	err := contentRecords("stdin", reader, c, output)
	if err != nil {
		t.Error(err)
	}

	if output.String() != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
	}
}

func TestContentRecords4(t *testing.T) {
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
	expected := `Seq1	4
Seq2	5
`
	output := bytes.NewBuffer(make([]byte, 0))

	err := contentRecords("stdin", reader, c, output)
	if err != nil {
		t.Error(err)
	}

	if output.String() != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
	}
}

func TestContentRecords5(t *testing.T) {
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
	expected := `stdin	9
`
	output := bytes.NewBuffer(make([]byte, 0))

	err := contentRecords("stdin", reader, c, output)
	if err != nil {
		t.Error(err)
	}

	if output.String() != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
	}
}

func TestContentRecords6(t *testing.T) {
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
	expected := `my.fasta	Seq1	4
my.fasta	Seq2	5
`
	output := bytes.NewBuffer(make([]byte, 0))

	err := contentRecords("/path/to/my.fasta", reader, c, output)
	if err != nil {
		t.Error(err)
	}

	if output.String() != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
	}
}

func TestContentRecords7(t *testing.T) {
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
	expected := `Seq1 Homo_sapiens	4
Seq2 Danio_rerio	5
`
	output := bytes.NewBuffer(make([]byte, 0))

	err := contentRecords("/path/to/my.fasta", reader, c, output)
	if err != nil {
		t.Error(err)
	}

	if output.String() != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
	}
}

func TestContentRecords8(t *testing.T) {
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
	expected := `my.fasta	Seq1 Homo_sapiens	6
my.fasta	Seq2 Danio_rerio	5
`
	output := bytes.NewBuffer(make([]byte, 0))

	err := contentRecords("/path/to/my.fasta", reader, c, output)
	if err != nil {
		t.Error(err)
	}

	if output.String() != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
	}
}

func TestContentRecords9(t *testing.T) {
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
	expected := `Seq1	1.000000	0.333333	2
Seq2	0.833333	0.000000	0
`
	output := bytes.NewBuffer(make([]byte, 0))

	err := contentRecords("/path/to/my.fasta", reader, c, output)
	if err != nil {
		t.Error(err)
	}

	if output.String() != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
	}
}

func TestContentRecords10(t *testing.T) {
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
	expected := `my.fasta	11
`
	output := bytes.NewBuffer(make([]byte, 0))

	err := contentRecords("/path/to/my.fasta", reader, c, output)
	if err != nil {
		t.Error(err)
	}

	if output.String() != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
	}
}
