#!/usr/bin/env bash
set -euo pipefail
DATADIR="/var/lib/bitcoind"
OUTDIR="$DATADIR/wallet-backups"
PASSPH="/root/.hyperhash/backup-passphrase"
CLI="/usr/local/bin/bitcoin-cli -datadir=$DATADIR"
mkdir -p "$OUTDIR"
# Use loaded wallets; fallback to poolhot
WALLETS=$($CLI listwallets | jq -r '.[]')
[ -n "$WALLETS" ] || WALLETS="poolhot"
TS=$(date -u +%Y%m%dT%H%M%SZ)
for W in $WALLETS; do
  TMP=$(mktemp -p "$OUTDIR" "${W}-${TS}.XXXX.tmp")
  sudo -u bitcoin $CLI -rpcwallet="$W" backupwallet "$TMP"
  OUT="$OUTDIR/${W}-${TS}.dat.gpg"
  /usr/bin/gpg --batch --yes --symmetric --cipher-algo AES256 --passphrase-file "$PASSPH" -o "$OUT" "$TMP"
  /usr/bin/sha256sum "$OUT" > "$OUT.sha256"
  /usr/bin/shred -u "$TMP"
  chmod 600 "$OUT" "$OUT.sha256"
  chown root:root "$OUT" "$OUT.sha256"
  echo "backup: $OUT"
done
# Retention: 30 days
find "$OUTDIR" -type f -name "*.dat.gpg"  -mtime +30 -delete || true
find "$OUTDIR" -type f -name "*.sha256"   -mtime +30 -delete || true