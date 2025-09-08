[![CI](https://github.com/hyperhash-org/hyperhash-core/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/hyperhash-org/hyperhash-core/actions/workflows/ci.yml)

# Hyper Hash — Core

Core services and libraries for the Hyper Hash mining pool.  
Provides job management, header generation, and utilities shared across pool, edge, and UI.

## Features
- Header generation (SV1, SV2, Hyper lanes)
- Version rolling & deduplication
- Midstate reuse and optimizations
- Common config & logging

## Usage
Imported by:
- `hyperhash-pool`
- `hyperhash-edge`
- `hyperhash-ui`

## Quickstart

### Build
```bash
make build
# or
go build -o ./bin/hh-core ./cmd/hh-core

### Run
```bash
./bin/hh-core
