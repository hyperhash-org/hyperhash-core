package main

import (
"flag"
"fmt"
"log"
"os"
"time"
)

func main() {
cfg := flag.String("config", "/opt/hyperhash/configs/pool.toml", "path to config")
flag.Parse()

log.SetFlags(0)
log.Printf("[core] starting with config %s", *cfg)

// prove binary picks up args and keeps running
for {
fmt.Printf("[core] alive %s\n", time.Now().Format(time.RFC3339))
time.Sleep(10 * time.Second)
}
}
