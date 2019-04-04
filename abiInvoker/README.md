# ./abiInvoker -h
Usage of ./abiInvoker:
+  -addr string
    	contract address (default "0x0000000000000000000000000000000000000011")
+  -config string
    	config file (default "../../conf/ctool.json")
+  -ctool string
    	path to ctool (default "../ctool")
+  -name string
    	contract name (default "cnsManager")
+  -path string
    	the folder contains $name.cpp.abi.json  (default "../conf/contracts/")

# description
generate scripts for call contract methods from contract abi files.

## example input
```json
[
    {
        "name": "cnsRegister",
        "inputs": [
            {
                "name": "name",
                "type": "string"
            },
            {
                "name": "version",
                "type": "string"
            },
            {
                "name": "address",
                "type": "string"
            }
        ],
        "outputs": [],
        "constant": "false",
        "type": "function"
    },
]
```

## example output
```bash
name_string=$1
version_string=$2
address_string=$3

config=../../conf/ctool.json
addr=0x0000000000000000000000000000000000000011
abi=.././cnsManager.cpp.abi.json

../ctool invoke --config $config --addr $addr --abi $abi --func cnsRegister --param $name_string --param $version_string --param $address_string 
echo "cnsRegister"
echo name_string = $name_string
echo version_string = $version_string
echo address_string = $address_string
```