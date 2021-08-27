package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sort"
)

func main() {
	fileName := os.Args[1]

	var inputData []InputData

	inputFile, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	var decoder *json.Decoder
	decoder = json.NewDecoder(inputFile)
	decoder.DisallowUnknownFields()

	err = decoder.Decode(&inputData)
	if err != nil {
		switch err.(type) {
		case *json.SyntaxError:
			log.Println("Input is not JSON")
			os.Exit(1)
		case *json.UnmarshalTypeError:
			log.Println("Incorrect input data format")
			os.Exit(2)
		default:
			if IsUnknownFieldError(err) {
				log.Println("Incorrect input data format")
				os.Exit(2)
			} else {
				log.Fatal("Необработанная ошибка парсинга json файла: ", err)
			}
		}

	}
	if !ValidateInputDataFormat(inputData) {
		log.Println("Incorrect input data format")
		os.Exit(2)
	}
	if !ValidateGraphFormat(inputData) {
		log.Println("Incorrect graph format")
		os.Exit(3)
	}

	for _, value := range VerticesMap {
		sort.Slice(value, func(i, j int) bool {
			return value[i].Index < value[j].Index
		})
	}

	vertices := make(map[string][]uint)
	for _, item := range inputData {
		vertices[item.From] = make([]uint, 0)
		vertices[item.To] = make([]uint, 0)
	}

	var outputData []OutputData
	for key, value := range vertices {
		outputData = append(outputData, OutputData{
			Name:      key,
			InputFrom: value,
		})
	}

	var VerticesArray []VertexData
	for key, value := range VerticesMap {
		VerticesArray = append(VerticesArray, VertexData{
			To:   key,
			From: value,
		})
	}

	for _, item := range inputData {
		if len(vertices[item.To]) > 0 {
			continue
		}
		idxVertex := GetIndexVertexData(VerticesArray, item.To)
		if idxVertex == -1 {
			log.Fatal("not found vertex")
			return
		}
		arrForVertex := VerticesArray[idxVertex].From
		for _, vertex := range arrForVertex {
			idx := GetIndexOutputData(outputData, vertex.From)
			if idx == -1 {
				log.Fatal("not found vertex")
			}
			currentIdx := GetIndexOutputData(outputData, vertex.To)
			if currentIdx == -1 {
				log.Fatal("not found vertex")
			}
			outputData[currentIdx].InputFrom = append(outputData[currentIdx].InputFrom, uint(idx))
		}
		vertices[item.To] = append(vertices[item.To], 1)
	}

	output, err := json.Marshal(outputData)
	if err != nil {
		log.Fatal("json.Marshal error : ", err)
	}
	err = ioutil.WriteFile("output.json", output, 0644)
	if err != nil {
		log.Fatal("ioutil.WriteFile error : ", err)
	}
	return
}
