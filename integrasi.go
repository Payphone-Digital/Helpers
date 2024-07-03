package helpers

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"fmt"

	"github.com/go-zoox/fetch"
)

func Gateway(data map[string]interface{}) ([]map[string]interface{}, error) {
	// Mendefinisikan output untuk setiap jenis
	headers := make(map[string]string)
	body := make(map[string]string)
	params := make(map[string]string)
	var output []map[string]interface{}
	var err error

	// Mengolah setiap data
	// for _, d := range data {
	for _, s := range data["source"].([]map[string]interface{}) {

		// Mengambil nilai untuk dienkripsi
		var valueToEncrypt string
		for _, v := range s["data"].([]map[string]interface{}) {
			valueToEncrypt += fmt.Sprintf("%v", v["value"])
		}

		// Melakukan enkripsi jika diperlukan
		var encryptedValue string
		if s["encrypt"].(string) == "sha256" {
			hash := sha256.New()
			hash.Write([]byte(valueToEncrypt))
			encryptedValue = hex.EncodeToString(hash.Sum(nil))
		} else if s["encrypt"].(string) == "md5" {
			hash := md5.Sum([]byte(valueToEncrypt))
			encryptedValue = hex.EncodeToString(hash[:])
		} else {
			encryptedValue = valueToEncrypt
		}

		// Memasukkan nilai ke output sesuai jenisnya
		switch s["type"].(string) {
		case "header":
			headers[s["nama"].(string)] = encryptedValue
		case "body":
			body[s["nama"].(string)] = encryptedValue
		case "param":
			params[s["nama"].(string)] = encryptedValue
		}

	}

	var req = data["url"].(string) + ":" + data["port"].(string) + data["uriSecondary"].(string) + data["uriPrimary"].(string) + data["uriCover"].(string)
	var opt = &fetch.Config{Headers: headers, Params: params, Body: body}

	var res *fetch.Response // Declare res variable outside the switch statement
	switch data["method"].(string) {
	case "GET":
		res, err = fetch.Get(req, opt)
	case "POST":
		res, err = fetch.Post(req, opt)
	case "PUT":
		res, err = fetch.Put(req, opt)
	case "DELETE":
		res, err = fetch.Delete(req, opt)
	}

	if err != nil {
		return output, err
	}

	//Hasil Data Akhir
	var object map[string]interface{}
	// var hasil []map[string]interface{}

	if data["xmlOutput"].(bool) {
		//XML Output Api
		var root Node
		if err := xml.Unmarshal(res.Body, &root); err != nil {
			return output, err
		}
		object = root.ToMap()
	} else {
		//Json Output Api
		if err := json.Unmarshal(res.Body, &object); err != nil {
			return output, err
		}
	}

	// //Pemangkasan Kedalaman Data
	// if len(data["parsing"].([]map[string]interface{})) > 0 {
	// 	for n, v := range data["parsing"].([]map[string]interface{}) {

	// 		jsonString, err := json.Marshal(object[v["variable"].(string)])
	// 		if err != nil {
	// 			return output, err
	// 		}

	// 		if IsObject(string(jsonString)) {
	// 			object = object[v["variable"].(string)].(map[string]interface{})
	// 			if (n + 1) == len(data["parsing"].([]map[string]interface{})) {
	// 				hasil = append(hasil, object)
	// 			}
	// 		} else {
	// 			if err := json.Unmarshal([]byte(jsonString), &hasil); err != nil {
	// 				return output, err
	// 			}
	// 		}
	// 	}
	// } else { hasil = append(hasil, object) }

	// output = append(output, hasil...)
	// }
	// return output, err
}
