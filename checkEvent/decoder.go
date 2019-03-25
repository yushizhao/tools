package main

import (
	"flag"
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
)

func main() {
	var tx = flag.String("tx", "0x", "tx hash")

	flag.Parse()

	client, err := rpc.Dial("http://localhost:6789")
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

	for _, l := range receipt.Logs {
		// b, err := hexutil.Decode(string(l.Data))

		// if err != nil {
		// 	fmt.Println(err)
		// }

		fmt.Println(string(l.Data))

		// var val interface{}
		// rlp.DecodeBytes(b, val)
	}

	fmt.Println("end of logs.")

}
