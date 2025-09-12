# Pool Fee Outputs — Hyper Hash Core

**Scope:** How `hyperhash-core` constructs coinbase outputs for the **2% pool fee** and the miner payout during template generation (Phase 1).

## Overview
When building the coinbase transaction, the block subsidy + fees (**total block reward**) is split into:
- **Miner Credit:** 98% (paid to the job/miner payout address)
- **Pool Fee:** 2% (paid to the Hyper Hash treasury address)

Both outputs are included directly in the **coinbase transaction `vout[]`**. No separate payout tx is used for the pool fee.

## Configuration (pool.toml)
```toml
[payout]
# Pool fee in basis points (bps). 200 bps = 2.00%
pool_fee_bps = 200

# Bech32/legacy address for the pool treasury
pool_fee_address = "bc1q...treasury"

# Default miner payout if a job/miner hasn't provided one
default_payout_address = "bc1q...default"
```

**Notes**
- `pool_fee_bps` is an integer (e.g., 200 = 2.00%).
- A miner/job-supplied payout address overrides `default_payout_address`.

## Calculation
Let:
- `R_total_sats` = block subsidy + tx fees (sats)
- `fee_bps` = `pool_fee_bps` (default **200**)
- `R_fee_sats = floor(R_total_sats * fee_bps / 10_000)`
- `R_miner_sats = R_total_sats - R_fee_sats`

Fee uses **floor**; any remainder stays with the miner.

## Coinbase Layout
`vout[]` (recommended order: miner first, pool second):
1) **Miner payout**
   - `value`: `R_miner_sats`
   - `scriptPubKey`: from miner/job payout address
2) **Pool treasury (fee)**
   - `value`: `R_fee_sats`
   - `scriptPubKey`: from `[payout.pool_fee_address]`

Optional **OP_RETURN** (branding/telemetry) may be appended with **0 sats**.

## Example (illustrative)
Assume `R_total_sats = 6_456_789_000`, `pool_fee_bps = 200`:
- `R_fee_sats = floor(6_456_789_000 * 200 / 10_000) = 129_135_780`
- `R_miner_sats = 6_456_789_000 - 129_135_780 = 6_327_653_220`

```json
{
  "vout": [
    { "value": 63.27653220, "scriptPubKey": { "addresses": ["bc1qminer..."] } },
    { "value": 1.29135780, "scriptPubKey": { "addresses": ["bc1qtreasury..."] } }
  ]
}
```

## Structured Log Event
```json
{
  "event": "coinbase_built",
  "height": n,
  "template_id": "<uuid>",
  "total_reward_sats": 6456789000,
  "pool_fee_bps": 200,
  "pool_fee_sats": 129135780,
  "miner_credit_sats": 6327653220,
  "miner_address": "bc1qminer...",
  "pool_fee_address": "bc1qtreasury..."
}
```

## Operator Checklist
1. **Config:** `pool_fee_bps=200` and a valid `pool_fee_address`.
2. **Decode:** `decoderawtransaction` → two pay-to outputs with expected sats.
3. **Rounding:** `pool_fee_sats + miner_credit_sats == total_reward_sats`.
4. **Audit Logs:** Structured log matches decoded coinbase.

## FAQs
- **Why bps?** Avoids floating-point; 200 bps = 2.00%.
- **Fees-only blocks?** Same formula; `R_total_sats` is subsidy+fees at height `n`.
- **Output order?** May be swapped, but keep stable across releases.
- **>2% due to rounding?** No—fee uses floor; remainder stays with miner.

---
*Owner:* Core Team • *Phase:* 1 (Coinbase/Template Engine) • *Issue:* T018
