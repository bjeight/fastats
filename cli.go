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
		Version:           "0.10.1",
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
	v bool

	cs  []string
	ps  []string
	vcs []string
	vps []string

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

func getFileBoolVarPArgs() (*bool, string, string, bool, string) {
	return &f, "file", "f", false, "calculate statistics per file (default is per record)"
}

func getDescriptionBoolVarPArgs() (*bool, string, string, bool, string) {
	return &d, "description", "d", false, "print record descriptions (default is IDs)"
}

func getFileNameBoolVarArgs() (*bool, string, bool, string) {
	return &fn, "fn", false, "always print a filename column"
}

func getCountBoolVarPArgs() (*bool, string, string, bool, string) {
	return &c, "count", "c", false, "print base content counts (default is proportions)"
}

func getInverseBoolVarPArgs() (*bool, string, string, bool, string) {
	return &v, "inverse", "v", false, "count bases that are NOT the given content"
}

func init() {

	rootCmd.AddCommand(contentCmd)
	rootCmd.AddCommand(atCmd)
	rootCmd.AddCommand(gcCmd)
	rootCmd.AddCommand(atgcCmd)
	rootCmd.AddCommand(nCmd)
	rootCmd.AddCommand(gapCmd)
	rootCmd.AddCommand(ambigCmd)
	rootCmd.AddCommand(softCmd)
	rootCmd.AddCommand(lenCmd)
	rootCmd.AddCommand(numCmd)
	rootCmd.AddCommand(nameCmd)
	rootCmd.AddCommand(assemblyCmd)

	contentCmd.Flags().StringSliceVarP(&cs, "count", "c", make([]string, 0), "comma-separated list of arbitrary base contents to count (case-sensitive)")
	contentCmd.Flags().StringSliceVarP(&ps, "prop", "p", make([]string, 0), "comma-separated list of arbitrary base contents to get proportions of (case-sensitive)")
	contentCmd.Flags().StringSliceVar(&vcs, "vc", make([]string, 0), "comma-separated list of inverse arbitrary base contents to count (case-sensitive)")
	contentCmd.Flags().StringSliceVar(&vps, "vp", make([]string, 0), "comma-separated list of inverse arbitrary base content get proportions of (case-sensitive)")
	contentCmd.Flags().BoolVarP(getFileBoolVarPArgs())
	contentCmd.Flags().BoolVarP(getDescriptionBoolVarPArgs())
	contentCmd.Flags().BoolVar(getFileNameBoolVarArgs())
	contentCmd.Flags().Lookup("file").NoOptDefVal = "true"
	contentCmd.Flags().Lookup("description").NoOptDefVal = "true"
	contentCmd.Flags().Lookup("fn").NoOptDefVal = "true"
	contentCmd.Flags().SortFlags = false

	atCmd.Flags().BoolVarP(getCountBoolVarPArgs())
	atCmd.Flags().BoolVarP(getFileBoolVarPArgs())
	atCmd.Flags().BoolVarP(getInverseBoolVarPArgs())
	atCmd.Flags().BoolVarP(getDescriptionBoolVarPArgs())
	atCmd.Flags().BoolVar(getFileNameBoolVarArgs())
	atCmd.Flags().Lookup("count").NoOptDefVal = "true"
	atCmd.Flags().Lookup("file").NoOptDefVal = "true"
	atCmd.Flags().Lookup("inverse").NoOptDefVal = "true"
	atCmd.Flags().Lookup("description").NoOptDefVal = "true"
	atCmd.Flags().Lookup("fn").NoOptDefVal = "true"
	atCmd.Flags().SortFlags = false

	gcCmd.Flags().BoolVarP(getCountBoolVarPArgs())
	gcCmd.Flags().BoolVarP(getFileBoolVarPArgs())
	gcCmd.Flags().BoolVarP(getInverseBoolVarPArgs())
	gcCmd.Flags().BoolVarP(getDescriptionBoolVarPArgs())
	gcCmd.Flags().BoolVar(getFileNameBoolVarArgs())
	gcCmd.Flags().Lookup("count").NoOptDefVal = "true"
	gcCmd.Flags().Lookup("file").NoOptDefVal = "true"
	gcCmd.Flags().Lookup("inverse").NoOptDefVal = "true"
	gcCmd.Flags().Lookup("description").NoOptDefVal = "true"
	gcCmd.Flags().Lookup("fn").NoOptDefVal = "true"
	gcCmd.Flags().SortFlags = false

	atgcCmd.Flags().BoolVarP(getCountBoolVarPArgs())
	atgcCmd.Flags().BoolVarP(getFileBoolVarPArgs())
	atgcCmd.Flags().BoolVarP(getInverseBoolVarPArgs())
	atgcCmd.Flags().BoolVarP(getDescriptionBoolVarPArgs())
	atgcCmd.Flags().BoolVar(getFileNameBoolVarArgs())
	atgcCmd.Flags().Lookup("count").NoOptDefVal = "true"
	atgcCmd.Flags().Lookup("file").NoOptDefVal = "true"
	atgcCmd.Flags().Lookup("inverse").NoOptDefVal = "true"
	atgcCmd.Flags().Lookup("description").NoOptDefVal = "true"
	atgcCmd.Flags().Lookup("fn").NoOptDefVal = "true"
	atgcCmd.Flags().SortFlags = false

	softCmd.Flags().BoolVarP(getCountBoolVarPArgs())
	softCmd.Flags().BoolVarP(getFileBoolVarPArgs())
	softCmd.Flags().BoolVarP(getInverseBoolVarPArgs())
	softCmd.Flags().BoolVarP(getDescriptionBoolVarPArgs())
	softCmd.Flags().BoolVar(getFileNameBoolVarArgs())
	softCmd.Flags().Lookup("count").NoOptDefVal = "true"
	softCmd.Flags().Lookup("file").NoOptDefVal = "true"
	softCmd.Flags().Lookup("inverse").NoOptDefVal = "true"
	softCmd.Flags().Lookup("description").NoOptDefVal = "true"
	softCmd.Flags().Lookup("fn").NoOptDefVal = "true"
	softCmd.Flags().SortFlags = false

	nCmd.Flags().BoolVarP(getCountBoolVarPArgs())
	nCmd.Flags().BoolVarP(getFileBoolVarPArgs())
	nCmd.Flags().BoolVarP(getInverseBoolVarPArgs())
	nCmd.Flags().BoolVarP(getDescriptionBoolVarPArgs())
	nCmd.Flags().BoolVar(getFileNameBoolVarArgs())
	nCmd.Flags().Lookup("count").NoOptDefVal = "true"
	nCmd.Flags().Lookup("file").NoOptDefVal = "true"
	nCmd.Flags().Lookup("inverse").NoOptDefVal = "true"
	nCmd.Flags().Lookup("description").NoOptDefVal = "true"
	nCmd.Flags().Lookup("fn").NoOptDefVal = "true"
	nCmd.Flags().SortFlags = false

	gapCmd.Flags().BoolVarP(getCountBoolVarPArgs())
	gapCmd.Flags().BoolVarP(getFileBoolVarPArgs())
	gapCmd.Flags().BoolVarP(getInverseBoolVarPArgs())
	gapCmd.Flags().BoolVarP(getDescriptionBoolVarPArgs())
	gapCmd.Flags().BoolVar(getFileNameBoolVarArgs())
	gapCmd.Flags().Lookup("count").NoOptDefVal = "true"
	gapCmd.Flags().Lookup("file").NoOptDefVal = "true"
	gapCmd.Flags().Lookup("inverse").NoOptDefVal = "true"
	gapCmd.Flags().Lookup("description").NoOptDefVal = "true"
	gapCmd.Flags().Lookup("fn").NoOptDefVal = "true"
	gapCmd.Flags().SortFlags = false

	ambigCmd.Flags().BoolVarP(getCountBoolVarPArgs())
	ambigCmd.Flags().BoolVarP(getFileBoolVarPArgs())
	ambigCmd.Flags().BoolVarP(getInverseBoolVarPArgs())
	ambigCmd.Flags().BoolVarP(getDescriptionBoolVarPArgs())
	ambigCmd.Flags().BoolVar(getFileNameBoolVarArgs())
	ambigCmd.Flags().Lookup("count").NoOptDefVal = "true"
	ambigCmd.Flags().Lookup("file").NoOptDefVal = "true"
	ambigCmd.Flags().Lookup("inverse").NoOptDefVal = "true"
	ambigCmd.Flags().Lookup("description").NoOptDefVal = "true"
	ambigCmd.Flags().Lookup("fn").NoOptDefVal = "true"
	ambigCmd.Flags().SortFlags = false

	lenCmd.Flags().BoolVarP(getFileBoolVarPArgs())
	lenCmd.Flags().BoolVar(getFileNameBoolVarArgs())
	lenCmd.Flags().BoolVar(&kb, "kb", false, "print sequence lengths in kilobases")
	lenCmd.Flags().BoolVar(&mb, "mb", false, "print sequence lengths in megabases")
	lenCmd.Flags().BoolVar(&gb, "gb", false, "print sequence lengths in gigabases")
	lenCmd.Flags().Lookup("file").NoOptDefVal = "true"
	lenCmd.Flags().Lookup("fn").NoOptDefVal = "true"
	lenCmd.Flags().Lookup("kb").NoOptDefVal = "true"
	lenCmd.Flags().Lookup("mb").NoOptDefVal = "true"
	lenCmd.Flags().Lookup("gb").NoOptDefVal = "true"
	lenCmd.Flags().SortFlags = false

	nameCmd.Flags().BoolVarP(getDescriptionBoolVarPArgs())
	nameCmd.Flags().Lookup("description").NoOptDefVal = "true"
	nameCmd.Flags().SortFlags = false

	assemblyCmd.Flags().IntSliceVarP(&nX, "N", "N", make([]int, 0), "arbitrary NX assembly statistics")
	assemblyCmd.Flags().IntSliceVarP(&lX, "L", "L", make([]int, 0), "arbitrary LX assembly statistics")
	assemblyCmd.Flags().IntSliceVarP(&ngX, "NG", "G", make([]int, 0), "arbitrary NGX assembly statistics (requires -g)")
	assemblyCmd.Flags().IntVarP(&gS, "genomesize", "g", -1, "genome size in bases")
	assemblyCmd.MarkFlagsRequiredTogether("NG", "genomesize")
	assemblyCmd.Flags().BoolVar(&kb, "kb", false, "print N and NG stats in kilobases")
	assemblyCmd.Flags().BoolVar(&mb, "mb", false, "print N and NG stats in megabases")
	assemblyCmd.Flags().BoolVar(&gb, "gb", false, "print N and NG stats in gigabases")
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

Use any combination of the -c, -p, --vc and --vp flags to override the defaults, e.g.:

fastats content -p GC,AT -c GC,AT <infile[s]>

Arguments provided to -p, -c, --vp, --vc are case-sensitive, so to calculate gc-content for
certain, use, e.g.: -p GCgc
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
		case len(ps) > 0, len(cs) > 0, len(vps) > 0, len(vcs) > 0:
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
			for _, bases := range vps {
				c.patterns = append(c.patterns, pattern{
					stat:    "prop",
					bases:   bases,
					inverse: true,
				})
			}
			for _, bases := range vcs {
				c.patterns = append(c.patterns, pattern{
					stat:    "count",
					bases:   bases,
					inverse: true,
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
					stat:         "prop",
					bases:        "-",
					headerPrefix: "gap",
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
	Long:                  `AT[+at] nucleotide content`,
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
		ctn := content{
			inputs:            files,
			perFile:           f,
			writeDescriptions: d,
			writeFileNames:    fn,
			patterns: []pattern{
				{
					stat:    stat,
					bases:   "ATat",
					inverse: v,
				},
			},
		}
		err = ctn.writeHeader(out)
		if err != nil {
			return err
		}
		err = ctn.writeBody(out)
		return err
	},
}

var gcCmd = &cobra.Command{
	Use:                   "gc <infile[s]>",
	Short:                 "GC content",
	Long:                  `GC[+gc] nucleotide content`,
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
		ctn := content{
			inputs:            files,
			perFile:           f,
			writeDescriptions: d,
			writeFileNames:    fn,
			patterns: []pattern{
				{
					stat:    stat,
					bases:   "GCgc",
					inverse: v,
				},
			},
		}
		err = ctn.writeHeader(out)
		if err != nil {
			return err
		}
		err = ctn.writeBody(out)
		return err
	},
}

var atgcCmd = &cobra.Command{
	Use:                   "atgc <infile[s]>",
	Short:                 "ATGC content",
	Long:                  `ATGC[+atgc] nucleotide content`,
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
		ctn := content{
			inputs:            files,
			perFile:           f,
			writeDescriptions: d,
			writeFileNames:    fn,
			patterns: []pattern{
				{
					stat:    stat,
					bases:   "ATGCatgc",
					inverse: v,
				},
			},
		}
		err = ctn.writeHeader(out)
		if err != nil {
			return err
		}
		err = ctn.writeBody(out)
		return err
	},
}

var nCmd = &cobra.Command{
	Use:                   "n <infile[s]>",
	Short:                 "N content",
	Long:                  `N[+n] content`,
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
		ctn := content{
			inputs:            files,
			perFile:           f,
			writeDescriptions: d,
			writeFileNames:    fn,
			patterns: []pattern{
				{
					stat:    stat,
					bases:   "Nn",
					inverse: v,
				},
			},
		}
		err = ctn.writeHeader(out)
		if err != nil {
			return err
		}
		err = ctn.writeBody(out)
		return err
	},
}

var gapCmd = &cobra.Command{
	Use:                   "gaps <infile[s]>",
	Short:                 "Gap content",
	Long:                  `Gap ("-") content`,
	Aliases:               []string{"gaps"},
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
		ctn := content{
			inputs:            files,
			perFile:           f,
			writeDescriptions: d,
			writeFileNames:    fn,
			patterns: []pattern{
				{
					stat:         stat,
					bases:        "-",
					headerPrefix: "gap",
					inverse:      v,
				},
			},
		}
		err = ctn.writeHeader(out)
		if err != nil {
			return err
		}
		err = ctn.writeBody(out)
		return err
	},
}

var ambigCmd = &cobra.Command{
	Use:                   "ambig <infile[s]>",
	Short:                 "Ambiguous content",
	Long:                  `Ambiguous ("RYMKSWHBVDNrymkswhbvdn") nucleotide content`,
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
		ctn := content{
			inputs:            files,
			perFile:           f,
			writeDescriptions: d,
			writeFileNames:    fn,
			patterns: []pattern{
				{
					stat:         stat,
					bases:        "RYMKSWHBVDNrymkswhbvdn",
					headerPrefix: "ambig",
					inverse:      v,
				},
			},
		}
		err = ctn.writeHeader(out)
		if err != nil {
			return err
		}
		err = ctn.writeBody(out)
		return err
	},
}

var softCmd = &cobra.Command{
	Use:                   "soft <infile[s]>",
	Short:                 "Softmasked content",
	Long:                  `Softmasked ("atgcrymkswhbvdn") nucleotide content`,
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
		ctn := content{
			inputs:            files,
			perFile:           f,
			writeDescriptions: d,
			writeFileNames:    fn,
			patterns: []pattern{
				{
					stat:         stat,
					bases:        "atgcrymkswhbvdn",
					headerPrefix: "soft",
					inverse:      v,
				},
			},
		}
		err = ctn.writeHeader(out)
		if err != nil {
			return err
		}
		err = ctn.writeBody(out)
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
