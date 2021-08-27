package main

type InputData struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Index uint   `json:"index"`
}

type OutputData struct {
	Name      string `json:"name"`
	InputFrom []uint `json:"input_from"`
}

var VerticesMap map[string][]InputData

type VertexData struct {
	To   string
	From []InputData
}

func GetIndexOutputData(data []OutputData, s string) int {
	for i, item := range data {
		if item.Name == s {
			return i
		}
	}
	return -1
}

func GetIndexVertexData(data []VertexData, s string) int {
	for i, item := range data {
		if item.To == s {
			return i
		}
	}
	return -1
}