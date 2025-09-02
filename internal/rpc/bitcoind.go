package rpc

import (
"bytes"
"encoding/base64"
"encoding/json"
"fmt"
"net/http"
"time"
)

type Client struct {
url     string
authHdr string
cli     *http.Client
}

func New(host string, port int, user, pass string) *Client {
u := fmt.Sprintf("http://%s:%d", host, port)
auth := base64.StdEncoding.EncodeToString([]byte(user + ":" + pass))
return &Client{
url:     u,
authHdr: "Basic " + auth,
cli: &http.Client{
Timeout: 8 * time.Second,
},
}
}

type rpcReq struct {
JSONRPC string      `json:"jsonrpc"`
Method  string      `json:"method"`
Params  interface{} `json:"params"`
ID      string      `json:"id"`
}

type rpcResp struct {
Result json.RawMessage `json:"result"`
Error  *struct {
Code    int    `json:"code"`
Message string `json:"message"`
} `json:"error"`
ID string `json:"id"`
}

func (c *Client) Call(method string, params interface{}, out interface{}) error {
body, _ := json.Marshal(rpcReq{JSONRPC: "1.0", Method: method, Params: params, ID: "hh"})
req, _ := http.NewRequest("POST", c.url, bytes.NewReader(body))
req.Header.Set("Content-Type", "text/plain")
req.Header.Set("Authorization", c.authHdr)

resp, err := c.cli.Do(req)
if err != nil {
return err
}
defer resp.Body.Close()

var r rpcResp
if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
return err
}
if r.Error != nil {
return fmt.Errorf("rpc error %d: %s", r.Error.Code, r.Error.Message)
}
if out != nil && len(r.Result) > 0 {
return json.Unmarshal(r.Result, out)
}
return nil
}
