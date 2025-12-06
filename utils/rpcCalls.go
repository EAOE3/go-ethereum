package utils

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

// getLatestBlockNumber returns the latest block number from the blockchain.
// It uses a static RPC URL to connect to the local node and query the current block number.
// This function is completely self-contained and doesn't require any parameters.
func getLatestBlockNumber() (uint64, error) {
	// Static RPC endpoint URL (default geth HTTP RPC port)
	const rpcURL = "http://localhost:8545"

	// Create RPC client connection
	ctx := context.Background()
	rpcClient, err := rpc.DialContext(ctx, rpcURL)
	if err != nil {
		return 0, fmt.Errorf("failed to connect to RPC endpoint %s: %w", rpcURL, err)
	}
	defer rpcClient.Close()

	// Create Ethereum client
	ethClient := ethclient.NewClient(rpcClient)
	defer ethClient.Close()

	// Get the latest block number
	blockNumber, err := ethClient.BlockNumber(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed to get block number: %w", err)
	}

	return blockNumber, nil
}
