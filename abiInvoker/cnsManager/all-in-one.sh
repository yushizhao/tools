address_string=
pageNum_int32=
pageSize_int32=
origin_string=
name_string=
version_string=

config=../../conf/ctool.json
addr=0x0000000000000000000000000000000000000011
abi=../cnsManager.cpp.abi.json

../ctool invoke --config $config --addr $addr --abi $abi --func cnsRegisterFromInit --param $name_string --param $version_string 
echo "cnsRegisterFromInit"
echo name_string = $name_string
echo version_string = $version_string

../ctool invoke --config $config --addr $addr --abi $abi --func cnsRegister --param $name_string --param $version_string --param $address_string 
echo "cnsRegister"
echo version_string = $version_string
echo address_string = $address_string
echo name_string = $name_string

../ctool invoke --config $config --addr $addr --abi $abi --func cnsUnregister --param $name_string --param $version_string 
echo "cnsUnregister"
echo name_string = $name_string
echo version_string = $version_string

../ctool invoke --config $config --addr $addr --abi $abi --func getContractAddress --param $name_string --param $version_string 
echo "getContractAddress"
echo name_string = $name_string
echo version_string = $version_string

../ctool invoke --config $config --addr $addr --abi $abi --func getRegisteredContracts --param $pageNum_int32 --param $pageSize_int32 
echo "getRegisteredContracts"
echo pageNum_int32 = $pageNum_int32
echo pageSize_int32 = $pageSize_int32

../ctool invoke --config $config --addr $addr --abi $abi --func getRegisteredContractsByAddress --param $origin_string --param $pageNum_int32 --param $pageSize_int32 
echo "getRegisteredContractsByAddress"
echo origin_string = $origin_string
echo pageNum_int32 = $pageNum_int32
echo pageSize_int32 = $pageSize_int32

../ctool invoke --config $config --addr $addr --abi $abi --func ifRegisteredByName --param $name_string 
echo "ifRegisteredByName"
echo name_string = $name_string

../ctool invoke --config $config --addr $addr --abi $abi --func ifRegisteredByAddress --param $address_string 
echo "ifRegisteredByAddress"
echo address_string = $address_string

../ctool invoke --config $config --addr $addr --abi $abi --func getHistoryContractsByName --param $name_string 
echo "getHistoryContractsByName"
echo name_string = $name_string
