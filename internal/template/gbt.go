package template

import (
"context"
"encoding/json"
"log"
"time"

"github.com/caldefenwycke/hyperhash-core/internal/rpc"
)

type GBTResult struct {
Height           int    `json:"height"`
PreviousBlockHash string `json:"previousblockhash"`
Bits             string `json:"bits"`
CurTime          int64  `json:"curtime"`
// add more as needed
}

type Cache struct {
Last *GBTResult
}

func StartGBTPoller(ctx context.Context, c *rpc.Client, interval time.Duration, cache *Cache) {
t := time.NewTicker(interval)
defer t.Stop()

call := func() {
var res GBTResult
// During IBD this can return error code -10, which we just log.
err := c.Call("getblocktemplate", []interface{}{map[string]interface{}{"rules": []string{"segwit"}}}, &res)
if err != nil {
log.Printf("[gbt] fetch error: %v", err)
return
}
// copy to cache
cache.Last = &res
// lightweight log line for observability
log.Printf("[gbt] height=%d prev=%s bits=%s time=%d", res.Height, short(res.PreviousBlockHash), res.Bits, res.CurTime)
}
// run immediately then on ticker
call()
for {
select {
case <-ctx.Done():
return
case <-t.C:
call()
}
}
}

func short(h string) string {
if len(h) <= 8 {
return h
}
return h[:8]
}
