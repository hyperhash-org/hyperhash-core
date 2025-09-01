#!/usr/bin/env bash
set -euo pipefail
echo "[build-core] building hh-core..."
go build -buildvcs=false -trimpath -o /opt/hyperhash/bin/hh-core ./cmd/core
echo "[build-core] done."
