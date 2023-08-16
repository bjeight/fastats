package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/virus-evolution/gofasta/pkg/encoding"
	"github.com/virus-evolution/gofasta/pkg/fastaio"
)

func main() {
	infile := os.Args[1]
	f, err := os.Open(infile)
	if err != nil {
		panic(err)
	}

	c_fasta := make(chan fastaio.EncodedFastaRecord)
	c_err := make(chan error)
	c_done := make(chan bool)

	go ReadEncodeFasta(f, false, c_fasta, c_err, c_done)

	go func() {
		<-c_done
		close(c_fasta)
	}()

	for record := range c_fasta {
		var lookup [256]int
		for _, nuc := range record.Seq {
			lookup[nuc] += 1
		}
		total := 0
		for _, v := range lookup {
			total += v
		}
		atgc_total := lookup[136] + lookup[72] + lookup[40] + lookup[24]
		prop := float64(atgc_total) / float64(total)
		fmt.Printf("%s proportion ATGC: %f\n", record.ID, prop)
	}
}

func ReadEncodeFasta(f io.Reader, hardGaps bool, chnl chan fastaio.EncodedFastaRecord, cErr chan error, cDone chan bool) {

	var err error

	var coding [256]byte
	switch hardGaps {
	case true:
		coding = encoding.MakeEncodingArrayHardGaps()
	case false:
		coding = encoding.MakeEncodingArray()
	}

	s := bufio.NewScanner(f)
	s.Buffer(make([]byte, 0), 1024*1024)

	first := true

	var id string
	var description string
	var seqBuffer []byte
	var line []byte
	var nuc byte

	var fr fastaio.EncodedFastaRecord

	counter := 0

	for s.Scan() {
		line = s.Bytes()

		if first {

			if len(line) == 0 || line[0] != '>' {
				cErr <- errors.New("badly formatted fasta file")
				return
			}

			description = string(line[1:])
			id = strings.Fields(description)[0]

			first = false

		} else if line[0] == '>' {

			fr = fastaio.EncodedFastaRecord{ID: id, Description: description, Seq: seqBuffer, Idx: counter}
			chnl <- fr
			counter++

			description = string(line[1:])
			id = strings.Fields(description)[0]
			seqBuffer = make([]byte, 0)

		} else {
			encodedLine := make([]byte, len(line))
			for i := range line {
				nuc = coding[line[i]]
				if nuc == 0 {
					cErr <- fmt.Errorf("invalid nucleotide in fasta file (\"%s\")", string(line[i]))
					return
				}
				encodedLine[i] = nuc
			}
			seqBuffer = append(seqBuffer, encodedLine...)
		}
	}

	if len(seqBuffer) > 0 {
		fr = fastaio.EncodedFastaRecord{ID: id, Description: description, Seq: seqBuffer, Idx: counter}
		chnl <- fr
		counter++
	}

	if counter == 0 {
		cErr <- errors.New("empty fasta file")
		return
	}

	err = s.Err()
	if err != nil {
		cErr <- err
		return
	}

	cDone <- true
}
