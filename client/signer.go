package client

import (
	"github.com/Sal-Ali/tron-grpc/address"
	"github.com/Sal-Ali/tron-grpc/core"
)

type Signer interface {
	Address() address.Address
	PublicKey() []byte
	SignTransaction(tx *core.Transaction) ([]byte, error)
	SignMessage(msg string) ([]byte, error)
}
