package main

import (
	"bytes"
	"testing"
)

func TestNamesWriteHeader1(t *testing.T) {
	n := names{
		writeDescriptions: false,
	}
	output := bytes.NewBuffer(make([]byte, 0))
	expected := "file\tid\n"

	err := n.writeHeader(output)
	if err != nil {
		t.Error(err)
	}

	if output.String() != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output.String())
	}
}

func TestNamesWriteHeader2(t *testing.T) {
	n := names{
		writeDescriptions: true,
	}
	output := bytes.NewBuffer(make([]byte, 0))
	expected := "file\tdescription\n"

	err := n.writeHeader(output)
	if err != nil {
		t.Error(err)
	}

	if output.String() != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output.String())
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
	reader := NewReader(bytes.NewReader(fastaFile))
	expected := `file.fasta	Seq1
file.fasta	Seq2
`
	output := bytes.NewBuffer(make([]byte, 0))

	err := namesRecords("/path/to/my/file.fasta", reader, n, output)
	if err != nil {
		t.Error(err)
	}

	if output.String() != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output.String())
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
	reader := NewReader(bytes.NewReader(fastaFile))
	expected := `file.fasta	Seq1 Homo_sapiens X
file.fasta	Seq2 Danio_rerio Y
`
	output := bytes.NewBuffer(make([]byte, 0))

	err := namesRecords("/path/to/my/file.fasta", reader, n, output)
	if err != nil {
		t.Error(err)
	}

	if output.String() != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output.String())
	}
}
