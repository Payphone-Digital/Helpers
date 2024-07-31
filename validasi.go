package helpers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

type DataError struct {
	DATA    string `json:"data"`
	ERROR   string `json:"error"`
	MESSAGE string `json:"message"`
}

type Validator struct{}

func (v *Validator) Validate(data map[string]interface{}, rules []map[string]interface{}, num string) (map[string]interface{}, string, error) {
	validate := validator.New()
	var arr []DataError
	var intData []string
	var extData []string
	var tags string
	var banding int
	fieldErrors := make(map[string]bool)
	convertedData := make(map[string]interface{})

	// Collect external data keys
	for al := range data {
		extData = append(extData, al)
	}

	for idx, tag := range rules {
		field := tag["column"].(string)
		dataType := tag["type"].(string)
		intData = append(intData, field)

		value, ok := data[field]
		if !ok || isEmpty(value) {
			// Apply default value if input is empty and default is provided
			if defaultValue, hasDefault := tag["default"]; hasDefault && !isEmpty(defaultValue) {
				convertedDefaultValue, err := convertType(defaultValue, dataType)
				if err != nil {
					arr = append(arr, DataError{
						DATA:    "Data Array " + num,
						MESSAGE: fmt.Sprintf("Cannot convert default value %v to %s", defaultValue, dataType),
						ERROR:   field + " Type conversion error",
					})
					continue
				}
				value = convertedDefaultValue
			} else {
				if !fieldErrors[field] {
					arr = append(arr, DataError{
						DATA:    "Data Array " + num,
						MESSAGE: "Not found " + field,
						ERROR:   field + " Field not found",
					})
					fieldErrors[field] = true
				}
				continue
			}
		}

		// Convert data type
		convertedValue, err := convertType(value, dataType)
		if err != nil {
			arr = append(arr, DataError{
				DATA:    "Data Array " + num,
				MESSAGE: fmt.Sprintf("Cannot convert %v to %s", value, dataType),
				ERROR:   field + " Type conversion error",
			})
			continue
		}
		convertedData[field] = convertedValue

		// Skip validation if the field is empty and default is applied
		if isEmpty(value) && tag["default"] != nil && !isEmpty(tag["default"]) {
			continue
		}

		// Prepare validation tags
		for _, v := range rules[idx]["validasi"].([]interface{}) {
			vMap, _ := v.(map[string]interface{})
			if banding != idx {
				banding = idx
				tags = ""
			}
			if rules[idx]["column"].(string) == field {
				if val, hasVal := vMap["value"]; hasVal && val != "" {
					tags = tags + "," + vMap["valid"].(string) + "=" + val.(string)
				} else {
					tags = tags + "," + vMap["valid"].(string)
				}
			}
		}
		tags = strings.TrimLeft(tags, ",")

		// Perform validation
		if !(dataType == "bool" && value == false) {
			if err := validate.Var(convertedValue, tags); err != nil {
				validationError := err.(validator.ValidationErrors)
				for _, fieldError := range validationError {
					if !fieldErrors[field] {
						arr = append(arr, DataError{
							DATA:    "Data Array " + num,
							MESSAGE: getMessage(rules, field, fieldError.Tag()),
							ERROR:   field + " Error " + fieldError.Tag(),
						})
						fieldErrors[field] = true
					}
				}
			}
		}
	}

	// Check for prohibited fields
	for _, item := range extData {
		if !contains(intData, item) {
			arr = append(arr, DataError{
				DATA:    "Data Array " + num,
				MESSAGE: "This data is prohibited from being input",
				ERROR:   item + " Forbidden Field " + item,
			})
		}
	}

	// Marshal the error array to JSON
	jsonOutput, err := json.MarshalIndent(arr, "", "  ")
	return convertedData, string(jsonOutput), err
}

// Get the first matching message by column and validation type
func getMessage(data []map[string]interface{}, columnName, validationType string) string {
	for _, item := range data {
		if item["column"] == columnName {
			for _, validation := range item["validasi"].([]interface{}) {
				dV, ok := validation.(map[string]interface{})
				if ok {
					if dV["valid"] == validationType {
						return dV["message"].(string)
					}
				}
			}
		}
	}
	return ""
}

// Convert the type of the value
func convertType(value interface{}, dataType string) (interface{}, error) {
	switch dataType {
	case "string":
		return fmt.Sprintf("%v", value), nil
	case "int":
		return strconv.Atoi(fmt.Sprintf("%v", value))
	case "float":
		return strconv.ParseFloat(fmt.Sprintf("%v", value), 64)
	case "bool":
		strVal := fmt.Sprintf("%v", value)
		if strVal == "true" || strVal == "false" {
			return strconv.ParseBool(strVal)
		}
		return nil, fmt.Errorf("invalid boolean value: %v", value)
	case "date":
		if value == "now" {
			return time.Now().Format("2006-01-02"), nil
		}
		return time.Parse("2006-01-02", fmt.Sprintf("%v", value))
	case "datetime":
		if value == "now" {
			return time.Now().Format("2006-01-02 15:04:05"), nil
		}
		return time.Parse("2006-01-02 15:04:05", fmt.Sprintf("%v", value))
	default:
		return value, fmt.Errorf("unsupported data type: %s", dataType)
	}
}

// Check if the value is empty
func isEmpty(value interface{}) bool {
	switch v := value.(type) {
	case bool:
		return false
	case string:
		return v == ""
	case int, int32, int64, float32, float64:
		return v == 0
	default:
		return value == nil
	}
}
