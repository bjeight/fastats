package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/bjeight/fastats/fasta"
)

func TestLengthWriteHeader1(t *testing.T) {
	l := length{
		perFile:           false,
		writeDescriptions: false,
		writeFileNames:    false,
		lenFormat:         "b",
	}
	out := bytes.NewBuffer(make([]byte, 0))
	desiredResult := "record\tlength\n"

	err := l.writeHeader(out)
	if err != nil {
		t.Error(err)
	}

	if out.String() != desiredResult {
		fmt.Print(out.String())
		t.Fail()
	}
}

func TestLengthWriteHeader2(t *testing.T) {
	l := length{
		perFile:           false,
		writeDescriptions: false,
		writeFileNames:    false,
		lenFormat:         "kb",
	}
	out := bytes.NewBuffer(make([]byte, 0))
	desiredResult := "record\tlength_kb\n"

	err := l.writeHeader(out)
	if err != nil {
		t.Error(err)
	}

	if out.String() != desiredResult {
		fmt.Print(out.String())
		t.Fail()
	}
}

func TestLengthWriteHeader3(t *testing.T) {
	l := length{
		perFile:           false,
		writeDescriptions: false,
		writeFileNames:    false,
		lenFormat:         "mb",
	}
	out := bytes.NewBuffer(make([]byte, 0))
	desiredResult := "record\tlength_mb\n"

	err := l.writeHeader(out)
	if err != nil {
		t.Error(err)
	}

	if out.String() != desiredResult {
		fmt.Print(out.String())
		t.Fail()
	}
}

func TestLengthWriteHeader4(t *testing.T) {
	l := length{
		perFile:           false,
		writeDescriptions: false,
		writeFileNames:    false,
		lenFormat:         "gb",
	}
	out := bytes.NewBuffer(make([]byte, 0))
	desiredResult := "record\tlength_gb\n"

	err := l.writeHeader(out)
	if err != nil {
		t.Error(err)
	}

	if out.String() != desiredResult {
		fmt.Print(out.String())
		t.Fail()
	}
}

func TestLengthWriteHeader5(t *testing.T) {
	l := length{
		perFile:           true,
		writeDescriptions: false,
		writeFileNames:    false,
		lenFormat:         "",
	}
	out := bytes.NewBuffer(make([]byte, 0))
	desiredResult := "file\tlength\n"

	err := l.writeHeader(out)
	if err != nil {
		t.Error(err)
	}

	if out.String() != desiredResult {
		fmt.Print(out.String())
		t.Fail()
	}
}

func TestLengthWriteHeader6(t *testing.T) {
	l := length{
		perFile:           false,
		writeDescriptions: false,
		writeFileNames:    true,
		lenFormat:         "",
	}
	out := bytes.NewBuffer(make([]byte, 0))
	desiredResult := "file\trecord\tlength\n"

	err := l.writeHeader(out)
	if err != nil {
		t.Error(err)
	}

	if out.String() != desiredResult {
		fmt.Print(out.String())
		t.Fail()
	}
}

func TestLengthRecords1(t *testing.T) {
	l := length{
		perFile:           false,
		writeDescriptions: false,
		writeFileNames:    false,
		lenFormat:         "b",
	}
	fastaFile := []byte(`>Seq1
ATGATG
>Seq2
ATTAT-
`)
	reader := fasta.NewReader(bytes.NewReader(fastaFile))
	desiredResult := `Seq1	6
Seq2	6
`

	out, err := lengthRecords("stdin", reader, l)
	if err != nil {
		t.Error(err)
	}

	if out != desiredResult {
		fmt.Print(out)
		t.Fail()
	}
}

func TestLengthRecords2(t *testing.T) {
	l := length{
		perFile:           true,
		writeDescriptions: false,
		writeFileNames:    false,
		lenFormat:         "b",
	}
	fastaFile := []byte(`>Seq1
ATGATG
>Seq2
ATTAT-
`)
	reader := fasta.NewReader(bytes.NewReader(fastaFile))
	desiredResult := `stdin	12
`

	out, err := lengthRecords("stdin", reader, l)
	if err != nil {
		t.Error(err)
	}

	if out != desiredResult {
		fmt.Print(out)
		t.Fail()
	}
}

func TestLengthRecords3(t *testing.T) {
	l := length{
		perFile:           false,
		writeDescriptions: true,
		writeFileNames:    false,
		lenFormat:         "b",
	}
	fastaFile := []byte(`>Seq1 Homo_sapiens
ATGATG
>Seq2 Danio_rerio
ATTAT-
`)
	reader := fasta.NewReader(bytes.NewReader(fastaFile))
	desiredResult := `Seq1 Homo_sapiens	6
Seq2 Danio_rerio	6
`

	out, err := lengthRecords("/path/to/test.fa", reader, l)
	if err != nil {
		t.Error(err)
	}

	if out != desiredResult {
		fmt.Print(out)
		t.Fail()
	}
}

func TestLengthRecords4(t *testing.T) {
	l := length{
		perFile:           false,
		writeDescriptions: true,
		writeFileNames:    true,
		lenFormat:         "b",
	}
	fastaFile := []byte(`>Seq1 Homo_sapiens
ATGATG
>Seq2 Danio_rerio
ATTAT-
`)
	reader := fasta.NewReader(bytes.NewReader(fastaFile))
	desiredResult := `test.fa	Seq1 Homo_sapiens	6
test.fa	Seq2 Danio_rerio	6
`

	out, err := lengthRecords("/path/to/test.fa", reader, l)
	if err != nil {
		t.Error(err)
	}

	if out != desiredResult {
		fmt.Print(out)
		t.Fail()
	}
}
