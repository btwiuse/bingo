# Go Binary Builder

This Go program is a binary builder for the k0s project. It supports cross-compilation for multiple operating system and architecture combinations.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine.

### Prerequisites

* Go (version 1.16 or later)
* UPX (optional, for binary compression)
* strip utility (optional, for removing symbol table and relocation information)

### Installing

Clone the repository to your local machine:

```bash
go install github.com/btwiuse/bingo/cmd/bingo@latest
```

## Usage

The binary builder can be run using the 'go run' command, followed by any flags you wish to use. Here's an example:

```bash
bingo -d ./bin -tags "tag1 tag2" -ldflags "-X main.version=1.0.0" -strip -upx -pkg ./cmd/k0s
```

### Flags

The binary builder supports the following flags:

* `-d`: Specify the output directory for the built binary.
* `-pkg`: Specify the path of the package to build.
* `-tags`: Specify build tags.
* `-ldflags`: Specify linker flags.
* `-strip`: Strip the binary (removes symbol table and relocation information).
* `-upx`: Compress the binary using UPX.
* `-dry`: Dry run. Show what would be done without actually doing anything.
