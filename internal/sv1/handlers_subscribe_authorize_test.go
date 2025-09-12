package sv1

import "testing"

func TestHandleMiningSubscribe_ReturnsShape(t *testing.T) {
res, err := HandleMiningSubscribe(nil)
if err != nil {
t.Fatalf("unexpected error: %v", err)
}
arr, ok := res.([]any)
if !ok {
t.Fatalf("expected []any, got %T", res)
}
if len(arr) != 3 {
t.Fatalf("expected len 3, got %d", len(arr))
}
// subscriptions
if _, ok := arr[0].([][]string); !ok {
t.Fatalf("subscriptions wrong type: %T", arr[0])
}
// extranonce1
if _, ok := arr[1].(string); !ok {
t.Fatalf("extranonce1 wrong type: %T", arr[1])
}
// extranonce2_size
if _, ok := arr[2].(int); !ok {
t.Fatalf("extranonce2_size wrong type: %T", arr[2])
}
}

func TestHandleMiningAuthorize(t *testing.T) {
// missing username -> error
if _, err := HandleMiningAuthorize([]any{}); err == nil {
t.Fatalf("expected error on missing username")
}
// happy path
out, err := HandleMiningAuthorize([]any{"acct.worker", "x"})
if err != nil {
t.Fatalf("unexpected error: %v", err)
}
ok, isBool := out.(bool)
if !isBool || !ok {
t.Fatalf("expected true bool, got %T %#v", out, out)
}
}
