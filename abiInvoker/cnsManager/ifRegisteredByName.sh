name_string=

config=../../conf/ctool.json
addr=0x0000000000000000000000000000000000000011
abi=../cnsManager.cpp.abi.json

../ctool invoke --config $config --addr $addr --abi $abi --func ifRegisteredByName --param $name_string 
echo "ifRegisteredByName"
echo name_string = $name_string
tring 
echo "ifRegisteredByName"
echo name_string = $name_string
