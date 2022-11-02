package tradeCrypto

import (
	"bytes"
	"embed"
	"os"
)

// Embeds has the entire ./embeds directory embed in read-only mode
//go:embed embeds/* env/*
var Embeds embed.FS

func LoadEnvFromFile(filepath string) {
	fileBytes, _ := Embeds.ReadFile(filepath)
	fileBytes = bytes.Replace(fileBytes, []byte("\""), nil, -1) //remove the " char. e.g.: name="john". We will remove the double quotes
	fileBytes = bytes.Replace(fileBytes, []byte("\r"), nil, -1) //remove the " char. e.g.: name="john". We will remove the double quotes

	envVariables := bytes.Split(fileBytes, []byte("\n"))

	for _, envVariable := range envVariables {
		keyValue := bytes.SplitN(envVariable, []byte("="), -1) // we will split to be ["name", "john"] in bytes
		key := string(keyValue[0])
		value := string(keyValue[1])
		os.Setenv(key, value)
	}

}
