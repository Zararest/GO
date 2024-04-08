package main

import (
	"flag"
	"os"

	"gopkg.in/yaml.v2"

	"currency-parser/pkg/currency"
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
	var dumpParsed bool

	// Parse flags
	flag.StringVar(&configFile, "config", "",
		"File with files to parse")
	flag.BoolVar(&dumpParsed, "dump-input", false,
		"Dump all currencies")
	flag.Parse()

	if configFile == "" {
		panic("config file hasn't been provided")
	}

	inputFile, outputFile := parseYamlFile(configFile)
	currencyRepresentation, err := currency.DecodeFile(inputFile)
	if err != nil {
		panic("can't decode file: " + err.Error())
	}

	if dumpParsed {
		currency.Dump(currencyRepresentation)
	}

	currency.SortByValue(currencyRepresentation)
	err = currency.DumpToJson(currencyRepresentation, outputFile)
	if err != nil {
		panic("can't dump info into json: " + err.Error())
	}
}
