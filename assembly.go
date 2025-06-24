package main

import (
	"errors"
	"io"
	"slices"
	"strconv"

	"github.com/bjeight/fastats/fasta"
)

type assembly struct {
	inputs     []string            // input files
	stats      []assemblyStatistic //
	genomeSize int                 //genome size to use when calculating NG-statistics
}

type assemblyStatistic struct {
	sType  string // the type of statistic to calculate: {N, NG, L}
	sValue int    // the value of the statistic to calculate (50, 90, etc.)
}

func (stat assemblyStatistic) string() string {
	return stat.sType + strconv.Itoa(stat.sValue)
}

func (args assembly) writeHeader(w io.Writer) error {
	_, err := w.Write([]byte("file"))
	if err != nil {
		return err
	}
	for _, stat := range args.stats {
		_, err := w.Write([]byte("\t" + stat.string()))
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte("\n"))
	return err
}

func (args assembly) writeBody(w io.Writer) error {
	for _, input := range args.inputs {
		reader, file, err := getReaderFile(input)
		if err != nil {
			return err
		}
		defer file.Close()
		err = assemblyRecords(input, reader, args, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func assemblyRecords(inputPath string, r *fasta.Reader, args assembly, w io.Writer) error {
	contigLengths := make([]int, 0)
	totalLength := 0
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		l := len(record.Seq)
		contigLengths = append(contigLengths, l)
		totalLength += l
	}
	slices.Sort(contigLengths)

	_, err := w.Write([]byte(returnFileName(inputPath)))
	if err != nil {
		return err
	}

	for _, stat := range args.stats {
		switch stat.sType {
		case "N":
			nX := nStat(contigLengths, totalLength, stat.sValue)
			_, err := w.Write([]byte("\t" + strconv.Itoa(nX)))
			if err != nil {
				return err
			}
		case "NG":
			nX := nStat(contigLengths, args.genomeSize, stat.sValue)
			_, err := w.Write([]byte("\t" + strconv.Itoa(nX)))
			if err != nil {
				return err
			}
		case "L":
			lX := lStat(contigLengths, totalLength, stat.sValue)
			_, err := w.Write([]byte("\t" + strconv.Itoa(lX)))
			if err != nil {
				return err
			}
		default:
			return errors.New("unknown assembly statistic type")
		}
	}
	_, err = w.Write([]byte("\n"))
	if err != nil {
		return err
	}

	return nil
}

// contigLengths must be sorted in ascending order
func nStat(contigLengths []int, genomeSize int, statValue int) int {
	var nX int
	runningTotal := 0
	for i := len(contigLengths) - 1; i >= 0; i-- {
		runningTotal += contigLengths[i]
		if float64(runningTotal) >= float64(genomeSize)*(float64(statValue)/100.0) {
			nX = contigLengths[i]
			break
		}
	}
	return nX
}

// contigLengths must be sorted in ascending order
func lStat(contigLengths []int, genomeSize int, statValue int) int {
	var lX int
	runningTotal := 0
	for i := len(contigLengths) - 1; i >= 0; i-- {
		runningTotal += contigLengths[i]
		if float64(runningTotal) >= float64(genomeSize)*(float64(statValue)/100.0) {
			lX = len(contigLengths) - i
			break
		}
	}
	return lX
}
