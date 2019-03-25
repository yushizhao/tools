name_string=

config=../../conf/ctool.json
addr=0x0000000000000000000000000000000000000011
abi=../cnsManager.cpp.abi.json

../ctool invoke --config $config --addr $addr --abi $abi --func getHistoryContractsByName --param $name_string 
echo "getHistoryContractsByName"
echo name_string = $name_string
echo "getHistoryContractsByName"
echo name_string = $name_string
