address_string=

config=../../conf/ctool.json
addr=0x0000000000000000000000000000000000000011
abi=../cnsManager.cpp.abi.json

../ctool invoke --config $config --addr $addr --abi $abi --func ifRegisteredByAddress --param $address_string 
echo "ifRegisteredByAddress"
echo address_string = $address_string
ho "ifRegisteredByAddress"
echo address_string = $address_string
