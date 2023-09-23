package dto

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

type TransactionDto struct {
	GasFee uint64
	Hash   *common.Hash
	Value  *big.Int
}

func NewTransactionDto(transaction *types.Transaction) *TransactionDto {
	hash := transaction.Hash()
	return &TransactionDto{
		GasFee: transaction.Gas(),
		Hash:   &hash,
		Value:  transaction.Value(),
	}
}
