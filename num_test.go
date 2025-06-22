package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/bjeight/fastats/fasta"
)

func TestNumWriteHeader(t *testing.T) {
	n := num{}
	out := bytes.NewBuffer(make([]byte, 0))
	desiredResult := "file\tn_records\n"

	err := n.writeHeader(out)
	if err != nil {
		t.Error(err)
	}

	if out.String() != desiredResult {
		fmt.Print(out.String())
		t.Fail()
	}
}

func TestNumRecords(t *testing.T) {
	n := num{}
	fastaFile := []byte(`>Seq1
ATGATG
>Seq2
ATTAT-
`)
	reader := fasta.NewReader(bytes.NewReader(fastaFile))
	desiredResult := `stdin	2
`

	out, err := numRecords("stdin", reader, n)
	if err != nil {
		t.Error(err)
	}

	if out != desiredResult {
		fmt.Print(out)
		t.Fail()
	}
}
