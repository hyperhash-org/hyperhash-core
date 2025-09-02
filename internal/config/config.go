package config

import (
"os"

"github.com/pelletier/go-toml/v2"
)

type Bitcoind struct {
RPCHost string `toml:"rpc_host"`
RPCPort int    `toml:"rpc_port"`
RPCUser string `toml:"rpc_user"`
RPCPass string `toml:"rpc_pass"`
Network string `toml:"network"`
}

type Jobs struct {
RefreshMS int `toml:"refresh_ms"`
}

type Config struct {
Bitcoind Bitcoind `toml:"bitcoind"`
Jobs     Jobs     `toml:"jobs"`
}

func Load(path string) (*Config, error) {
b, err := os.ReadFile(path)
if err != nil {
return nil, err
}
var c Config
if err := toml.Unmarshal(b, &c); err != nil {
return nil, err
}
// sensible defaults
if c.Jobs.RefreshMS == 0 {
c.Jobs.RefreshMS = 1000
}
return &c, nil
}
