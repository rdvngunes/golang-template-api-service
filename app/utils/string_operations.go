package utils

import (
	"fmt"
	"strconv"
	"strings"
)

// For adjacent areas pdf operations
func FindMatchingString(Rpas string, target string) bool {
	Rpasarray := strings.Split(strings.ReplaceAll(Rpas, " ", ""), ",") //(Rpas, ",")
	for _, value := range Rpasarray {
		if value == target {
			return true
		}
	}
	return false
}

// remove element from the string

func RemoveMatchingString(Rpas string, target string) string {
	Rpasarray := strings.Split(strings.ReplaceAll(Rpas, " ", ""), ",") //(Rpas, ",")
	for i, value := range Rpasarray {
		if value == target {
			Rpasarray = append(Rpasarray[:i], Rpasarray[i+1:]...)
			break
		}
	}
	finalRpasarray := strings.Join(Rpasarray, ",")
	return finalRpasarray
}

// split the string by prefix label

func GetLinkByLabel(arr []string, targetLabel string) string {
	for _, element := range arr {
		if strings.HasPrefix(element, targetLabel) {
			return strings.TrimPrefix(element, targetLabel)
		}
	}
	return ""
}

func SeparateString(str string) (float64, string) {
	// Find the index of the first non-digit character
	idx := strings.IndexFunc(str, func(r rune) bool { return r < '0' || r > '9' })

	// Separate the number and letter parts
	numberStr := str[:idx]
	letterStr := str[idx:]

	// Convert the number part to an integer
	number, err := strconv.Atoi(numberStr)
	if err != nil {
		fmt.Println("Error:", err)
	}

	return float64(number), letterStr
}

// check if input string is present in the list of strings

func ContainsAnyField(input string, stringList []string) bool {
	for _, field := range stringList {
		if strings.Contains(input, field) {
			return true
		}
	}
	return false
}
