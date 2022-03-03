// $ go get github.com/Arafatk/glot

package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file := "dataFile"

	_, err := os.Stat(file) // Checking whether a file already exists or not
	if err != nil {
		log.Panic(err)
	}

	f, err := os.Open(file)
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.FieldsPerRecord = -1 // ?
	allRecords, err := reader.ReadAll()
	if err != nil {
		log.Panic(err)
	}

	xP := []float64{}
	yP := []float64{}
	for _, rec := range allRecords {
		x, _ := strconv.ParseFloat(rec[0], 64)
		y, _ := strconv.ParseFloat(rec[1], 64)
		xP = append(xP, x)
		yP = append(yP, y)
	}

	points := [][]float64{}
	points = append(points, xP)
	points = append(points, yP)
	fmt.Println(points)

}
