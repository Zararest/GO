package currency

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

type Currency struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

type xmlCurrency struct {
	NumCode  int    `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type xmlFileStruct struct {
	Records []xmlCurrency `xml:"Valute"`
}

type Descending []Currency

func (d Descending) Len() int           { return len(d) }
func (d Descending) Less(i, j int) bool { return d[i].Value > d[j].Value }
func (d Descending) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }

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

func Dump(currencyToDump []Currency) {
	fmt.Println("List of currencies:")
	for i, cur := range currencyToDump {
		fmt.Printf("Currency %d\n", i)
		fmt.Printf("\t%d\n", cur.NumCode)
		fmt.Printf("\t%s\n", cur.CharCode)
		fmt.Printf("\t%f\n", cur.Value)
	}
}

func DecodeFile(xmlFileName string) ([]Currency, error) {
	xmlFile, err := os.Open(xmlFileName)
	if err != nil {
		return nil, err
	}
	defer xmlFile.Close()

	xmlDecoder := xml.NewDecoder(xmlFile)
	xmlDecoder.CharsetReader =
		func(charset string, input io.Reader) (io.Reader, error) {
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		}

	var xmlFileParsed xmlFileStruct
	err = xmlDecoder.Decode(&xmlFileParsed)
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

func SortByValue(currencies []Currency) {
	sort.Sort(Descending(currencies))
}

func DumpToJson(currencies []Currency, outFile string) error {
	file, err := os.Create(outFile)
	if err != nil {
		return err
	}
	defer file.Close()

	// Convert structure representation to json
	jsonRepresentation, err := json.Marshal(currencies)
	if err != nil {
		return err
	}

	_, err = file.Write(jsonRepresentation)
	if err != nil {
		return err
	}

	return nil
}
