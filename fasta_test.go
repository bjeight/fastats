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

	bc := [256]int64{}
	for _, b := range "ATGC" {
		bc[b]++
	}
	expected := Record{
		ID:          "seq1",
		Description: "seq1",
		Len:         4,
		BaseCounts:  bc,
	}

	record, err := r.ReadCalcBaseCounts()
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(record, expected) {
		fmt.Println(record)
		t.Errorf("problem in TestRead(seq1)")
	}

	bc = [256]int64{}
	for _, b := range "ATG-ATG-ATGCATGCATGC" {
		bc[b]++
	}
	expected = Record{
		ID:          "seq2",
		Description: "seq2 something",
		Len:         20,
		BaseCounts:  bc,
	}
	record, err = r.ReadCalcBaseCounts()
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(record, expected) {
		fmt.Println(record)
		t.Errorf("problem in TestRead(seq2)")
	}

	bc = [256]int64{}
	for _, b := range "AT" {
		bc[b]++
	}
	expected = Record{
		ID:          "seq3",
		Description: "seq3",
		Len:         2,
		BaseCounts:  bc,
	}
	record, err = r.ReadCalcBaseCounts()
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(record, expected) {
		fmt.Println(record)
		t.Errorf("problem in TestRead(seq3)")
	}

	expected = Record{}
	record, err = r.ReadCalcBaseCounts()
	if !errors.Is(err, io.EOF) {
		t.Error(err)
	}
	if !reflect.DeepEqual(record, expected) {
		fmt.Println(record)
		t.Errorf("problem in TestRead(last read)")
	}
}
