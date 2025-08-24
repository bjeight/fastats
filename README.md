# fastats

Very simple statistics from fasta files


### Installation

First, [install go](https://go.dev/dl/),

then:

```
go install github.com/bjeight/fastats@latest
```

or

```
git clone https://github.com/bjeight/fastats.git
cd fastats
go build -o fastats
```

### Example usage

```
❯ fastats at PlasmoDB-64_Pfalciparum3D7_Genome.fasta
record	ATat_prop
Pf3D7_01_v3	0.794985
Pf3D7_02_v3	0.802509
Pf3D7_03_v3	0.799358
Pf3D7_04_v3	0.794851
Pf3D7_05_v3	0.806723
Pf3D7_06_v3	0.802128
Pf3D7_07_v3	0.801507
Pf3D7_08_v3	0.804305
Pf3D7_09_v3	0.809843
Pf3D7_10_v3	0.803563
Pf3D7_11_v3	0.810128
Pf3D7_12_v3	0.806968
Pf3D7_13_v3	0.810485
Pf3D7_14_v3	0.815636
Pf3D7_API_v3	0.857839
Pf3D7_MIT_v3	0.684096
```

```
❯ fastats soft -f *softmasked.fasta
file	atgcn_prop
Pad.softmasked.fasta	0.612057
Pbi.softmasked.fasta	0.545409
Pbl.softmasked.fasta	0.576649
Pfa.softmasked.fasta	0.519676
Pga.softmasked.fasta	0.610234
Ppr.softmasked.fasta	0.539591
Pre.softmasked.fasta	0.535700
```

```
❯ bgzip -d -c GRCh38.primary_assembly.genome.bgz | fastats gc | grep "Y"
chrY	0.184749
```

```
❯ fastats gc GRCh38.primary_assembly.genome.fa.bgz | grep "X"
chrX	0.392343
```

```
❯ fastats assembly GRCh38.primary_assembly.genome.fa
file  N50 N90 L50 L90
GRCh38.primary_assembly.genome.fa	145138636	80373285	9	19
```

### Help

```
❯ fastats -h
Very simple statistics from fasta files

Usage:
  fastats [command]

Available Commands:
  ambig       Ambiguous content
  assembly    Assembly statistics
  at          AT content
  atgc        ATGC content
  content     Arbitrary base content
  gaps        Gap content
  gc          GC content
  help        Help about any command
  len         Sequence length
  n           N content
  names       Record names
  num         Number of records
  soft        Softmasked content

Flags:
  -h, --help      help for fastats
  -v, --version   version for fastats

Use "fastats [command] --help" for more information about a command.
```
