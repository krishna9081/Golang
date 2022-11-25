package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func readData(f string) (map[string]string, map[string]float64) {
	//CSV Reading
	fileRecords, err := os.Open(f)

	if err != nil {
		fmt.Println("File opening error", err)
	}

	reader := csv.NewReader(fileRecords)

	if _, err := reader.Read(); err != nil {
		panic(err)
	}

	records, _ := reader.ReadAll()
	err = fileRecords.Close()
	if err != nil {
		fmt.Println("An error encountered ::", err)
	}

	m := make(map[string]float64)
	p := make(map[string]string)
	for i := range records {

		m[records[i][0]], _ = strconv.ParseFloat(records[i][1], 64)
		p[records[i][0]] = records[i][2]
	}
	return p, m
}

type Dino struct {
	Name  string
	speed float64
}

func main() {

	var s float64
	_, legLength := readData("dataset1.csv")
	legTypedata, strideLength := readData("dataset2.csv")
	var res []Dino
	for k, v := range legLength {

		if val, ok := strideLength[k]; ok {

			s = (((val / v) - 1) * math.Sqrt(v*9.8))

			if legTypedata[k] == "bipedal" {
				res = append(res, Dino{Name: k, speed: s})
			}
		}

	}

	sort.SliceStable(res, func(i, j int) bool {
		return res[i].speed > res[j].speed
	})

	fmt.Println("Sorted by speed", res)
}
