package main

import (
"context"
"flag"
"log"
"time"

"github.com/caldefenwycke/hyperhash-core/internal/config"
"github.com/caldefenwycke/hyperhash-core/internal/rpc"
"github.com/caldefenwycke/hyperhash-core/internal/template"
)

func main() {
cfgPath := flag.String("config", "/opt/hyperhash/configs/pool.toml", "path to config")
flag.Parse()
log.SetFlags(0)

// load config
c, err := config.Load(*cfgPath)
if err != nil {
log.Fatalf("[core] config load error: %v", err)
}
log.Printf("[core] starting with config %s (network=%s)", *cfgPath, c.Bitcoind.Network)

// rpc client
r := rpc.New(c.Bitcoind.RPCHost, c.Bitcoind.RPCPort, c.Bitcoind.RPCUser, c.Bitcoind.RPCPass)

// start GBT poller
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
cache := &template.Cache{}
go template.StartGBTPoller(ctx, r, time.Duration(c.Jobs.RefreshMS)*time.Millisecond, cache)

// simple heartbeat
t := time.NewTicker(10 * time.Second)
defer t.Stop()
for range t.C {
log.Printf("[core] alive")
}
}
