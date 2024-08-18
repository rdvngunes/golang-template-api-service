package utils

import (
	"fmt"
	"reflect"
	"strings"
)

func HasMandatoryFields(json interface{}) []string {
	mandatoryFields := []string{}
	valueType := reflect.ValueOf(json)
	typeOfT := valueType.Type()

	for i := 0; i < valueType.NumField(); i++ {
		fieldName := typeOfT.Field(i).Name
		bindingTag, found := typeOfT.Field(i).Tag.Lookup("binding")
		if found && bindingTag == "required" {
			fieldValue := valueType.Field(i)

			switch fieldValue.Kind() {
			case reflect.String:
				if fieldValue.String() == "" {
					mandatoryFields = append(mandatoryFields, fieldName)
				}
			case reflect.Slice:
				if fieldValue.Len() == 0 {
					mandatoryFields = append(mandatoryFields, fieldName)
				}
			case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
				if fieldValue.Int() == 0 {
					mandatoryFields = append(mandatoryFields, fieldName)
				}
			case reflect.Uint, reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8:
				if fieldValue.Uint() == 0 {
					mandatoryFields = append(mandatoryFields, fieldName)
				}
			case reflect.Bool:
				if !fieldValue.Bool() {
					mandatoryFields = append(mandatoryFields, fieldName)
				}
			case reflect.Struct:
				if isZero(fieldValue) {
					mandatoryFields = append(mandatoryFields, fieldName)
				}
			case reflect.Ptr:
				if fieldValue.IsNil() {
					mandatoryFields = append(mandatoryFields, fieldName)
				}
			default:
				mandatoryFields = append(mandatoryFields, fieldName)

			}
		}
	}

	return mandatoryFields
}

// Helper function to check if a struct field is zero value
func isZero(v reflect.Value) bool {
	zero := reflect.Zero(v.Type())
	return v.Interface() == zero.Interface()
}

func ValidateMandatoryFields(data interface{}, mandatoryFields []string) error {
	missingFields := []string{}
	value := reflect.ValueOf(data)

	// Check if the data is a slice
	if value.Kind() != reflect.Slice {
		for _, fieldName := range mandatoryFields {
			if reflect.ValueOf(data).FieldByName(fieldName).IsZero() {
				missingFields = append(missingFields, fieldName)
			}
		}
		if len(missingFields) > 0 {
			errorMessage := fmt.Sprintf("%s", strings.Join(missingFields, ", "))
			return fmt.Errorf(errorMessage)
		} else {
			return nil
		}
	} else {
		for i := 0; i < value.Len(); i++ {
			element := value.Index(i)
			for _, fieldName := range mandatoryFields {
				field := element.FieldByName(fieldName)
				if !field.IsValid() {
					return fmt.Errorf("field %s not found in element %d", fieldName, i)
				}
				if field.Kind() == reflect.String && field.Len() == 0 {
					missingFields = append(missingFields, fieldName)
				}
			}
		}
		if len(missingFields) > 0 {
			errorMessage := fmt.Sprintf("%s", strings.Join(missingFields, ", "))
			return fmt.Errorf(errorMessage)
		} else {
			return nil
		}
	}

}
