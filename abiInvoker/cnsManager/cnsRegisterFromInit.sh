name_string=
version_string=

config=../../conf/ctool.json
addr=0x0000000000000000000000000000000000000011
abi=../cnsManager.cpp.abi.json

../ctool invoke --config $config --addr $addr --abi $abi --func cnsRegisterFromInit --param $name_string --param $version_string 
echo "cnsRegisterFromInit"
echo name_string = $name_string
echo version_string = $version_string
name_string = $name_string
echo version_string = $version_string
