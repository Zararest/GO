package xmlParser

import (
	"encoding/xml"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Currency struct {
	NumCode  int
	CharCode string
	Value    float64
}

type xmlCurrency struct {
	NumCode  int
	CharCode string
	Value    string
}

type xmlFileStruct struct {
	Records []xmlCurrency
}

func transformRecord(record xmlCurrency) (Currency, error) {
	var res Currency
	res.CharCode = record.CharCode
	res.NumCode = record.NumCode

	record.Value = strings.ReplaceAll(record.Value, ",", ".")
	f, err := strconv.ParseFloat(record.Value /*BitSize*/, 64)
	if err != nil {
		return res, err
	}
	res.Value = f
	return res, nil
}

func DumpCurrency(currencyToDump []Currency) {
	for i, cur := range currencyToDump {
		fmt.Printf("Currency %d\n", i)
		fmt.Printf("\t%d\n", cur.NumCode)
		fmt.Printf("\t%s\n", cur.CharCode)
		fmt.Printf("\t%f\n", cur.Value)
	}
}

func DecodeFile(xmlFileName string) ([]Currency, error) {
	xmlFile, err := os.ReadFile(xmlFileName)
	if err != nil {
		return nil, err
	}

	var xmlFileParsed xmlFileStruct
	err = xml.Unmarshal(xmlFile, &xmlFileParsed)
	if err != nil {
		return nil, err
	}

	decodedFile := make([]Currency, len(xmlFileParsed.Records))
	for i, xmlRecord := range xmlFileParsed.Records {
		decodedFile[i], err = transformRecord(xmlRecord)
		if err != nil {
			return nil, err
		}
	}
	return decodedFile, nil
}
