# ./checkEvent -h
Usage of ./checkEvent:
+  -tx string
    	tx hash (default "0x")
+  -url string
    	eth rpc ip:port (default "http://localhost:6789")

# check events on the chain
On receiving a transaction which invokes a method call, smart contracts may record the return value of the call into 'events'. These events can be found in the same block containing the transaction.

Via eth_getTransactionReceipt, transactions' events can be retrieved by its hash.

However, events are RLP encoded. This checkEvent tool will help us retrieve events and decode them.

# decode events defined in system contracts
Within system contracts, all events can be parsed into the same struct, which is pre-defined in decoder.go:
```golang
type Notify struct {
	Code uint
	Msg  string
}
``` 
This is because system contracts use
```c++
BCOS_EVENT(Notify, uint64_t, const char*)
``` 
to define one event template and reuse it in all methods.

Note, neither the name of the event ,'Notify', nor the type of data, 'uint64_t, const char*' are included in (the byte stream of) events. 

Note, to correctly parse an event, the name of the event is irrelevant, however, the data structure need to be known in advance.

# decode arbitrary events
As stated in the previous note, the data structure of events need to be known. It will be helpful if we have something similar to *.abi.json called *.event.json.

So far, search in the contract source code is the way to go.

For example, if there is a event template says:
```c++
BCOS_EVENT(EXAMPLE, uint64_t,  const char*, uint64_t,)
``` 

we can decode it into
```golang
type EXAMPLE struct {
    Number1 uint
    String  string
    Number2 uint
}
``` 
use rlp.Decode. See https://godoc.org/github.com/ethereum/go-ethereum/rlp#Decode for more details.

Now there are some tricky parts in our c++ wasm contracts.

## signed int
As we all know, rlp don't do signed int. But it is possible to define an event template via BCOS_EVENT using signed int as variables. 

This is because in bcoslib, for events, all kinds of signed int inputs will be explicit type converted to 64-bit unsigned int. See event.hpp for more details.

Therefore, to decode them in golang, use uint.

## single value event
If there is an event template says:
```c++
BCOS_EVENT(ID, uint64_t)
``` 

use
```golang
type ID struct {
    Number uint
}
``` 
instead of uint directly.

This is because the expansion of the macro BCOS_EVENT is recursive through every input. Therefore, even if these is only one input, the list of inputs will be generated; This list will contain only one elements and be encoded as an event.