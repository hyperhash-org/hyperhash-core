#!/usr/bin/env bash
set -euo pipefail

UNIT_SRC="$(dirname "$0")/hh-core.service"
UNIT_DST="/etc/systemd/system/hh-core.service"

# ensure dirs exist
sudo mkdir -p /opt/hyperhash/bin /opt/hyperhash/configs /opt/hyperhash/logs
sudo chown -R hyperhash:hyperhash /opt/hyperhash

# install unit
sudo cp "$UNIT_SRC" "$UNIT_DST"
sudo systemctl daemon-reload
sudo systemctl enable hh-core

echo "Installed hh-core.service. Next steps:"
echo "  1) put binary at /opt/hyperhash/bin/hh-core (chmod +x)"
echo "  2) create /opt/hyperhash/configs/pool.toml (copy example)"
echo "  3) sudo systemctl start hh-core && journalctl -u hh-core -f"
