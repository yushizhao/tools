address_string=
name_string=
version_string=

config=../../conf/ctool.json
addr=0x0000000000000000000000000000000000000011
abi=../cnsManager.cpp.abi.json

../ctool invoke --config $config --addr $addr --abi $abi --func cnsRegister --param $name_string --param $version_string --param $address_string 
echo "cnsRegister"
echo version_string = $version_string
echo address_string = $address_string
echo name_string = $name_string
n_string = $version_string
echo address_string = $address_string
