package main

import (
	"errors"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:               "fastats {command}",
		Short:             "Very simple statistics from fasta files",
		Long:              ``,
		Version:           "0.8.0",
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	}
)

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var (
	f bool
	c bool
	d bool

	fn bool

	b string

	kb bool
	mb bool
	gb bool
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&f, "file", "f", false, "calculate statistics per file (default is per record)")
	rootCmd.PersistentFlags().BoolVarP(&c, "count", "c", false, "print base content counts (default is proportions)")
	rootCmd.PersistentFlags().BoolVarP(&d, "description", "d", false, "print record descriptions (default is IDs)")
	rootCmd.PersistentFlags().BoolVarP(&fn, "fn", "", false, "always print a filename column")

	rootCmd.PersistentFlags().Lookup("file").NoOptDefVal = "true"
	rootCmd.PersistentFlags().Lookup("count").NoOptDefVal = "true"
	rootCmd.PersistentFlags().Lookup("description").NoOptDefVal = "true"
	rootCmd.PersistentFlags().Lookup("fn").NoOptDefVal = "true"

	rootCmd.AddCommand(atCmd)
	rootCmd.AddCommand(gcCmd)
	rootCmd.AddCommand(atgcCmd)
	rootCmd.AddCommand(nCmd)
	rootCmd.AddCommand(gapCmd)
	rootCmd.AddCommand(lenCmd)
	rootCmd.AddCommand(softCmd)
	rootCmd.AddCommand(contentCmd)
	rootCmd.AddCommand(numCmd)
	rootCmd.AddCommand(nameCmd)

	contentCmd.Flags().StringVarP(&b, "bases", "b", "", "arbitrary base content to count (case-sensitive)")

	lenCmd.Flags().BoolVarP(&kb, "kb", "", false, "print sequence lengths in kilobases")
	lenCmd.Flags().BoolVarP(&mb, "mb", "", false, "print sequence lengths in megabases")
	lenCmd.Flags().BoolVarP(&gb, "gb", "", false, "print sequence lengths in gigabases")

	lenCmd.Flags().Lookup("kb").NoOptDefVal = "true"
	lenCmd.Flags().Lookup("mb").NoOptDefVal = "true"
	lenCmd.Flags().Lookup("gb").NoOptDefVal = "true"
}

func resolveCommandLine(files []string) ([]string, error) {
	if len(files) > 1 {
		fn = true
	}
	if len(files) == 0 {
		files = []string{"stdin"}
	}
	return files, nil
}

var atCmd = &cobra.Command{
	Use:                   "at <infile[s]>",
	Short:                 "AT content",
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, files []string) (err error) {
		files, err = resolveCommandLine(files)
		if err != nil {
			return err
		}
		p := pattern{
			inputs:            files,
			perFile:           f,
			writeCounts:       c,
			writeDescriptions: d,
			writeFileNames:    fn,
			bases:             "ATat",
		}
		err = p.writeHeader(os.Stdout)
		if err != nil {
			return err
		}
		err = p.writeBody(os.Stdout)
		return err
	},
}

var gcCmd = &cobra.Command{
	Use:                   "gc <infile[s]>",
	Short:                 "GC content",
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, files []string) (err error) {
		files, err = resolveCommandLine(files)
		if err != nil {
			return err
		}
		p := pattern{
			inputs:            files,
			perFile:           f,
			writeCounts:       c,
			writeDescriptions: d,
			writeFileNames:    fn,
			bases:             "GCgc",
		}
		err = p.writeHeader(os.Stdout)
		if err != nil {
			return err
		}
		err = p.writeBody(os.Stdout)
		return err
	},
}

var atgcCmd = &cobra.Command{
	Use:                   "atgc <infile[s]>",
	Short:                 "ATGC content",
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, files []string) (err error) {
		files, err = resolveCommandLine(files)
		if err != nil {
			return err
		}
		p := pattern{
			inputs:            files,
			perFile:           f,
			writeCounts:       c,
			writeDescriptions: d,
			writeFileNames:    fn,
			bases:             "ATGCatgc",
		}
		err = p.writeHeader(os.Stdout)
		if err != nil {
			return err
		}
		err = p.writeBody(os.Stdout)
		return err
	},
}

var nCmd = &cobra.Command{
	Use:                   "n <infile[s]>",
	Short:                 "N content",
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, files []string) (err error) {
		files, err = resolveCommandLine(files)
		if err != nil {
			return err
		}
		p := pattern{
			inputs:            files,
			perFile:           f,
			writeCounts:       c,
			writeDescriptions: d,
			writeFileNames:    fn,
			bases:             "Nn",
		}
		err = p.writeHeader(os.Stdout)
		if err != nil {
			return err
		}
		err = p.writeBody(os.Stdout)
		return err
	},
}

var gapCmd = &cobra.Command{
	Use:                   "gaps <infile[s]>",
	Short:                 "Gap content",
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, files []string) (err error) {
		files, err = resolveCommandLine(files)
		if err != nil {
			return err
		}
		p := pattern{
			inputs:            files,
			perFile:           f,
			writeCounts:       c,
			writeDescriptions: d,
			writeFileNames:    fn,
			bases:             "-",
		}
		err = p.writeHeader(os.Stdout)
		if err != nil {
			return err
		}
		err = p.writeBody(os.Stdout)
		return err
	},
}

var softCmd = &cobra.Command{
	Use:                   "soft <infile[s]>",
	Short:                 "Softmasked content",
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, files []string) (err error) {
		files, err = resolveCommandLine(files)
		if err != nil {
			return err
		}
		p := pattern{
			inputs:            files,
			perFile:           f,
			writeCounts:       c,
			writeDescriptions: d,
			writeFileNames:    fn,
			bases:             "atgcn",
		}
		err = p.writeHeader(os.Stdout)
		if err != nil {
			return err
		}
		err = p.writeBody(os.Stdout)
		return err
	},
}

var contentCmd = &cobra.Command{
	Use:   "content -b BASES <infile[s]>",
	Short: "Arbitrary base content",
	Long: `e.g. fastats content -b AG <infile[s]>
`,
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, files []string) (err error) {
		files, err = resolveCommandLine(files)
		if err != nil {
			return err
		}
		p := pattern{
			inputs:            files,
			perFile:           f,
			writeCounts:       c,
			writeDescriptions: d,
			writeFileNames:    fn,
			bases:             b,
		}
		err = p.writeHeader(os.Stdout)
		if err != nil {
			return err
		}
		err = p.writeBody(os.Stdout)
		return err
	},
}

var lenCmd = &cobra.Command{
	Use:                   "len <infile[s]>",
	Short:                 "Sequence length",
	DisableFlagsInUseLine: true,
	Aliases:               []string{"length"},
	RunE: func(cmd *cobra.Command, files []string) (err error) {
		files, err = resolveCommandLine(files)
		if err != nil {
			return err
		}
		formatCount := 0
		lenFormat := "b"
		if kb {
			formatCount += 1
			lenFormat = "kb"
		}
		if mb {
			formatCount += 1
			lenFormat = "mb"
		}
		if gb {
			formatCount += 1
			lenFormat = "gb"
		}
		if formatCount > 1 {
			return errors.New("Choose one of --kb, --mb, or --gb")
		}
		l := length{
			inputs:            files,
			perFile:           f,
			writeDescriptions: d,
			writeFileNames:    fn,
			lenFormat:         lenFormat,
		}
		err = l.writeHeader(os.Stdout)
		if err != nil {
			return err
		}
		err = l.writeBody(os.Stdout)
		return err
	},
}

var numCmd = &cobra.Command{
	Use:                   "num <infile[s]>",
	Short:                 "Number of records",
	DisableFlagsInUseLine: true,
	Aliases:               []string{"number"},
	RunE: func(cmd *cobra.Command, files []string) (err error) {
		files, err = resolveCommandLine(files)
		if err != nil {
			return err
		}
		n := num{
			inputs: files,
		}
		err = n.writeHeader(os.Stdout)
		if err != nil {
			return err
		}
		err = n.writeBody(os.Stdout)
		return err
	},
}

var nameCmd = &cobra.Command{
	Use:                   "names <infile[s]>",
	Short:                 "Record names",
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, files []string) (err error) {
		files, err = resolveCommandLine(files)
		if err != nil {
			return err
		}
		n := names{
			inputs:            files,
			writeDescriptions: d,
		}
		err = n.writeHeader(os.Stdout)
		if err != nil {
			return err
		}
		err = n.writeBody(os.Stdout)
		return err
	},
}
