package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

type shellFlag struct {
	name  string
	value string
}

func (f shellFlag) define() string {
	return f.name + "=" + f.value + "\n"
}

func (f shellFlag) add() string {
	return "--" + f.name + " $" + f.name + " "
}

type paramList map[string]string

func (l paramList) define() string {
	tmp := ""
	for k, v := range l {
		tmp = tmp + fmt.Sprintf("%s=%s\n", k, v)
	}
	return tmp + "\n"
}

func (l paramList) echo() string {
	tmp := ""
	for p, _ := range l {
		tmp = tmp + fmt.Sprintf("echo %s = $%s\n", p, p)
	}
	return tmp
}

func (l paramList) add() string {
	tmp := ""
	for p, _ := range l {
		tmp = tmp + "--param $" + p + " "
	}
	return tmp
}

func main() {
	var contract = flag.String("name", "cnsManager", "contract name")
	var addr = flag.String("addr", "0x0000000000000000000000000000000000000011", "contract address")
	var config = flag.String("config", "../../conf/ctool.json", "config file")
	var path = flag.String("path", "../conf/contracts/", "the folder contains $name.cpp.abi.json ")
	var ctool = flag.String("ctool", "../ctool", "path to ctool")

	flag.Parse()
	// place abitester in bin
	filePath := *path + *contract + ".cpp.abi.json"

	var addrFlag = &shellFlag{"addr", *addr}
	var abiFlag = &shellFlag{"abi", "../" + filePath}
	var configFlag = &shellFlag{"config", *config}

	defineFlag := configFlag.define() + addrFlag.define() + abiFlag.define()
	addFlag := configFlag.add() + addrFlag.add() + abiFlag.add()

	var raw interface{}
	// Open our jsonFile
	jsonFile, err := os.Open(filePath)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Successfully Opened %s\n", filePath)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &raw)
	abiFile := raw.([]interface{})

	// output scripts here
	os.Mkdir(*contract, os.ModePerm)

	// map[param][default value]
	wholeList := make(paramList)
	// [shell]
	shellList := []string{}

	// individual ctool invoke
	for _, element := range abiFile {

		abi := element.(map[string]interface{})
		if abi["type"].(string) == "function" {

			functionName := abi["name"].(string)

			pList := make(paramList)
			funcFlag := "--func " + functionName + " "

			f, err := os.OpenFile("./"+*contract+"/"+functionName+".sh", os.O_RDWR|os.O_CREATE, 0755)
			if err != nil {
				fmt.Println(err)
			}
			defer f.Close()

			inputList := abi["inputs"].([]interface{})
			for _, e := range inputList {
				input := e.(map[string]interface{})
				p := input["name"].(string) + "_" + input["type"].(string)
				//To Do: set default value
				wholeList[p] = ""
				pList[p] = ""
			}

			fmt.Fprintf(f, pList.define())
			fmt.Fprintf(f, defineFlag)

			shell := "\n" + *ctool + " invoke " + addFlag + funcFlag + pList.add() + "\n"
			shell = shell + "echo \"" + functionName + "\"\n"
			shell = shell + pList.echo()
			fmt.Fprintf(f, shell)

			shellList = append(shellList, shell)
		}
	}

	// all-in-one ctool invoke
	f, err := os.OpenFile("./"+*contract+"/"+"all-in-one"+".sh", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	fmt.Fprintf(f, wholeList.define())
	fmt.Fprintf(f, defineFlag)

	for _, e := range shellList {
		fmt.Fprintf(f, e)
	}
}
