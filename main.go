package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:               "fastats",
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

func init() {
	rootCmd.AddCommand(atCmd)
	rootCmd.AddCommand(gcCmd)
	rootCmd.AddCommand(atgcCmd)
	rootCmd.AddCommand(nCmd)
	rootCmd.AddCommand(gapCmd)
	rootCmd.AddCommand(lenCmd)
	rootCmd.AddCommand(softCmd)
}

var atCmd = &cobra.Command{
	Use:                   "at <infile>",
	Args:                  cobra.ExactArgs(1),
	Short:                 "AT content",
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		err = at(args[0])
		return err
	},
}

var gcCmd = &cobra.Command{
	Use:                   "gc <infile>",
	Args:                  cobra.ExactArgs(1),
	Short:                 "GC content",
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := gc(args[0])
		return err
	},
}

var atgcCmd = &cobra.Command{
	Use:                   "atgc <infile>",
	Args:                  cobra.ExactArgs(1),
	Short:                 "ATGC content",
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := atgc(args[0])
		return err
	},
}

var nCmd = &cobra.Command{
	Use:                   "n <infile>",
	Args:                  cobra.ExactArgs(1),
	Short:                 "N content",
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := n(args[0])
		return err
	},
}

var gapCmd = &cobra.Command{
	Use:                   "gaps <infile>",
	Args:                  cobra.ExactArgs(1),
	Short:                 "Gap content",
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := gap(args[0])
		return err
	},
}

var lenCmd = &cobra.Command{
	Use:                   "len <infile>",
	Args:                  cobra.ExactArgs(1),
	Short:                 "Sequence length",
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := length(args[0])
		return err
	},
}

var softCmd = &cobra.Command{
	Use:                   "soft <infile>",
	Args:                  cobra.ExactArgs(1),
	Short:                 "Softmasked content",
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := soft(args[0])
		return err
	},
}
