package helpers

import (
	"bytes"
	"math/rand"
)

func Implode(slices []string, sep string) string {
	result := ""
	length := len(slices)
	for i, item := range slices {
		result += item
		if i != length-1 {
			result += sep
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

func SliceEqual(src []string, dist []string) bool {
	if len(src) != len(dist) {
		return false
	}
	for idx, value := range src {
		if value != dist[idx] {
			return false
		}
	}
	return true
}

func Keys(maps map[string]interface{}) []string {
	result := []string{}
	for key, _ := range maps {
		result = append(result, key)
	}

	return result
}

func RandString(length int) string {
	template := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	return RandTemplateString(template, length)
}

func RandNumberString(length int) string {
	template := "0123456789"
	return RandTemplateString(template, length)
}

func RandLetterString(length int) string {
	template := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return RandTemplateString(template, length)
}

func RandTemplateString(template string, length int) string {
	runes := []rune(template)
	b := make([]rune, length)
	ler := len(runes)
	for i := range b {
		b[i] = runes[rand.Intn(ler)]
	}
	return string(b)
}
