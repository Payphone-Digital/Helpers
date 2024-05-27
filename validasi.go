package helpers

import (
	"encoding/json"
	"strings"

	"github.com/go-playground/validator/v10"
)

type DataError struct {
	DATA    string `json:"data"`
	FIELD   string `json:"field"`
	ERROR   string `json:"error"`
	MESSAGE string `json:"message"`
}
type Validator struct{}

func (v *Validator) Validate(data map[string]interface{}, rules []map[string]interface{}, num string) (string, error) {
	//Variabel Data Awal
	validate := validator.New()
	var arr []DataError
	var intData []string
	var extData []string
	var tags string
	var banding int

	//Mengecek data yang tidak diizinkan Masuk
	for al := range data {
		extData = append(extData, al)
	}

	for idx, tag := range rules {
		field := tag["column"].(string)
		intData = append(intData, field)

		//Mengecek Column Apakah Sudah Ada
		value, ok := data[field]
		if !ok {
			arr = append(arr, DataError{
				DATA:    "Data Array " + num,
				FIELD:   field,
				ERROR:   "Field not found",
				MESSAGE: "Not found " + field,
			})
		}

		//Menyusun Validasi Apa Yang Harus Digunakan Pada Data Json
		for _, v := range rules[idx]["validasi"].([]interface{}) {
			vMap, _ := v.(map[string]interface{})
			// if !ok {
			// 	// Menangani kasus di mana nilai setiap elemen tidak sesuai dengan tipe yang diharapkan
			// 	fmt.Println("Error: Failed to convert validasi item to map[string]interface{}")
			// 	continue
			// }
			if banding != idx {
				banding = idx
				tags = ""
			}
			if rules[idx]["column"].(string) == field {
				tags = tags + "," + vMap["valid"].(string)
			}
		}
		tags = strings.TrimLeft(tags, ",")

		//Core Validasi Untuk Mengecek Validasi
		if err := validate.Var(value, tags); err != nil {
			validationError := err.(validator.ValidationErrors)
			for _, fieldError := range validationError {
				arr = append(arr, DataError{
					DATA:    "Data Array " + num,
					FIELD:   field,
					ERROR:   "Error " + fieldError.Tag(),
					MESSAGE: getMessage(rules, field, fieldError.Tag()),
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
				FIELD:   item,
				ERROR:   "Forbiden Field " + item,
				MESSAGE: "This data is prohibited from being input",
			})
		}
	}

	//Hasil Yang Akan ditampilkan ke output
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
