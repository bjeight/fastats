package main

import (
	"errors"
	"io"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:               "fastats {command}",
		Short:             "Very simple statistics from fasta files",
		Long:              ``,
		Version:           "0.9.1",
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

	cs []string
	ps []string

	fn bool

	kb bool
	mb bool
	gb bool

	nX  []int
	ngX []int
	lX  []int
	gS  int

	out io.Writer = os.Stdout
)

func init() {

	rootCmd.AddCommand(contentCmd)
	rootCmd.AddCommand(atCmd)
	rootCmd.AddCommand(gcCmd)
	rootCmd.AddCommand(atgcCmd)
	rootCmd.AddCommand(nCmd)
	rootCmd.AddCommand(gapCmd)
	rootCmd.AddCommand(softCmd)
	rootCmd.AddCommand(lenCmd)
	rootCmd.AddCommand(numCmd)
	rootCmd.AddCommand(nameCmd)
	rootCmd.AddCommand(assemblyCmd)

	contentCmd.Flags().StringSliceVarP(&cs, "counts", "c", make([]string, 0), "arbitrary base content counts (case-sensitive)")
	contentCmd.Flags().StringSliceVarP(&ps, "proportions", "p", make([]string, 0), "arbitrary base content proportions (case-sensitive)")
	contentCmd.Flags().BoolVarP(&f, "file", "f", false, "calculate statistics per file (default is per record)")
	contentCmd.Flags().BoolVarP(&d, "description", "d", false, "print record descriptions (default is IDs)")
	contentCmd.Flags().BoolVarP(&fn, "fn", "", false, "always print a filename column")
	contentCmd.Flags().Lookup("file").NoOptDefVal = "true"
	contentCmd.Flags().Lookup("description").NoOptDefVal = "true"
	contentCmd.Flags().Lookup("fn").NoOptDefVal = "true"
	contentCmd.Flags().SortFlags = false

	atCmd.Flags().BoolVarP(&c, "count", "c", false, "print base content counts (default is proportions)")
	atCmd.Flags().BoolVarP(&f, "file", "f", false, "calculate statistics per file (default is per record)")
	atCmd.Flags().BoolVarP(&d, "description", "d", false, "print record descriptions (default is IDs)")
	atCmd.Flags().BoolVarP(&fn, "fn", "", false, "always print a filename column")
	atCmd.Flags().Lookup("count").NoOptDefVal = "true"
	atCmd.Flags().Lookup("file").NoOptDefVal = "true"
	atCmd.Flags().Lookup("description").NoOptDefVal = "true"
	atCmd.Flags().Lookup("fn").NoOptDefVal = "true"
	atCmd.Flags().SortFlags = false

	gcCmd.Flags().BoolVarP(&c, "count", "c", false, "print base content counts (default is proportions)")
	gcCmd.Flags().BoolVarP(&f, "file", "f", false, "calculate statistics per file (default is per record)")
	gcCmd.Flags().BoolVarP(&d, "description", "d", false, "print record descriptions (default is IDs)")
	gcCmd.Flags().BoolVarP(&fn, "fn", "", false, "always print a filename column")
	gcCmd.Flags().Lookup("count").NoOptDefVal = "true"
	gcCmd.Flags().Lookup("file").NoOptDefVal = "true"
	gcCmd.Flags().Lookup("description").NoOptDefVal = "true"
	gcCmd.Flags().Lookup("fn").NoOptDefVal = "true"
	gcCmd.Flags().SortFlags = false

	atgcCmd.Flags().BoolVarP(&c, "count", "c", false, "print base content counts (default is proportions)")
	atgcCmd.Flags().BoolVarP(&f, "file", "f", false, "calculate statistics per file (default is per record)")
	atgcCmd.Flags().BoolVarP(&d, "description", "d", false, "print record descriptions (default is IDs)")
	atgcCmd.Flags().BoolVarP(&fn, "fn", "", false, "always print a filename column")
	atgcCmd.Flags().Lookup("count").NoOptDefVal = "true"
	atgcCmd.Flags().Lookup("file").NoOptDefVal = "true"
	atgcCmd.Flags().Lookup("description").NoOptDefVal = "true"
	atgcCmd.Flags().Lookup("fn").NoOptDefVal = "true"
	atgcCmd.Flags().SortFlags = false

	softCmd.Flags().BoolVarP(&c, "count", "c", false, "print base content counts (default is proportions)")
	softCmd.Flags().BoolVarP(&f, "file", "f", false, "calculate statistics per file (default is per record)")
	softCmd.Flags().BoolVarP(&d, "description", "d", false, "print record descriptions (default is IDs)")
	softCmd.Flags().BoolVarP(&fn, "fn", "", false, "always print a filename column")
	softCmd.Flags().Lookup("count").NoOptDefVal = "true"
	softCmd.Flags().Lookup("file").NoOptDefVal = "true"
	softCmd.Flags().Lookup("description").NoOptDefVal = "true"
	softCmd.Flags().Lookup("fn").NoOptDefVal = "true"
	softCmd.Flags().SortFlags = false

	nCmd.Flags().BoolVarP(&c, "count", "c", false, "print base content counts (default is proportions)")
	nCmd.Flags().BoolVarP(&f, "file", "f", false, "calculate statistics per file (default is per record)")
	nCmd.Flags().BoolVarP(&d, "description", "d", false, "print record descriptions (default is IDs)")
	nCmd.Flags().BoolVarP(&fn, "fn", "", false, "always print a filename column")
	nCmd.Flags().Lookup("count").NoOptDefVal = "true"
	nCmd.Flags().Lookup("file").NoOptDefVal = "true"
	nCmd.Flags().Lookup("description").NoOptDefVal = "true"
	nCmd.Flags().Lookup("fn").NoOptDefVal = "true"
	nCmd.Flags().SortFlags = false

	gapCmd.Flags().BoolVarP(&c, "count", "c", false, "print base content counts (default is proportions)")
	gapCmd.Flags().BoolVarP(&f, "file", "f", false, "calculate statistics per file (default is per record)")
	gapCmd.Flags().BoolVarP(&d, "description", "d", false, "print record descriptions (default is IDs)")
	gapCmd.Flags().BoolVarP(&fn, "fn", "", false, "always print a filename column")
	gapCmd.Flags().Lookup("count").NoOptDefVal = "true"
	gapCmd.Flags().Lookup("file").NoOptDefVal = "true"
	gapCmd.Flags().Lookup("description").NoOptDefVal = "true"
	gapCmd.Flags().Lookup("fn").NoOptDefVal = "true"
	gapCmd.Flags().SortFlags = false

	lenCmd.Flags().BoolVarP(&f, "file", "f", false, "calculate statistics per file (default is per record)")
	lenCmd.Flags().BoolVarP(&fn, "fn", "", false, "always print a filename column")
	lenCmd.Flags().BoolVarP(&kb, "kb", "", false, "print sequence lengths in kilobases")
	lenCmd.Flags().BoolVarP(&mb, "mb", "", false, "print sequence lengths in megabases")
	lenCmd.Flags().BoolVarP(&gb, "gb", "", false, "print sequence lengths in gigabases")
	lenCmd.Flags().Lookup("file").NoOptDefVal = "true"
	lenCmd.Flags().Lookup("kb").NoOptDefVal = "true"
	lenCmd.Flags().Lookup("mb").NoOptDefVal = "true"
	lenCmd.Flags().Lookup("gb").NoOptDefVal = "true"
	lenCmd.Flags().SortFlags = false

	nameCmd.Flags().BoolVarP(&d, "description", "d", false, "print record descriptions (default is IDs)")
	nameCmd.Flags().Lookup("description").NoOptDefVal = "true"
	nameCmd.Flags().SortFlags = false

	assemblyCmd.Flags().IntSliceVarP(&nX, "N", "N", make([]int, 0), "arbitrary NX assembly statistics")
	assemblyCmd.Flags().IntSliceVarP(&lX, "L", "L", make([]int, 0), "arbitrary LX assembly statistics")
	assemblyCmd.Flags().IntSliceVarP(&ngX, "NG", "G", make([]int, 0), "arbitrary NGX assembly statistics (requires -g)")
	assemblyCmd.Flags().IntVarP(&gS, "genomesize", "g", -1, "genome size in bases")
	assemblyCmd.MarkFlagsRequiredTogether("NG", "genomesize")
	assemblyCmd.Flags().BoolVarP(&kb, "kb", "", false, "print N and NG stats in kilobases")
	assemblyCmd.Flags().BoolVarP(&mb, "mb", "", false, "print N and NG stats in megabases")
	assemblyCmd.Flags().BoolVarP(&gb, "gb", "", false, "print N and NG stats in gigabases")
	assemblyCmd.Flags().Lookup("kb").NoOptDefVal = "true"
	assemblyCmd.Flags().Lookup("mb").NoOptDefVal = "true"
	assemblyCmd.Flags().Lookup("gb").NoOptDefVal = "true"
	assemblyCmd.Flags().SortFlags = false

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

func lenMutuallyExclusive(kb, mb, gb bool) (string, error) {
	set := make([]string, 0)
	if kb {
		set = append(set, "kb")
	}
	if mb {
		set = append(set, "mb")
	}
	if gb {
		set = append(set, "gb")
	}
	if len(set) > 1 {
		return "", errors.New("flags --kb, --mb, and --gb are mutually exclusive")
	}
	if len(set) == 1 {
		return set[0], nil
	}
	return "b", nil
}

var contentCmd = &cobra.Command{
	Use:   "content [-c BASES -p BASES] <infile[s]>",
	Short: "Arbitrary base content",
	Long: `Arbitrary base content
	
Default stats when no arguments provided are: GC content, AT content, N content, gap content

Use any combination of the -c and -p flags to override the defaults, e.g.:

fastats content -p GC -c GC -p AT -c AT <infile[s]>

Arguments provided to -p and -c are case-sensitive, so to calculate gc-content for
certain, use, e.g. -p GCgc
`,
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, files []string) (err error) {
		files, err = resolveCommandLine(files)
		if err != nil {
			return err
		}
		c := content{
			inputs:            files,
			perFile:           f,
			writeDescriptions: d,
			writeFileNames:    fn,
			patterns:          make([]pattern, 0),
		}
		switch {
		case len(ps) > 0, len(cs) > 0:
			for _, bases := range ps {
				c.patterns = append(c.patterns, pattern{
					stat:  "prop",
					bases: bases,
				})
			}
			for _, bases := range cs {
				c.patterns = append(c.patterns, pattern{
					stat:  "count",
					bases: bases,
				})
			}
		default:
			c.patterns = []pattern{
				{
					stat:  "prop",
					bases: "GCgc",
				},
				{
					stat:  "prop",
					bases: "ATat",
				},
				{
					stat:  "prop",
					bases: "Nn",
				},
				{
					stat:  "prop",
					bases: "-",
				},
			}
		}
		err = c.writeHeader(out)
		if err != nil {
			return err
		}
		err = c.writeBody(out)
		return err
	},
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
		stat := "prop"
		if c {
			stat = "count"
		}
		p := content{
			inputs:            files,
			perFile:           f,
			writeDescriptions: d,
			writeFileNames:    fn,
			patterns: []pattern{
				{
					stat:  stat,
					bases: "ATat",
				},
			},
		}
		err = p.writeHeader(out)
		if err != nil {
			return err
		}
		err = p.writeBody(out)
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
		stat := "prop"
		if c {
			stat = "count"
		}
		p := content{
			inputs:            files,
			perFile:           f,
			writeDescriptions: d,
			writeFileNames:    fn,
			patterns: []pattern{
				{
					stat:  stat,
					bases: "GCgc",
				},
			},
		}
		err = p.writeHeader(out)
		if err != nil {
			return err
		}
		err = p.writeBody(out)
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
		stat := "prop"
		if c {
			stat = "count"
		}
		p := content{
			inputs:            files,
			perFile:           f,
			writeDescriptions: d,
			writeFileNames:    fn,
			patterns: []pattern{
				{
					stat:  stat,
					bases: "ATGCatgc",
				},
			},
		}
		err = p.writeHeader(out)
		if err != nil {
			return err
		}
		err = p.writeBody(out)
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
		stat := "prop"
		if c {
			stat = "count"
		}
		p := content{
			inputs:            files,
			perFile:           f,
			writeDescriptions: d,
			writeFileNames:    fn,
			patterns: []pattern{
				{
					stat:  stat,
					bases: "Nn",
				},
			},
		}
		err = p.writeHeader(out)
		if err != nil {
			return err
		}
		err = p.writeBody(out)
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
		stat := "prop"
		if c {
			stat = "count"
		}
		p := content{
			inputs:            files,
			perFile:           f,
			writeDescriptions: d,
			writeFileNames:    fn,
			patterns: []pattern{
				{
					stat:  stat,
					bases: "-",
				},
			},
		}
		err = p.writeHeader(out)
		if err != nil {
			return err
		}
		err = p.writeBody(out)
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
		stat := "prop"
		if c {
			stat = "count"
		}
		p := content{
			inputs:            files,
			perFile:           f,
			writeDescriptions: d,
			writeFileNames:    fn,
			patterns: []pattern{
				{
					stat:  stat,
					bases: "atgcn",
				},
			},
		}
		err = p.writeHeader(out)
		if err != nil {
			return err
		}
		err = p.writeBody(out)
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
		lenFormat, err := lenMutuallyExclusive(kb, mb, gb)
		if err != nil {
			return err
		}
		l := length{
			inputs:            files,
			perFile:           f,
			writeDescriptions: d,
			writeFileNames:    fn,
			lenFormat:         lenFormat,
		}
		err = l.writeHeader(out)
		if err != nil {
			return err
		}
		err = l.writeBody(out)
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
		err = n.writeHeader(out)
		if err != nil {
			return err
		}
		err = n.writeBody(out)
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
		err = n.writeHeader(out)
		if err != nil {
			return err
		}
		err = n.writeBody(out)
		return err
	},
}

var assemblyCmd = &cobra.Command{
	Use:   "assembly [-N50 -L50 -G50 [-g int]] <infile[s]>",
	Short: "Assembly statistics",
	Long: `Assembly statistics

Default stats when no arguments provided are: N50, N90, L50, L90

Use any combination of the --N, --L, --NG (-g) flags to override the defaults, e.g.:

fastats assembly --N50 --N90 --NG50 --NG90 -g 3000000000 <infile[s]>
`,
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, files []string) (err error) {
		files, err = resolveCommandLine(files)
		if err != nil {
			return err
		}
		lenFormat, err := lenMutuallyExclusive(kb, mb, gb)
		if err != nil {
			return err
		}
		a := assembly{
			inputs:     files,
			stats:      make([]assemblyStatistic, 0),
			genomeSize: gS,
			lenFormat:  lenFormat,
		}
		switch {
		case len(nX) > 0, len(lX) > 0, len(ngX) > 0:
			for _, v := range nX {
				a.stats = append(a.stats, assemblyStatistic{
					sType:  "N",
					sValue: v,
				})
			}
			for _, v := range lX {
				a.stats = append(a.stats, assemblyStatistic{
					sType:  "L",
					sValue: v,
				})
			}
			for _, v := range ngX {
				a.stats = append(a.stats, assemblyStatistic{
					sType:  "NG",
					sValue: v,
				})
			}
		default:
			a.stats = []assemblyStatistic{
				{
					sType:  "N",
					sValue: 50,
				},
				{
					sType:  "N",
					sValue: 90,
				},
				{
					sType:  "L",
					sValue: 50,
				},
				{
					sType:  "L",
					sValue: 90,
				},
			}
		}

		err = a.writeHeader(out)
		if err != nil {
			return err
		}
		err = a.writeBody(out)
		return err
	},
}
