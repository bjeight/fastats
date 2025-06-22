package fasta

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"errors"
	"io"
)

// A struct for one fasta record
type Record struct {
	ID          string
	Description string
	Seq         []byte
}

var (
	errBadlyFormedFasta = errors.New("badly formed fasta")
)

type Reader struct {
	*bufio.Reader
}

func NewReader(f io.Reader) *Reader {
	return &Reader{bufio.NewReader(f)}
}

func NewZReader(f io.Reader) *Reader {
	rz, _ := gzip.NewReader(f)
	return &Reader{bufio.NewReader(rz)}
}

// Read reads one fasta record from the underlying reader. The final record is returned with error = nil,
// and the next call to Read() returns an empty FastaRecord struct and error = io.EOF.
func (r *Reader) Read() (Record, error) {

	var (
		buffer, line, peek []byte
		fields             [][]byte
		err                error
		record             Record
	)

	first := true

	for {
		if first {
			// "ReadBytes reads until the first occurrence of delim in the input,
			// returning a slice containing the data up to and including the delimiter.
			// If ReadBytes encounters an error before finding a delimiter,
			// it returns the data read before the error and the error itself (often io.EOF).
			// ReadBytes returns err != nil if and only if the returned data does not end in delim.
			// For simple uses, a Scanner may be more convenient."
			line, err = r.ReadBytes('\n')

			// return even if err == io.EOF, because the file should never end on a fasta header line
			if err != nil {
				return Record{}, err

				// if the header doesn't start with a > then something is also wrong
			} else if line[0] != '>' {
				return Record{}, errBadlyFormedFasta
			}

			drop := 0
			// Strip unix or dos newline characters from the header before setting the description.
			if line[len(line)-1] == '\n' {
				drop = 1
				if len(line) > 1 && line[len(line)-2] == '\r' {
					drop = 2
				}
				line = line[:len(line)-drop]
			}

			// split the header on whitespace
			fields = bytes.Fields(line[1:])
			// fasta ID
			record.ID = string(fields[0])
			// fasta description
			record.Description = string(line[1:])

			// we are no longer on a header line
			first = false

		} else {
			// peek at the first next byte of the underlying reader, in order
			// to see if we've reached the end of this record (or the file)
			peek, err = r.Peek(1)

			// both these cases are fine if first = false, so we can exit the loop and return the fasta record
			if err == io.EOF || peek[0] == '>' {
				err = nil
				break

				// other errors are returned along with an empty fasta record
			} else if err != nil {
				return Record{}, err
			}

			// If we've got this far, this should be a sequence line.
			// The err from ReadBytes() may be io.EOF if the file ends before a newline character, but this is okay because it will
			// be caught when we peek in the next iteration of the while loop.
			line, err = r.ReadBytes('\n')
			if err != nil && err != io.EOF {
				return Record{}, err
			}

			drop := 0
			// Strip unix or dos newline characters from the sequence before appending it.
			if line[len(line)-1] == '\n' {
				drop = 1
				if len(line) > 1 && line[len(line)-2] == '\r' {
					drop = 2
				}
				line = line[:len(line)-drop]
			}

			buffer = append(buffer, line...)
		}
	}
	record.Seq = buffer

	return record, err
}
