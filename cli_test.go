package main

import (
	"bytes"
	"testing"

	"github.com/spf13/cobra"
)

var (
	buf = new(bytes.Buffer)
)

func init() {
	out = buf
}

// borrowed from https://github.com/spf13/cobra/blob/6dec1ae26659a130bdb4c985768d1853b0e1bc06/command_test.go
func executeCommand(root *cobra.Command, args ...string) (output string, err error) {
	_, output, err = executeCommandC(root, args...)
	return output, err
}

// modified from https://github.com/spf13/cobra/blob/6dec1ae26659a130bdb4c985768d1853b0e1bc06/command_test.go
func executeCommandC(root *cobra.Command, args ...string) (c *cobra.Command, output string, err error) {
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs(args)

	c, err = root.ExecuteC()

	out := buf.String()
	buf.Reset()

	return c, out, err
}

// In the program flags are shared between subcommands, which is ok because only
// one subcommand is ever called. But for testing, they need resetting to their
// default values between test function calls
func resetFlags() {
	f = false
	c = false
	d = false

	cs = make([]string, 0)
	ps = make([]string, 0)

	fn = false

	kb = false
	mb = false
	gb = false

	nX = make([]int, 0)
	ngX = make([]int, 0)
	lX = make([]int, 0)
	gS = 0
}

func TestContentCmd1(t *testing.T) {
	output, err := executeCommand(rootCmd, "content", "testdata/n2b6.fasta")
	if err != nil {
		t.Fatal(err)
	}

	expected := `record	GCgc_prop	ATat_prop	Nn_prop	-_prop
Seq1	0.333333	0.666667	0.000000	0.000000
Seq2	0.000000	0.833333	0.000000	0.166667
`
	if output != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
	}

	resetFlags()
}

func TestContentCmd2(t *testing.T) {
	output, err := executeCommand(rootCmd, "content", "testdata/n2b6.fasta", "testdata/n2b6.fasta")
	if err != nil {
		t.Fatal(err)
	}

	expected := `file	record	GCgc_prop	ATat_prop	Nn_prop	-_prop
n2b6.fasta	Seq1	0.333333	0.666667	0.000000	0.000000
n2b6.fasta	Seq2	0.000000	0.833333	0.000000	0.166667
n2b6.fasta	Seq1	0.333333	0.666667	0.000000	0.000000
n2b6.fasta	Seq2	0.000000	0.833333	0.000000	0.166667
`
	if output != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
	}

	resetFlags()
}

func TestContentCmd3(t *testing.T) {
	output, err := executeCommand(rootCmd, "content", "testdata/n2b6.fasta", "testdata/n2b6.fasta", "-f")
	if err != nil {
		t.Fatal(err)
	}

	expected := `file	GCgc_prop	ATat_prop	Nn_prop	-_prop
n2b6.fasta	0.166667	0.750000	0.000000	0.083333
n2b6.fasta	0.166667	0.750000	0.000000	0.083333
`
	if output != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
	}

	resetFlags()
}

func TestATCmd1(t *testing.T) {
	output, err := executeCommand(rootCmd, "at", "testdata/n2b6.fasta")
	if err != nil {
		t.Fatal(err)
	}

	expected := `record	ATat_prop
Seq1	0.666667
Seq2	0.833333
`
	if output != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
	}

	resetFlags()
}

func TestATCmd2(t *testing.T) {
	output, err := executeCommand(rootCmd, "at", "testdata/n2b6.fasta", "-f")
	if err != nil {
		t.Fatal(err)
	}

	expected := `file	ATat_prop
n2b6.fasta	0.750000
`
	if output != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
	}

	resetFlags()
}

func TestATCmd3(t *testing.T) {
	output, err := executeCommand(rootCmd, "at", "testdata/n2b6.fasta", "-c")
	if err != nil {
		t.Fatal(err)
	}

	expected := `record	ATat_count
Seq1	4
Seq2	5
`
	if output != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
	}

	resetFlags()
}

func TestATCmd4(t *testing.T) {
	output, err := executeCommand(rootCmd, "at", "testdata/n2b6.fasta", "-f", "-c")
	if err != nil {
		t.Fatal(err)
	}

	expected := `file	ATat_count
n2b6.fasta	9
`
	if output != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
	}

	resetFlags()
}

func TestATCmd5(t *testing.T) {
	output, err := executeCommand(rootCmd, "at", "testdata/n2b6.fasta", "--fn", "-c")
	if err != nil {
		t.Fatal(err)
	}

	expected := `file	record	ATat_count
n2b6.fasta	Seq1	4
n2b6.fasta	Seq2	5
`
	if output != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
	}

	resetFlags()
}

func TestATCmd6(t *testing.T) {
	output, err := executeCommand(rootCmd, "at", "testdata/n2b6.fasta", "--fn", "-c", "-d")
	if err != nil {
		t.Fatal(err)
	}

	expected := `file	record	ATat_count
n2b6.fasta	Seq1 Hsapiens	4
n2b6.fasta	Seq2 Ptroglodytes	5
`
	if output != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
	}

	resetFlags()
}

func TestGCCmd1(t *testing.T) {
	output, err := executeCommand(rootCmd, "gc", "testdata/n2b6.fasta")
	if err != nil {
		t.Fatal(err)
	}

	expected := `record	GCgc_prop
Seq1	0.333333
Seq2	0.000000
`
	if output != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
	}

	resetFlags()
}

func TestGCCmd2(t *testing.T) {
	output, err := executeCommand(rootCmd, "gc", "testdata/n2b6.fasta", "-c")
	if err != nil {
		t.Fatal(err)
	}

	expected := `record	GCgc_count
Seq1	2
Seq2	0
`
	if output != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
	}

	resetFlags()
}

func TestATGCCmd1(t *testing.T) {
	output, err := executeCommand(rootCmd, "atgc", "testdata/n2b6.fasta")
	if err != nil {
		t.Fatal(err)
	}

	expected := `record	ATGCatgc_prop
Seq1	1.000000
Seq2	0.833333
`
	if output != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
	}

	resetFlags()
}

func TestNCmd1(t *testing.T) {
	output, err := executeCommand(rootCmd, "n", "testdata/n2b6.fasta")
	if err != nil {
		t.Fatal(err)
	}

	expected := `record	Nn_prop
Seq1	0.000000
Seq2	0.000000
`
	if output != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
	}

	resetFlags()
}

func TestGapsCmd1(t *testing.T) {
	output, err := executeCommand(rootCmd, "gaps", "testdata/n2b6.fasta")
	if err != nil {
		t.Fatal(err)
	}

	expected := `record	-_prop
Seq1	0.000000
Seq2	0.166667
`
	if output != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
	}

	resetFlags()
}

func TestSoftCmd1(t *testing.T) {
	output, err := executeCommand(rootCmd, "soft", "testdata/n2b6.fasta")
	if err != nil {
		t.Fatal(err)
	}

	expected := `record	atgcn_prop
Seq1	0.000000
Seq2	0.000000
`
	if output != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
	}

	resetFlags()
}

func TestLenCmd1(t *testing.T) {
	output, err := executeCommand(rootCmd, "len", "testdata/n2b6.fasta")
	if err != nil {
		t.Fatal(err)
	}

	expected := `record	length
Seq1	6
Seq2	6
`
	if output != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
	}

	resetFlags()
}

func TestLenCmd2(t *testing.T) {
	output, err := executeCommand(rootCmd, "len", "testdata/n2b6.fasta", "-f")
	if err != nil {
		t.Fatal(err)
	}

	expected := `file	length
n2b6.fasta	12
`
	if output != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
	}

	resetFlags()
}

func TestLenCmd3(t *testing.T) {
	output, err := executeCommand(rootCmd, "len", "testdata/n2b6.fasta", "testdata/n2b6.fasta", "-f")
	if err != nil {
		t.Fatal(err)
	}

	expected := `file	length
n2b6.fasta	12
n2b6.fasta	12
`
	if output != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
	}

	resetFlags()
}

func TestLenCmd4(t *testing.T) {
	output, err := executeCommand(rootCmd, "len", "testdata/n2b6.fasta", "testdata/n2b6.fasta", "-f", "--kb")
	if err != nil {
		t.Fatal(err)
	}

	expected := `file	length_kb
n2b6.fasta	0.012
n2b6.fasta	0.012
`
	if output != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
	}

	resetFlags()
}

func TestLenCmd5(t *testing.T) {
	output, err := executeCommand(rootCmd, "len", "testdata/n2b6.fasta", "testdata/n2b6.fasta", "-f", "--mb")
	if err != nil {
		t.Fatal(err)
	}

	expected := `file	length_mb
n2b6.fasta	0.000012
n2b6.fasta	0.000012
`
	if output != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
	}

	resetFlags()
}

func TestLenCmd6(t *testing.T) {
	output, err := executeCommand(rootCmd, "len", "testdata/n2b6.fasta", "testdata/n2b6.fasta", "-f", "--gb")
	if err != nil {
		t.Fatal(err)
	}

	expected := `file	length_gb
n2b6.fasta	0.000000012
n2b6.fasta	0.000000012
`
	if output != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
	}

	resetFlags()
}

func TestLenCmd7(t *testing.T) {
	_, err := executeCommand(rootCmd, "len", "testdata/n2b6.fasta", "testdata/n2b6.fasta", "-f", "--kb", "--gb")
	if err.Error() != "flags --kb, --mb, and --gb are mutually exclusive" {
		t.Fatal(err)
	}

	resetFlags()
}

func TestNumCmd1(t *testing.T) {
	output, err := executeCommand(rootCmd, "num", "testdata/n2b6.fasta")
	if err != nil {
		t.Fatal(err)
	}

	expected := `file	n_records
n2b6.fasta	2
`
	if output != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
	}

	resetFlags()
}

func TestNumCmd2(t *testing.T) {
	output, err := executeCommand(rootCmd, "num", "testdata/n2b6.fasta", "testdata/n2b6.fasta")
	if err != nil {
		t.Fatal(err)
	}

	expected := `file	n_records
n2b6.fasta	2
n2b6.fasta	2
`
	if output != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
	}

	resetFlags()
}

func TestNamesCmd1(t *testing.T) {
	output, err := executeCommand(rootCmd, "names", "testdata/n2b6.fasta")
	if err != nil {
		t.Fatal(err)
	}

	expected := `file	id
n2b6.fasta	Seq1
n2b6.fasta	Seq2
`
	if output != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
	}

	resetFlags()
}

func TestNamesCmd2(t *testing.T) {
	output, err := executeCommand(rootCmd, "names", "testdata/n2b6.fasta", "-d")
	if err != nil {
		t.Fatal(err)
	}

	expected := `file	description
n2b6.fasta	Seq1 Hsapiens
n2b6.fasta	Seq2 Ptroglodytes
`
	if output != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
	}

	resetFlags()
}

func TestNamesCmd3(t *testing.T) {
	output, err := executeCommand(rootCmd, "names", "testdata/n2b6.fasta", "testdata/n2b6.fasta", "-d")
	if err != nil {
		t.Fatal(err)
	}

	expected := `file	description
n2b6.fasta	Seq1 Hsapiens
n2b6.fasta	Seq2 Ptroglodytes
n2b6.fasta	Seq1 Hsapiens
n2b6.fasta	Seq2 Ptroglodytes
`
	if output != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
	}

	resetFlags()
}

func TestAssemblyCmd1(t *testing.T) {
	output, err := executeCommand(rootCmd, "assembly", "testdata/asmbly.fasta")
	if err != nil {
		t.Fatal(err)
	}

	expected := `file	N50	N90	L50	L90
asmbly.fasta	8	4	3	7
`
	if output != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
	}

	resetFlags()
}

func TestAssemblyCmd2(t *testing.T) {
	output, err := executeCommand(rootCmd, "assembly", "testdata/asmbly.fasta", "-N50")
	if err != nil {
		t.Fatal(err)
	}

	expected := `file	N50
asmbly.fasta	8
`
	if output != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
	}

	resetFlags()
}

func TestAssemblyCmd3(t *testing.T) {
	output, err := executeCommand(rootCmd, "assembly", "testdata/asmbly.fasta", "-N50", "--kb")
	if err != nil {
		t.Fatal(err)
	}

	expected := `file	N50_kb
asmbly.fasta	0.008
`
	if output != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
	}

	resetFlags()
}

func TestAssemblyCmd4(t *testing.T) {
	output, err := executeCommand(rootCmd, "assembly", "testdata/asmbly.fasta", "-N50", "--mb")
	if err != nil {
		t.Fatal(err)
	}

	expected := `file	N50_mb
asmbly.fasta	0.000008
`
	if output != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
	}

	resetFlags()
}

func TestAssemblyCmd5(t *testing.T) {
	output, err := executeCommand(rootCmd, "assembly", "testdata/asmbly.fasta", "-N50", "--gb")
	if err != nil {
		t.Fatal(err)
	}

	expected := `file	N50_gb
asmbly.fasta	0.000000008
`
	if output != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
	}

	resetFlags()
}

func TestAssemblyCmd6(t *testing.T) {
	_, err := executeCommand(rootCmd, "assembly", "testdata/asmbly.fasta", "-N50", "--kb", "--gb")
	if err.Error() != "flags --kb, --mb, and --gb are mutually exclusive" {
		t.Fatal(err)
	}

	resetFlags()
}

func TestAssemblyCmd7(t *testing.T) {
	output, err := executeCommand(rootCmd, "assembly", "testdata/asmbly.fasta", "testdata/asmbly.fasta", "-N50")
	if err != nil {
		t.Fatal(err)
	}

	expected := `file	N50
asmbly.fasta	8
asmbly.fasta	8
`
	if output != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
	}

	resetFlags()
}

func TestAssemblyCmd8(t *testing.T) {
	output, err := executeCommand(rootCmd, "assembly", "testdata/asmbly.fasta", "-N50", "-N90")
	if err != nil {
		t.Fatal(err)
	}

	expected := `file	N50	N90
asmbly.fasta	8	4
`
	if output != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
	}

	resetFlags()
}

func TestAssemblyCmd9(t *testing.T) {
	output, err := executeCommand(rootCmd, "assembly", "testdata/asmbly.fasta", "-N50", "-G50", "-g54")
	if err != nil {
		t.Fatal(err)
	}

	expected := `file	N50	NG50
asmbly.fasta	8	8
`
	if output != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
	}

	resetFlags()
}

func TestAssemblyCmd10(t *testing.T) {
	output, err := executeCommand(rootCmd, "assembly", "testdata/asmbly.fasta", "-G50", "-g80")
	if err != nil {
		t.Fatal(err)
	}

	expected := `file	NG50
asmbly.fasta	6
`
	if output != expected {
		t.Errorf("expected:\n %s, got:\n %s", expected, output)
	}

	resetFlags()
}
