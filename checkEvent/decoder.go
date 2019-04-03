package main

import (
	"bytes"
	"flag"
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"
)

type Notify struct {
	Code uint
	Msg  string
}

func main() {
	var tx = flag.String("tx", "0x", "tx hash")
	var url = flag.String("url", "http://localhost:6789", "eth rpc ip:port")

	flag.Parse()

	client, err := rpc.Dial(*url)
	if err != nil {
		fmt.Println("rpc.Dial err", err)
		return
	}

	var receipt types.Receipt
	err = client.Call(&receipt, "eth_getTransactionReceipt", *tx)

	if err != nil {
		fmt.Println("client.Call err", err)
		return
	}

	fmt.Printf("Status: %v\n", receipt.Status)

	for i, log := range receipt.Logs {
		var event Notify
		err := rlp.Decode(bytes.NewReader(log.Data), &event)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("event %d: %#v\n", i, event)
	}

	fmt.Println("end of logs.")

}
