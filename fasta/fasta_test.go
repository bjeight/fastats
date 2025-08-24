package fasta

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"reflect"
	"testing"
)

func TestRead(t *testing.T) {
	fastaData := []byte(
		`>seq1
ATGC
>seq2 something
ATG-ATG-
ATGCATGC
ATGC
>seq3
AT
`)
	fastaR := bytes.NewReader(fastaData)
	r := NewReader(fastaR)

	expected := Record{
		ID:          "seq1",
		Description: "seq1",
		Seq:         []byte("ATGC"),
	}
	record, err := r.Read()
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(record, expected) {
		fmt.Println(record)
		t.Errorf("problem in TestRead(seq1)")
	}

	expected = Record{
		ID:          "seq2",
		Description: "seq2 something",
		Seq:         []byte("ATG-ATG-ATGCATGCATGC"),
	}
	record, err = r.Read()
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(record, expected) {
		fmt.Println(record)
		t.Errorf("problem in TestRead(seq2)")
	}

	expected = Record{
		ID:          "seq3",
		Description: "seq3",
		Seq:         []byte("AT"),
	}
	record, err = r.Read()
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(record, expected) {
		fmt.Println(record)
		t.Errorf("problem in TestRead(seq3)")
	}

	expected = Record{}
	record, err = r.Read()
	if !errors.Is(err, io.EOF) {
		t.Error(err)
	}
	if !reflect.DeepEqual(record, expected) {
		fmt.Println(record)
		t.Errorf("problem in TestRead(last read)")
	}
}
