package main

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

	desiredResult := FastaRecord{
		ID:          "seq1",
		Description: "seq1",
		Seq:         []byte("ATGC"),
	}
	record, err := r.Read()
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(record, desiredResult) {
		fmt.Println(record)
		t.Errorf("problem in TestRead(seq1)")
	}

	desiredResult = FastaRecord{
		ID:          "seq2",
		Description: "seq2 something",
		Seq:         []byte("ATG-ATG-ATGCATGCATGC"),
	}
	record, err = r.Read()
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(record, desiredResult) {
		fmt.Println(record)
		t.Errorf("problem in TestRead(seq2)")
	}

	desiredResult = FastaRecord{
		ID:          "seq3",
		Description: "seq3",
		Seq:         []byte("AT"),
	}
	record, err = r.Read()
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(record, desiredResult) {
		fmt.Println(record)
		t.Errorf("problem in TestRead(seq3)")
	}

	desiredResult = FastaRecord{}
	record, err = r.Read()
	if !errors.Is(err, io.EOF) {
		t.Error(err)
	}
	if !reflect.DeepEqual(record, desiredResult) {
		fmt.Println(record)
		t.Errorf("problem in TestRead(last read)")
	}
}
