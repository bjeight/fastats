package main

import (
	"errors"
	"io"
	"slices"
	"strconv"
)

type assembly struct {
	inputs     []string            // input files
	stats      []assemblyStatistic //
	genomeSize int64               //genome size to use when calculating NG-statistics
	lenFormat  string
}

type assemblyStatistic struct {
	sType  string // the type of statistic to calculate: {N, NG, L}
	sValue int64  // the value of the statistic to calculate (50, 90, etc.)
}

func (stat assemblyStatistic) string() string {
	return stat.sType + strconv.FormatInt(stat.sValue, 10)

}

func (args assembly) writeHeader(w io.Writer) error {
	_, err := w.Write([]byte("file\tn_records\tlength"))
	if err != nil {
		return err
	}
	switch args.lenFormat {
	case "kb", "mb", "gb":
		_, err := w.Write([]byte("_" + args.lenFormat))
		if err != nil {
			return err
		}
	}
	for _, stat := range args.stats {
		_, err := w.Write([]byte("\t" + stat.string()))
		if err != nil {
			return err
		}
		switch stat.sType {
		case "N", "NG":
			switch args.lenFormat {
			case "kb", "mb", "gb":
				_, err := w.Write([]byte("_" + args.lenFormat))
				if err != nil {
					return err
				}
			}
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

func assemblyRecords(inputPath string, r *Reader, args assembly, w io.Writer) error {
	contigLengths := make([]int64, 0)
	var totalLength int64 = 0
	var nRecords int64 = 0
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		l := record.Len
		contigLengths = append(contigLengths, l)
		totalLength += l
		nRecords++
	}
	slices.Sort(contigLengths)

	_, err := w.Write([]byte(returnFileName(inputPath) + "\t" + strconv.FormatInt(nRecords, 10) + "\t" + returnLengthFormatted(totalLength, args.lenFormat)))
	if err != nil {
		return err
	}

	for _, stat := range args.stats {
		switch stat.sType {
		case "N":
			nX := nStat(contigLengths, totalLength, stat.sValue)
			_, err := w.Write([]byte("\t" + returnLengthFormatted(nX, args.lenFormat)))
			if err != nil {
				return err
			}
		case "NG":
			nX := nStat(contigLengths, args.genomeSize, stat.sValue)
			_, err := w.Write([]byte("\t" + returnLengthFormatted(nX, args.lenFormat)))
			if err != nil {
				return err
			}
		case "L":
			lX := lStat(contigLengths, totalLength, stat.sValue)
			_, err := w.Write([]byte("\t" + strconv.FormatInt(lX, 10)))
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
func nStat(contigLengths []int64, genomeSize int64, statValue int64) int64 {
	var nX int64
	var runningTotal int64 = 0
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
func lStat(contigLengths []int64, genomeSize int64, statValue int64) int64 {
	var lX int64
	var runningTotal int64 = 0
	for i := len(contigLengths) - 1; i >= 0; i-- {
		runningTotal += contigLengths[i]
		if float64(runningTotal) >= float64(genomeSize)*(float64(statValue)/100.0) {
			lX = int64(len(contigLengths) - i)
			break
		}
	}
	return lX
}
