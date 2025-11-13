package main

import (
	"bytes"
	"testing"
)

func TestNumWriteHeader(t *testing.T) {
	n := num{}
	output := bytes.NewBuffer(make([]byte, 0))
	expected := "file\tn_records\n"

	err := n.writeHeader(output)
	if err != nil {
		t.Error(err)
	}

	if output.String() != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output.String())
	}
}

func TestNumRecords(t *testing.T) {
	n := num{}
	fastaFile := []byte(`>Seq1
ATGATG
>Seq2
ATTAT-
`)
	reader := NewReader(bytes.NewReader(fastaFile))
	expected := `stdin	2
`
	output := bytes.NewBuffer(make([]byte, 0))

	err := numRecords("stdin", reader, n, output)
	if err != nil {
		t.Error(err)
	}

	if output.String() != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output.String())
	}
}
