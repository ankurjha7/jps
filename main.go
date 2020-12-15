package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"./jps"
)

func main() {
	matrix := loadCsv()
	log.Printf("matrix loaded with y %d and x %d", len(matrix[0]), len(matrix))
	path, err := jps.AStarWithJump(matrix, jps.GetNode(3, 2), jps.GetNode(0, 0), 1)
	if err != nil {
		log.Printf("%s", err.Error())
	} else {
		for _, pathNode := range path.Nodes {
			log.Printf("%d %d", pathNode.GetX(), pathNode.GetY())
		}
	}
}

func loadCsv() [][]uint8 {
	csvFile, err := os.Open("map.csv")
	if err != nil {
		log.Panic(err)
	}

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		log.Panic(err)
	}
	matrix := make([][]uint8, 0)
	for _, line := range csvLines {
		row := make([]uint8, 0)
		for _, val := range line {
			intVal, _ := strconv.Atoi(val)
			row = append(row, uint8(intVal))
		}
		matrix = append(matrix, row)
	}
	return matrix
}
