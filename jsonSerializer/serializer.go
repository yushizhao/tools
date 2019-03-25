package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {

	// connect mysql
	var filePath = flag.String("f", "./raw.json", "json file")
	flag.Parse()
	jsonFile, err := os.Open(*filePath)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	var raw interface{}
	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &raw)
	jsonList := raw.([]interface{})

	// output, err := os.OpenFile("serialized", os.O_RDWR|os.O_CREATE, 0755)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer output.Close()

	for _, element := range jsonList {

		j := element.(map[string]interface{})
		bytes, err := json.Marshal(j)
		if err != nil {
			fmt.Println(err)
		}
		str := strconv.Quote(string(bytes))
		// fmt.Fprintf(output, str+"\n")
		fmt.Println(str)

	}
}
