package ipfs

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/ipfs/kubo/client/rpc"

	"go.uber.org/zap"
)

type IpfsDriver struct {
	Bucket string
	client *rpc.HttpApi
	logger *zap.Logger
}

func NewIpfsDriver(endpoint, bucket string) *IpfsDriver {
	client, err := rpc.NewURLApiWithClient(endpoint, http.DefaultClient)
	if err != nil {
		panic(fmt.Sprintf("da/ipfs-driver: %s", err.Error()))
	}

	// TODO set Auth
	// client.Headers.Set()

	logger, _ := zap.NewProduction()

	return &IpfsDriver{
		client: client,
		Bucket: bucket,
		logger: logger.With(zap.Any("bucket", bucket)),
	}
}

func (ipfs *IpfsDriver) Put(ctx context.Context, key string, data []byte) (txHash string, err error) {
	path := fmt.Sprintf("/%s/%s", ipfs.Bucket, key)
	err = ipfs.client.Request("files/write", path).Option("parents", true).Option("create", true).FileBody(bytes.NewReader(data)).Exec(ctx, nil)
	if err != nil {
		return
	}

	txHash, _, err = ipfs.stat(ctx, key)
	return
}

func (ipfs *IpfsDriver) stat(ctx context.Context, key string) (hash string, size int64, err error) {
	path := fmt.Sprintf("/%s/%s", ipfs.Bucket, key)
	var res struct {
		Hash string
		Size int64
	}
	// use cidv1: https://github.com/ipfs/kubo/issues/9529
	err = ipfs.client.Request("files/stat", path).Option("cid-base", "base32").Exec(ctx, &res)
	hash = res.Hash
	size = res.Size
	return
}

func (ipfs *IpfsDriver) Get(ctx context.Context, key string) (data []byte, txHash string, err error) {
	path := fmt.Sprintf("/%s/%s", ipfs.Bucket, key)
	txHash, _, err = ipfs.stat(ctx, key)
	if err != nil {
		return
	}
	resp, err := ipfs.client.Request("files/read", path).Send(ctx)
	if err != nil {
		return
	}
	defer resp.Close()
	data, err = io.ReadAll(resp.Output)
	return
}

func (ipfs *IpfsDriver) Type() string {
	return "ipfs"
}

func (ipfs *IpfsDriver) DaID(dataHash, txHash string) string {
	return fmt.Sprintf("%s://%s/%s?txHash=%s", ipfs.Type(), ipfs.Bucket, dataHash, txHash)
}
