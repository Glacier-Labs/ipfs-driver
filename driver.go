package ipfs

import (
	"context"
)

type IDriver interface {
	Put(ctx context.Context, key string, data []byte) (txHash string, err error)
	Get(ctx context.Context, key string) (data []byte, txHash string, err error)
	Type() string
	DaID(dataHash string, txHash string) string
}

func GetIpfsDriver(endpoint, bucket string) IDriver {
	return NewIpfsDriver(endpoint, bucket)
}
