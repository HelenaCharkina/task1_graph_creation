package main

import (
	"log"
	"regexp"
	"strings"
)

func IsUnknownFieldError(err error) bool {
	if strings.Contains(err.Error(), "unknown field") {
		return true
	}
	return false
}

func ValidateInputDataFormat(data []InputData) bool {
	for _, item := range data {
		matched, err := regexp.MatchString(`[a-zA-Z0-9_]+`, item.From)
		if err != nil {
			log.Fatalf("regex error: %+v\n", err)
		}
		if !matched {
			return false
		}
		matched, err = regexp.MatchString(`[a-zA-Z0-9_]+`, item.To)
		if err != nil {
			log.Fatalf("regex error: %+v\n", err)
		}
		if !matched {
			return false
		}
	}
	return true
}

func ValidateGraphFormat(data []InputData) bool {
	vertices := make(map[string][]InputData)
	for _, item := range data {
		if _, ok := vertices[item.To]; !ok {
			vertices[item.To] = append(vertices[item.To], item)
		} else {
			for _, v := range vertices[item.To] {
				if v.Index == item.Index {
					return false
				}
			}
			vertices[item.To] = append(vertices[item.To], item)
		}
	}
	VerticesMap = vertices
	return true
}