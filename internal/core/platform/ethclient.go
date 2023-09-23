package platform

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
)

func NewClient(host string) (*ethclient.Client, error) {
	client, err := ethclient.DialContext(context.Background(), host)
	if err != nil {
		return nil, err
	}

	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("block number=", block.Number().Uint64())
	return client, nil
}
