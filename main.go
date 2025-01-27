package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func MakeContacts(records [][]string) {
	for index, value := range records {
		fmt.Printf("\"%v\",\"%v\",\"%v\",\"%v\",\"%v\"\n", index, value[11], value[9], value[10], "None")
	}
}

func main() {
	records := readCsvFile("channel.csv")
	MakeContacts(records)
}
