pageSize_int32=
origin_string=
pageNum_int32=

config=../../conf/ctool.json
addr=0x0000000000000000000000000000000000000011
abi=../cnsManager.cpp.abi.json

../ctool invoke --config $config --addr $addr --abi $abi --func getRegisteredContractsByAddress --param $origin_string --param $pageNum_int32 --param $pageSize_int32 
echo "getRegisteredContractsByAddress"
echo origin_string = $origin_string
echo pageNum_int32 = $pageNum_int32
echo pageSize_int32 = $pageSize_int32
Num_int32 = $pageNum_int32
echo pageSize_int32 = $pageSize_int32
