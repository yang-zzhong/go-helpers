package helpers

import (
	"bytes"
)

func Implode(slices []string, seperator string) string {
	result := ""
	length := len(slices)
	for i, item := range slices {
		result += item
		if i != length-1 {
			result += seperator
		}
	}

	return result
}

func Explode(str string, sep string) []string {
	result := []string{}
	for _, item := range bytes.Split(([]byte)(str), ([]byte)(sep)) {
		result = append(result, (string)(item))
	}

	return result
}
