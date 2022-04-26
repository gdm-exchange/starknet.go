package jsonrpc

import (
	"testing"
	"context"
	"math/big"

	// "github.com/dontpanicdao/caigo"
	// "github.com/dontpanicdao/caigo/types"
)

// Requires a StarkNet JSON-RPC compliant node (e.g. pathfinder)
// (ref: https://github.com/eqlabs/pathfinder)
func TestJsonRpc(t *testing.T) {
	client, err := DialContext(context.Background(), "http://localhost:9545")
	if err != nil {
		t.Errorf("Could not connect to local StarkNet node: %v\n", err)
	}
	defer client.Close()

	// block, err := client.BlockByHash(context.Background(), "0x14b05305c69bcfa91945cd2a1a0cd4d9e8879b96e57a1688843a0719afce7c2", "TXN_HASH")
	_, err = client.BlockByHash(context.Background(), "0x14b05305c69bcfa91945cd2a1a0cd4d9e8879b96e57a1688843a0719afce7c2", "FULL_TXNS")
	if err != nil {
		t.Errorf("Could not retrieve block: %v\n", err)
	}

	_, err = client.BlockByHash(context.Background(), "0x14b05305c69bcfa91945cd2a1a0cd4d9e8879b96e57a1688843a0719afce7c2", "FULL_TXN_AND_RECEIPTS")
	if err != nil {
		t.Errorf("Could not retrieve block: %v\n", err)
	}

	_, err = client.BlockByNumber(context.Background(), big.NewInt(1), "FULL_TXNS")
	if err != nil {
		t.Errorf("Could not retrieve block: %v\n", err)
	}

	_, err = client.BlockByNumber(context.Background(), big.NewInt(1), "FULL_TXN_AND_RECEIPTS")
	if err != nil {
		t.Errorf("Could not retrieve block: %v\n", err)
	}
}