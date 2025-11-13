package main

import (
	"bytes"
	"testing"
)

func TestLengthWriteHeader1(t *testing.T) {
	l := length{
		perFile:           false,
		writeDescriptions: false,
		writeFileNames:    false,
		lenFormat:         "b",
	}
	output := bytes.NewBuffer(make([]byte, 0))
	expected := "record\tlength\n"

	err := l.writeHeader(output)
	if err != nil {
		t.Error(err)
	}

	if output.String() != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
	}
}

func TestLengthWriteHeader2(t *testing.T) {
	l := length{
		perFile:           false,
		writeDescriptions: false,
		writeFileNames:    false,
		lenFormat:         "kb",
	}
	output := bytes.NewBuffer(make([]byte, 0))
	expected := "record\tlength_kb\n"

	err := l.writeHeader(output)
	if err != nil {
		t.Error(err)
	}

	if output.String() != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
	}
}

func TestLengthWriteHeader3(t *testing.T) {
	l := length{
		perFile:           false,
		writeDescriptions: false,
		writeFileNames:    false,
		lenFormat:         "mb",
	}
	output := bytes.NewBuffer(make([]byte, 0))
	expected := "record\tlength_mb\n"

	err := l.writeHeader(output)
	if err != nil {
		t.Error(err)
	}

	if output.String() != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
	}
}

func TestLengthWriteHeader4(t *testing.T) {
	l := length{
		perFile:           false,
		writeDescriptions: false,
		writeFileNames:    false,
		lenFormat:         "gb",
	}
	output := bytes.NewBuffer(make([]byte, 0))
	expected := "record\tlength_gb\n"

	err := l.writeHeader(output)
	if err != nil {
		t.Error(err)
	}

	if output.String() != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
	}
}

func TestLengthWriteHeader5(t *testing.T) {
	l := length{
		perFile:           true,
		writeDescriptions: false,
		writeFileNames:    false,
		lenFormat:         "",
	}
	output := bytes.NewBuffer(make([]byte, 0))
	expected := "file\tlength\n"

	err := l.writeHeader(output)
	if err != nil {
		t.Error(err)
	}

	if output.String() != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
	}
}

func TestLengthWriteHeader6(t *testing.T) {
	l := length{
		perFile:           false,
		writeDescriptions: false,
		writeFileNames:    true,
		lenFormat:         "",
	}
	output := bytes.NewBuffer(make([]byte, 0))
	expected := "file\trecord\tlength\n"

	err := l.writeHeader(output)
	if err != nil {
		t.Error(err)
	}

	if output.String() != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
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
	reader := NewReader(bytes.NewReader(fastaFile))
	expected := `Seq1	6
Seq2	6
`
	output := bytes.NewBuffer(make([]byte, 0))

	err := lengthRecords("stdin", reader, l, output)
	if err != nil {
		t.Error(err)
	}

	if output.String() != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
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
	reader := NewReader(bytes.NewReader(fastaFile))
	expected := `stdin	12
`
	output := bytes.NewBuffer(make([]byte, 0))

	err := lengthRecords("stdin", reader, l, output)
	if err != nil {
		t.Error(err)
	}

	if output.String() != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
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
	reader := NewReader(bytes.NewReader(fastaFile))
	expected := `Seq1 Homo_sapiens	6
Seq2 Danio_rerio	6
`
	output := bytes.NewBuffer(make([]byte, 0))

	err := lengthRecords("/path/to/test.fa", reader, l, output)
	if err != nil {
		t.Error(err)
	}

	if output.String() != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
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
	reader := NewReader(bytes.NewReader(fastaFile))
	expected := `test.fa	Seq1 Homo_sapiens	6
test.fa	Seq2 Danio_rerio	6
`
	output := bytes.NewBuffer(make([]byte, 0))

	err := lengthRecords("/path/to/test.fa", reader, l, output)
	if err != nil {
		t.Error(err)
	}

	if output.String() != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
	}
}
