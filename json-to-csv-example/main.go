package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

// Define a struct to match the JSON structure
type Record map[string]interface{}

func main() {
	jsonFile := "data.json"
	csvFile := "output.csv"

	// Read JSON file
	data, err := os.ReadFile(jsonFile)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	// Parse JSON
	var records []Record
	err = json.Unmarshal(data, &records)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Open CSV file for writing
	file, err := os.Create(csvFile)
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Extract headers from the first record
	if len(records) == 0 {
		fmt.Println("No data to write")
		return
	}

	var headers []string
	for key := range records[0] {
		headers = append(headers, key)
	}

	writer.Write(headers)

	// Write data rows
	for _, record := range records {
		var row []string
		for _, header := range headers {
			value := fmt.Sprintf("%v", record[header])
			row = append(row, value)
		}
		writer.Write(row)
	}

	fmt.Println("CSV file successfully created as", csvFile)
}
