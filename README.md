# q-go

A terminal-based todo list application with subject organization.

## Installation

### Build from source

```bash
make
```

### Install system-wide

```bash
make install
```

This installs the binary to `/usr/local/bin/q-go`.

## Running

After building:
```bash
./build/q-go
```

After installing:
```bash
q-go
```

## Data Storage

The application saves your todo lists and subjects to `~/.q-go/data.yaml`. This file is automatically created when you first use the application.

## Building

### Build the application
```bash
make
```

### Clean build artifacts
```bash
make clean
```

## Available Makefile Commands

- `make` or `make all` - Build the application to `./build/q-go`
- `make clean` - Remove the built binary
- `make install` - Copy the binary to `/usr/local/bin/q-go` (requires build first)