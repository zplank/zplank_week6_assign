package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type boston struct {
	mv      float64
	nox     float64
	crim    float64
	zn      float64
	indus   float64
	chas    int
	rooms   float64
	age     float64
	dis     float64
	rad     float64
	tax     float64
	ptratio float64
	lstat   float64
}

func main() {
	// Open the CSV file
	file, err := os.Open("boston.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read all the records from the CSV file
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	// Initialize a slice to hold the data
	var bostonData []boston

	// Iterate over each record and parse the fields into boston struct
	for _, record := range records {
		// Initialize a new boston struct
		var entry boston

		// Parse each field into corresponding struct field
		entry.mv, _ = strconv.ParseFloat(record[0], 64)
		entry.nox, _ = strconv.ParseFloat(record[1], 64)
		entry.crim, _ = strconv.ParseFloat(record[2], 64)
		entry.zn, _ = strconv.ParseFloat(record[3], 64)
		entry.indus, _ = strconv.ParseFloat(record[4], 64)
		entry.chas, _ = strconv.Atoi(record[5])
		entry.rooms, _ = strconv.ParseFloat(record[6], 64)
		entry.age, _ = strconv.ParseFloat(record[7], 64)
		entry.dis, _ = strconv.ParseFloat(record[8], 64)
		entry.rad, _ = strconv.ParseFloat(record[9], 64)
		entry.tax, _ = strconv.ParseFloat(record[10], 64)
		entry.ptratio, _ = strconv.ParseFloat(record[11], 64)
		entry.lstat, _ = strconv.ParseFloat(record[12], 64)

		// Append the parsed struct to the slice
		bostonData = append(bostonData, entry)
	}

	// Print the read data
	fmt.Println(bostonData)
}
