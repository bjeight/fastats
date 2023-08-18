package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:               "fastats [command] <infiles>",
		Short:             "Very simple statistics from fasta files",
		Long:              ``,
		Version:           "0.1.0",
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}

var p string
var f bool
var c bool

func init() {
	rootCmd.PersistentFlags().BoolVarP(&f, "file", "f", true, "calculate statists per file (default is per record)")
	rootCmd.PersistentFlags().BoolVarP(&c, "count", "c", false, "print counts (default is proportions)")

	rootCmd.PersistentFlags().Lookup("file").NoOptDefVal = "false"
	rootCmd.PersistentFlags().Lookup("file").DefValue = "false"

	rootCmd.AddCommand(atCmd)
	rootCmd.AddCommand(gcCmd)
	rootCmd.AddCommand(atgcCmd)
	rootCmd.AddCommand(nCmd)
	rootCmd.AddCommand(gapCmd)
	rootCmd.AddCommand(lenCmd)
	rootCmd.AddCommand(softCmd)
	rootCmd.AddCommand(patternCmd)
	patternCmd.Flags().StringVarP(&p, "pattern", "p", "", "arbitrary pattern to parse")
}

var atCmd = &cobra.Command{
	Use:                   "at <infile>",
	Args:                  cobra.ExactArgs(1),
	Short:                 "AT content",
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		err = pattern(args[0], "ATat", f, c)
		return err
	},
}

var gcCmd = &cobra.Command{
	Use:                   "gc <infile>",
	Args:                  cobra.ExactArgs(1),
	Short:                 "GC content",
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := pattern(args[0], "GCgc", f, c)
		return err
	},
}

var atgcCmd = &cobra.Command{
	Use:                   "atgc <infile>",
	Args:                  cobra.ExactArgs(1),
	Short:                 "ATGC content",
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := pattern(args[0], "ATGCatgc", f, c)
		return err
	},
}

var nCmd = &cobra.Command{
	Use:                   "n <infile>",
	Args:                  cobra.ExactArgs(1),
	Short:                 "N content",
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := pattern(args[0], "Nn", f, c)
		return err
	},
}

var gapCmd = &cobra.Command{
	Use:                   "gaps <infile>",
	Args:                  cobra.ExactArgs(1),
	Short:                 "Gap content",
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := pattern(args[0], "-", f, c)
		return err
	},
}

var lenCmd = &cobra.Command{
	Use:                   "len <infile>",
	Args:                  cobra.ExactArgs(1),
	Short:                 "Sequence length",
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := length(args[0], f, c)
		return err
	},
}

var softCmd = &cobra.Command{
	Use:                   "soft <infile>",
	Args:                  cobra.ExactArgs(1),
	Short:                 "Softmasked content",
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := pattern(args[0], "atgcn", f, c)
		return err
	},
}

var patternCmd = &cobra.Command{
	Use: "pattern -p PATTERN <infile>",
	Long: `e.g. fastats pattern -p AG <infile>
`,
	Args:                  cobra.ExactArgs(1),
	Short:                 "PATTERN content",
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := pattern(args[0], p, f, c)
		return err
	},
}
