package ipfs

import (
	"bytes"
	"context"
	"fmt"
	"testing"

	"github.com/joho/godotenv"
)

func TestDriver(t *testing.T) {
	godotenv.Load()

	bucket := "glc001"
	endpoint := "http://127.0.0.1:5001"

	key := "hello.txt"
	data := []byte("hello ipfs!")

	driver := GetIpfsDriver(endpoint, bucket)

	ctx := context.Background()

	txHash, err := driver.Put(ctx, key, data)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("txHash", txHash)

	data0, txHash0, err := driver.Get(ctx, key)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(data, data0) {
		t.Fatal("data not match")
	}
	if txHash != txHash0 {
		t.Fatal("txHash not match")
	}

	fmt.Println("daType", driver.Type())
	fmt.Println("daID", driver.DaID(txHash, txHash))
}
