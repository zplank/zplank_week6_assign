package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

type Boston struct {
	neighborhood string
	crim         float64
	zn           float64
	indus        float64
	chas         int
	nox          float64
	rooms        float64
	age          float64
	dis          float64
	rad          float64
	tax          float64
	ptratio      float64
	lstat        float64
	mv           float64
}

func main() {
	fmt.Println("Starting all 100 iterations")
	start := time.Now()

	for i := 0; i < 100; i++ {
		fmt.Printf("Run %d:\n", i+1)
		startRun := time.Now()
		run()
		fmt.Printf("Run %d finished. Time taken: %v\n", i+1, time.Since(startRun))
	}

	fmt.Printf("All runs finished. Total time taken: %v\n", time.Since(start))
}

func run() {
	//read data in from boston.csv
	data, err := readDataFromCSV("boston.csv")
	if err != nil {
		fmt.Println("Error reading data from CSV:", err)
		return
	}

	//linear regression model
	coefficients, mse := linearRegression(data)

	//print coefficients and MSE
	fmt.Println("Coefficients:")
	for feature, coefficient := range coefficients {
		fmt.Printf("%s: %.6f\n", feature, coefficient)
	}
	fmt.Printf("Mean-Square Error: %.6f\n", mse)

	//aic and bic calulations
	n := float64(len(data))
	k := float64(len(coefficients))
	aic := n*math.Log(mse) + 2*k
	bic := n*math.Log(mse) + k*math.Log(n)

	//print aic and bic
	fmt.Printf("AIC: %.6f\n", aic)
	fmt.Printf("BIC: %.6f\n", bic)
}

// read in and parse data
func readDataFromCSV(filename string) ([]Boston, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 14
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var data []Boston
	for _, record := range records {
		var b Boston
		b.neighborhood = record[0]
		b.crim, _ = strconv.ParseFloat(record[1], 64)
		b.zn, _ = strconv.ParseFloat(record[2], 64)
		b.indus, _ = strconv.ParseFloat(record[3], 64)
		b.chas, _ = strconv.Atoi(record[4])
		b.nox, _ = strconv.ParseFloat(record[5], 64)
		b.rooms, _ = strconv.ParseFloat(record[6], 64)
		b.age, _ = strconv.ParseFloat(record[7], 64)
		b.dis, _ = strconv.ParseFloat(record[8], 64)
		b.rad, _ = strconv.ParseFloat(record[9], 64)
		b.tax, _ = strconv.ParseFloat(record[10], 64)
		b.ptratio, _ = strconv.ParseFloat(record[11], 64)
		b.lstat, _ = strconv.ParseFloat(record[12], 64)
		b.mv, _ = strconv.ParseFloat(record[13], 64)
		data = append(data, b)
	}

	return data, nil
}

func linearRegression(data []Boston) (map[string]float64, float64) {
	coefficients := make(map[string]float64)

	//initiate variables
	var sumSquaredErrors float64

	for _, feature := range []string{"crim", "zn", "indus", "chas", "nox", "rooms", "age", "dis", "rad", "tax", "ptratio", "lstat"} {
		sumX := 0.0
		sumY := 0.0
		sumXY := 0.0
		sumX2 := 0.0
		n := float64(len(data))

		for _, d := range data {
			sumX += getFeatureValue(d, feature)
			sumY += d.mv
			sumXY += getFeatureValue(d, feature) * d.mv
			sumX2 += getFeatureValue(d, feature) * getFeatureValue(d, feature)
		}

		slope := (n*sumXY - sumX*sumY) / (n*sumX2 - sumX*sumX)
		coefficients[feature] = slope

		//calulate predicted values and sqr errors
		for _, d := range data {
			predicted := slope * getFeatureValue(d, feature)
			sumSquaredErrors += math.Pow(d.mv-predicted, 2)
		}
	}

	//MSE calculation
	mse := sumSquaredErrors / float64(len(data))

	return coefficients, mse
}

// return results
func getFeatureValue(b Boston, feature string) float64 {
	switch feature {
	case "crim":
		return b.crim
	case "zn":
		return b.zn
	case "indus":
		return b.indus
	case "chas":
		return float64(b.chas)
	case "nox":
		return b.nox
	case "rooms":
		return b.rooms
	case "age":
		return b.age
	case "dis":
		return b.dis
	case "rad":
		return b.rad
	case "tax":
		return b.tax
	case "ptratio":
		return b.ptratio
	case "lstat":
		return b.lstat
	default:
		return 0.0
	}
}
