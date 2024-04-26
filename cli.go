package main

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:               "fastats {command}",
		Short:             "Very simple statistics from fasta files",
		Long:              ``,
		Version:           "0.6.0",
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	}
)

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var f bool
var c bool
var p string
var d bool

func init() {
	rootCmd.PersistentFlags().BoolVarP(&f, "file", "f", false, "calculate statistics per file (default is per record)")
	rootCmd.PersistentFlags().BoolVarP(&c, "count", "c", false, "print counts (default is proportions)")
	rootCmd.PersistentFlags().BoolVarP(&d, "description", "d", false, "write record descriptions (default is IDs)")

	rootCmd.PersistentFlags().Lookup("file").NoOptDefVal = "true"
	rootCmd.PersistentFlags().Lookup("count").NoOptDefVal = "true"
	rootCmd.PersistentFlags().Lookup("description").NoOptDefVal = "true"

	rootCmd.AddCommand(atCmd)
	rootCmd.AddCommand(gcCmd)
	rootCmd.AddCommand(atgcCmd)
	rootCmd.AddCommand(nCmd)
	rootCmd.AddCommand(gapCmd)
	rootCmd.AddCommand(lenCmd)
	rootCmd.AddCommand(softCmd)
	rootCmd.AddCommand(patternCmd)
	rootCmd.AddCommand(numCmd)
	rootCmd.AddCommand(nameCmd)

	patternCmd.Flags().StringVarP(&p, "pattern", "p", "", "arbitrary pattern to parse")
}

var atCmd = &cobra.Command{
	Use:                   "at <infile[s]>",
	Short:                 "AT content",
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		err = pattern(os.Stdout, args, "ATat", f, c, d)
		return err
	},
}

var gcCmd = &cobra.Command{
	Use:                   "gc <infile[s]>",
	Short:                 "GC content",
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		err = pattern(os.Stdout, args, "GCgc", f, c, d)
		return err
	},
}

var atgcCmd = &cobra.Command{
	Use:                   "atgc <infile[s]>",
	Short:                 "ATGC content",
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		err = pattern(os.Stdout, args, "ATGCatgc", f, c, d)
		return err
	},
}

var nCmd = &cobra.Command{
	Use:                   "n <infile[s]>",
	Short:                 "N content",
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		err = pattern(os.Stdout, args, "Nn", f, c, d)
		return err
	},
}

var gapCmd = &cobra.Command{
	Use:                   "gaps <infile[s]>",
	Short:                 "Gap content",
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		err = pattern(os.Stdout, args, "-", f, c, d)
		return err
	},
}

var softCmd = &cobra.Command{
	Use:                   "soft <infile[s]>",
	Short:                 "Softmasked content",
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		err = pattern(os.Stdout, args, "atgcn", f, c, d)
		return err
	},
}

var patternCmd = &cobra.Command{
	Use: "pattern -p PATTERN <infile[s]>",
	Long: `e.g. fastats pattern -p AG <infile[s]>
`,
	Short:                 "Arbitrary PATTERN content",
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		err = pattern(os.Stdout, args, p, f, c, d)
		return err
	},
}

var lenCmd = &cobra.Command{
	Use:                   "len <infile[s]>",
	Short:                 "Sequence length",
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		err = length(os.Stdout, args, p, f, c, d)
		return err
	},
}

var numCmd = &cobra.Command{
	Use:                   "num <infile[s]>",
	Short:                 "Number of records",
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		err = num(os.Stdout, args, p, f, c, d)
		return err
	},
}

var nameCmd = &cobra.Command{
	Use:                   "names <infile[s]>",
	Short:                 "Record names",
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		err = names(os.Stdout, args, p, f, c, d)
		return err
	},
}
