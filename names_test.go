package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/bjeight/fastats/fasta"
)

func TestNamesWriteHeader1(t *testing.T) {
	n := names{
		writeDescriptions: false,
	}
	out := bytes.NewBuffer(make([]byte, 0))
	desiredResult := "file\tid\n"

	err := n.writeHeader(out)
	if err != nil {
		t.Error(err)
	}

	if out.String() != desiredResult {
		fmt.Print(out.String())
		t.Fail()
	}
}

func TestNamesWriteHeader2(t *testing.T) {
	n := names{
		writeDescriptions: true,
	}
	out := bytes.NewBuffer(make([]byte, 0))
	desiredResult := "file\tdescription\n"

	err := n.writeHeader(out)
	if err != nil {
		t.Error(err)
	}

	if out.String() != desiredResult {
		fmt.Print(out.String())
		t.Fail()
	}
}

func TestNamesRecords1(t *testing.T) {
	n := names{
		writeDescriptions: false,
	}
	fastaFile := []byte(`>Seq1 Homo_sapiens X
ATGATG
>Seq2 Danio_rerio Y
ATTAT-
`)
	reader := fasta.NewReader(bytes.NewReader(fastaFile))
	desiredResult := `file.fasta	Seq1
file.fasta	Seq2
`
	out := bytes.NewBuffer(make([]byte, 0))

	err := namesRecords("/path/to/my/file.fasta", reader, n, out)
	if err != nil {
		t.Error(err)
	}

	if out.String() != desiredResult {
		fmt.Print(out.String())
		t.Fail()
	}
}

func TestNamesRecords2(t *testing.T) {
	n := names{
		writeDescriptions: true,
	}
	fastaFile := []byte(`>Seq1 Homo_sapiens X
ATGATG
>Seq2 Danio_rerio Y
ATTAT-
`)
	reader := fasta.NewReader(bytes.NewReader(fastaFile))
	desiredResult := `file.fasta	Seq1 Homo_sapiens X
file.fasta	Seq2 Danio_rerio Y
`
	out := bytes.NewBuffer(make([]byte, 0))

	err := namesRecords("/path/to/my/file.fasta", reader, n, out)
	if err != nil {
		t.Error(err)
	}

	if out.String() != desiredResult {
		fmt.Print(out.String())
		t.Fail()
	}
}
