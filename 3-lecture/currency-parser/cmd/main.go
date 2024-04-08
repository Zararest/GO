package main

import (
	"flag"
	"os"

	"gopkg.in/yaml.v2"

	"currency-parser/pkg/xml"
)

type _YAMLStructure struct {
	InputFile string `yaml:"input-file"`
	OutFile   string `yaml:"output-file"`
}

func parseYamlFile(fileName string) (string, string) {
	yamlFile, err := os.ReadFile(fileName)
	if err != nil {
		panic("failed to read YAML file" + err.Error())
	}

	var config _YAMLStructure
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic("failed to parse YAML" + err.Error())
	}

	return config.InputFile, config.OutFile
}

func main() {
	// Define flags
	var configFile string

	// Parse flags
	flag.StringVar(&configFile, "config", "",
		"File with files to parse")
	flag.Parse()

	if configFile == "" {
		panic("config file hasn't been provided")
	}

	inputFile, _ := parseYamlFile(configFile)
	currencyRepresentation, err := xmlParser.DecodeFile(inputFile)
	if err != nil {
		panic("can't decode file:" + err.Error())
	}
	xmlParser.DumpCurrency(currencyRepresentation)
}
