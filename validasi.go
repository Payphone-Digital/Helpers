package helpers

import (
	"encoding/json"
	"strings"

	"github.com/go-playground/validator/v10"
)

type DataError struct {
	DATA    string `json:"data"`
	ERROR   string `json:"error"`
	MESSAGE string `json:"message"`
}

type Validator struct{}

func (v *Validator) Validate(data map[string]interface{}, rules []map[string]interface{}, num string) (string, error) {
	// Variabel Data Awal
	validate := validator.New()
	var arr []DataError
	var intData []string
	var extData []string
	var tags string
	var banding int

	// Mengecek data yang tidak diizinkan Masuk
	for al := range data {
		extData = append(extData, al)
	}

	for idx, tag := range rules {
		field := tag["column"].(string)
		intData = append(intData, field)

		// Mengecek Column Apakah Sudah Ada
		value, ok := data[field]
		if !ok || value == "" {
			// Cek apakah ada default value dan apakah default tidak kosong
			if defaultValue, hasDefault := tag["default"]; hasDefault && defaultValue != "" {
				value = defaultValue
			} else {
				arr = append(arr, DataError{
					DATA:    "Data Array " + num,
					MESSAGE: "Not found " + field,
					ERROR:   field + " Field not found",
				})
				continue
			}
		}

		// Menyusun Validasi Apa Yang Harus Digunakan Pada Data Json
		for _, v := range rules[idx]["validasi"].([]interface{}) {
			vMap, _ := v.(map[string]interface{})
			if banding != idx {
				banding = idx
				tags = ""
			}
			if rules[idx]["column"].(string) == field {
				if val, hasVal := vMap["value"]; hasVal {
					if valStr, ok := val.(string); ok && valStr != "" {
						tags = tags + "," + vMap["valid"].(string) + "=" + valStr
					} else {
						// Handle missing or empty value for validation
						arr = append(arr, DataError{
							DATA:    "Data Array " + num,
							MESSAGE: "Missing or invalid value for validation " + vMap["valid"].(string),
							ERROR:   field + " Error " + vMap["valid"].(string),
						})
					}
				} else {
					tags = tags + "," + vMap["valid"].(string)
				}
			}
		}
		tags = strings.TrimLeft(tags, ",")

		// Core Validasi Untuk Mengecek Validasi
		if err := validate.Var(value, tags); err != nil {
			validationError := err.(validator.ValidationErrors)
			for _, fieldError := range validationError {
				arr = append(arr, DataError{
					DATA:    "Data Array " + num,
					MESSAGE: getMessage(rules, field, fieldError.Tag()),
					ERROR:   field + " Error " + fieldError.Tag(),
				})
			}
		}
	}

	// Periksa setiap elemen dalam data yang masuk
	for _, item := range extData {
		// Periksa apakah elemen tersebut tidak ada dalam variabel internal
		if !contains(intData, item) {
			arr = append(arr, DataError{
				DATA:    "Data Array " + num,
				MESSAGE: "This data is prohibited from being input",
				ERROR:   item + " Forbiden Field " + item,
			})
		}
	}

	// Hasil Yang Akan ditampilkan ke output
	jsonOutput, err := json.MarshalIndent(arr, "", "  ")
	return string(jsonOutput), err
}

// Function to get the first matching message by column and validation type
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
