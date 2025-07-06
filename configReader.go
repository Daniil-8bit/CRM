package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type ConfigData struct {
	DbUser         string
	DbUserPassword string
	DbName         string
	AppPort        int
}

func readConfigFile() ConfigData {

	fileData, err := os.ReadFile("configStrings.json")

	if err != nil {
		fmt.Println(err)
	}

	var cd ConfigData

	err = json.Unmarshal(fileData, &cd)

	if err != nil {
		fmt.Println(err)
	}

	return cd
}
