package main

import (
	"testing"

	"github.com/bjeight/fastats/fasta"
)

func TestReturnFileName(t *testing.T) {
	path := "/path/to/file.fasta"
	filename := returnFileName(path)
	if filename != "file.fasta" {
		t.Fail()
	}
}

func TestReturnRecordName(t *testing.T) {
	record := fasta.Record{
		ID:          "seq1",
		Description: "seq1 Homo_sapiens",
	}
	id := returnRecordName(record, false)
	if id != "seq1" {
		t.Fail()
	}
	desc := returnRecordName(record, true)
	if desc != "seq1 Homo_sapiens" {
		t.Fail()
	}
}
