package utils

import (
	"fmt"
	"strings"
)

// EnumValidation validates if the given value is in the list of allowed enums.
// Enums are represented as a map of strings for compatibility.
func EnumValidation(value string, enumMap map[string]string) error {
	if _, exists := enumMap[value]; exists {
		return nil
	}

	expectedEnums := make([]string, 0, len(enumMap))
	for enum := range enumMap {
		expectedEnums = append(expectedEnums, enum)
	}

	return fmt.Errorf("invalid enum type: %s. Expected enums: %s", value, strings.Join(expectedEnums, ", "))
}
