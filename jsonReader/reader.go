package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var path = flag.String("path", "list.json", "the file")

	flag.Parse()

	var rawFile interface{}
	// Open our jsonFile
	jsonFile, err := os.Open(*path)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Successfully Opened %s\n", *path)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &rawFile)
	objList := rawFile.([]interface{})

	for _, element := range objList {
		obj := element.(map[string]interface{})
		for k, v := range obj {
			fmt.Printf("k: %v, v: %v\n", k, v)
		}
	}

}
