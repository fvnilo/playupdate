# playupdate
> Update music players in batch

## Getting Started

### Prerequisites
- [Golang](https://golang.org/)

### How to build

#### All Environments
```
make
```

#### MacOSX
```
make build_macosx
```

#### Linux
```
make build_linux
```

#### Windows
```
make build_windows
```

## Usage
```
NAME:
   playupdate - Update music players in batch

USAGE:
   main [global options] command [command options] [arguments...]

VERSION:
   0.1.0

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --input FILE, -i FILE  Load input from FILE (mandatory)
   --api-url URL          Api base URL (optional, defaults to http://localhost:8080)
   --help, -h             show help
   --version, -v          print the version
```

Most of the time you will run:
```
playupdate --input file-path-here.csv
```

If you want to use another base url for the update api, you can specify `--api-url url-here.com`.


### Local usage
In local, you can run the `dummyserver` command first in order to have a simple http server that will respond with a 200:
```
go run main/dummyserver/main.go
```

When it is running, you can run the cli with `go run cmd/playupdate/main.go -i examples/input.csv`.
