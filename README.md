# fastats

Very simple statistics from fasta files


### Installation

```
git clone https://github.com/bjeight/fastats.git
cd fastats
go build -o fastats
```

### Usage

```
Very simple statistics from fasta files

Usage:
  fastats [command]

Available Commands:
  at          AT content
  atgc        ATGC content
  gaps        Gap content
  gc          GC content
  help        Help about any command
  len         Sequence length
  n           N content
  pattern     PATTERN content
  soft        Softmasked content

Flags:
  -c, --count          print counts, not proportions
  -f, --file[=false]   calculate statists per file (default is per record)
  -h, --help           help for fastats
  -v, --version        version for fastats

Use "fastats [command] --help" for more information about a command.
```
