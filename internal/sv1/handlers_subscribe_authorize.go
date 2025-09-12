package sv1

import "errors"

// HandleMiningSubscribe implements "mining.subscribe".
// Params: [ userAgent?, sessionID? ]
// Returns: [ subscriptions, extranonce1, extranonce2_size ]
func HandleMiningSubscribe(params []any) (any, error) {
var userAgent, sessionID string
if len(params) > 0 {
if s, ok := params[0].(string); ok { userAgent = s }
}
if len(params) > 1 {
if s, ok := params[1].(string); ok { sessionID = s }
}
_ = userAgent
_ = sessionID

extranonce1 := "00000001" // TODO: per-connection bytes (hex)
extranonce2Size := 4      // TODO: confirm against coinbase builder

subscriptions := [][]string{
{"mining.set_difficulty", "client.subscription"},
{"mining.notify", "client.subscription"},
}
return []any{subscriptions, extranonce1, extranonce2Size}, nil
}

// HandleMiningAuthorize implements "mining.authorize".
// Params: [ username, password? ]
// Returns: true on success
func HandleMiningAuthorize(params []any) (any, error) {
if len(params) < 1 {
return nil, errors.New("missing username")
}
username, ok := params[0].(string)
if !ok || username == "" {
return nil, errors.New("invalid username")
}
// TODO: validate <account.worker>, check bans/ACLs
_ = username
return true, nil
}