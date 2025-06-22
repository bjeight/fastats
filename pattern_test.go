package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/bjeight/fastats/fasta"
)

func TestPatternWriteHeader1(t *testing.T) {
	p := pattern{
		perFile:           false,
		writeCounts:       false,
		writeDescriptions: false,
		writeFileNames:    false,
		bases:             "ATat",
	}
	out := bytes.NewBuffer(make([]byte, 0))
	desiredResult := "record\tATat_prop\n"

	err := p.writeHeader(out)
	if err != nil {
		t.Error(err)
	}

	if out.String() != desiredResult {
		fmt.Print(out.String())
		t.Fail()
	}
}

func TestPatternWriteHeader2(t *testing.T) {
	p := pattern{
		perFile:           true,
		writeCounts:       false,
		writeDescriptions: false,
		writeFileNames:    false,
		bases:             "ATat",
	}
	out := bytes.NewBuffer(make([]byte, 0))
	desiredResult := "file\tATat_prop\n"

	err := p.writeHeader(out)
	if err != nil {
		t.Error(err)
	}

	if out.String() != desiredResult {
		fmt.Print(out.String())
		t.Fail()
	}
}

func TestPatternWriteHeader3(t *testing.T) {
	p := pattern{
		perFile:           false,
		writeCounts:       true,
		writeDescriptions: false,
		writeFileNames:    false,
		bases:             "ATat",
	}
	out := bytes.NewBuffer(make([]byte, 0))
	desiredResult := "record\tATat_count\n"

	err := p.writeHeader(out)
	if err != nil {
		t.Error(err)
	}

	if out.String() != desiredResult {
		fmt.Print(out.String())
		t.Fail()
	}
}

func TestPatternWriteHeader4(t *testing.T) {
	p := pattern{
		perFile:           false,
		writeCounts:       true,
		writeDescriptions: false,
		writeFileNames:    true,
		bases:             "ATat",
	}
	out := bytes.NewBuffer(make([]byte, 0))
	desiredResult := "file\trecord\tATat_count\n"

	err := p.writeHeader(out)
	if err != nil {
		t.Error(err)
	}

	if out.String() != desiredResult {
		fmt.Print(out.String())
		t.Fail()
	}
}

func TestPatternWriteHeader5(t *testing.T) {
	p := pattern{
		perFile:           true,
		writeCounts:       true,
		writeDescriptions: false,
		writeFileNames:    true,
		bases:             "ATat",
	}
	out := bytes.NewBuffer(make([]byte, 0))
	desiredResult := "file\tATat_count\n"

	err := p.writeHeader(out)
	if err != nil {
		t.Error(err)
	}

	if out.String() != desiredResult {
		fmt.Print(out.String())
		t.Fail()
	}
}

func TestPatternRecords1(t *testing.T) {
	p := pattern{
		perFile:           false,
		writeCounts:       false,
		writeDescriptions: false,
		writeFileNames:    false,
		bases:             "ATat",
	}
	fastaFile := []byte(`>Seq1
ATGATG
>Seq2
ATTAT-
`)
	reader := fasta.NewReader(bytes.NewReader(fastaFile))
	desiredResult := `Seq1	0.666667
Seq2	0.833333
`
	out := bytes.NewBuffer(make([]byte, 0))

	err := patternRecords("stdin", reader, p, out)
	if err != nil {
		t.Error(err)
	}

	if out.String() != desiredResult {
		fmt.Print(out.String())
		t.Fail()
	}
}

func TestPatternRecords2(t *testing.T) {
	p := pattern{
		perFile:           false,
		writeCounts:       false,
		writeDescriptions: false,
		writeFileNames:    false,
		bases:             "GCgc",
	}
	fastaFile := []byte(`>Seq1
ATGATG
>Seq2
ATTAT-
`)
	reader := fasta.NewReader(bytes.NewReader(fastaFile))
	desiredResult := `Seq1	0.333333
Seq2	0.000000
`
	out := bytes.NewBuffer(make([]byte, 0))

	err := patternRecords("stdin", reader, p, out)
	if err != nil {
		t.Error(err)
	}

	if out.String() != desiredResult {
		fmt.Print(out.String())
		t.Fail()
	}
}

func TestPatternRecords3(t *testing.T) {
	p := pattern{
		perFile:           true,
		writeCounts:       false,
		writeDescriptions: false,
		writeFileNames:    false,
		bases:             "ATat",
	}
	fastaFile := []byte(`>Seq1
ATGATG
>Seq2
ATTAT-
`)
	reader := fasta.NewReader(bytes.NewReader(fastaFile))
	desiredResult := `stdin	0.750000
`
	out := bytes.NewBuffer(make([]byte, 0))

	err := patternRecords("stdin", reader, p, out)
	if err != nil {
		t.Error(err)
	}

	if out.String() != desiredResult {
		fmt.Print(out.String())
		t.Fail()
	}
}

func TestPatternRecords4(t *testing.T) {
	p := pattern{
		perFile:           false,
		writeCounts:       true,
		writeDescriptions: false,
		writeFileNames:    false,
		bases:             "ATat",
	}
	fastaFile := []byte(`>Seq1
ATGATG
>Seq2
ATTAT-
`)
	reader := fasta.NewReader(bytes.NewReader(fastaFile))
	desiredResult := `Seq1	4
Seq2	5
`
	out := bytes.NewBuffer(make([]byte, 0))

	err := patternRecords("stdin", reader, p, out)
	if err != nil {
		t.Error(err)
	}

	if out.String() != desiredResult {
		fmt.Print(out.String())
		t.Fail()
	}
}

func TestPatternRecords5(t *testing.T) {
	p := pattern{
		perFile:           true,
		writeCounts:       true,
		writeDescriptions: false,
		writeFileNames:    false,
		bases:             "ATat",
	}
	fastaFile := []byte(`>Seq1
ATGATG
>Seq2
ATTAT-
`)
	reader := fasta.NewReader(bytes.NewReader(fastaFile))
	desiredResult := `stdin	9
`
	out := bytes.NewBuffer(make([]byte, 0))

	err := patternRecords("stdin", reader, p, out)
	if err != nil {
		t.Error(err)
	}

	if out.String() != desiredResult {
		fmt.Print(out.String())
		t.Fail()
	}
}

func TestPatternRecords6(t *testing.T) {
	p := pattern{
		perFile:           false,
		writeCounts:       true,
		writeDescriptions: false,
		writeFileNames:    true,
		bases:             "ATat",
	}
	fastaFile := []byte(`>Seq1
ATGATG
>Seq2
ATTAT-
`)
	reader := fasta.NewReader(bytes.NewReader(fastaFile))
	desiredResult := `my.fasta	Seq1	4
my.fasta	Seq2	5
`
	out := bytes.NewBuffer(make([]byte, 0))

	err := patternRecords("/path/to/my.fasta", reader, p, out)
	if err != nil {
		t.Error(err)
	}

	if out.String() != desiredResult {
		fmt.Print(out.String())
		t.Fail()
	}
}

func TestPatternRecords7(t *testing.T) {
	p := pattern{
		perFile:           false,
		writeCounts:       true,
		writeDescriptions: true,
		writeFileNames:    false,
		bases:             "ATat",
	}
	fastaFile := []byte(`>Seq1 Homo_sapiens
ATGATG
>Seq2 Danio_rerio
ATTAT-
`)
	reader := fasta.NewReader(bytes.NewReader(fastaFile))
	desiredResult := `Seq1 Homo_sapiens	4
Seq2 Danio_rerio	5
`
	out := bytes.NewBuffer(make([]byte, 0))

	err := patternRecords("/path/to/my.fasta", reader, p, out)
	if err != nil {
		t.Error(err)
	}

	if out.String() != desiredResult {
		fmt.Print(out.String())
		t.Fail()
	}
}

func TestPatternRecords8(t *testing.T) {
	p := pattern{
		perFile:           false,
		writeCounts:       true,
		writeDescriptions: true,
		writeFileNames:    true,
		bases:             "ATGCatgc",
	}
	fastaFile := []byte(`>Seq1 Homo_sapiens
ATGATG
>Seq2 Danio_rerio
ATTAT-
`)
	reader := fasta.NewReader(bytes.NewReader(fastaFile))
	desiredResult := `my.fasta	Seq1 Homo_sapiens	6
my.fasta	Seq2 Danio_rerio	5
`
	out := bytes.NewBuffer(make([]byte, 0))

	err := patternRecords("/path/to/my.fasta", reader, p, out)
	if err != nil {
		t.Error(err)
	}

	if out.String() != desiredResult {
		fmt.Print(out.String())
		t.Fail()
	}
}
