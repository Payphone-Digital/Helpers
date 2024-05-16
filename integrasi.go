package helpers

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/go-zoox/fetch"
)

func Gateway(data []map[string]interface{}) ([]map[string]interface{}, error) {
	// Mendefinisikan output untuk setiap jenis
	headers := make(map[string]string)
	body := make(map[string]string)
	params := make(map[string]string)
	var output []map[string]interface{}
	var err error

	// Mengolah setiap data
	for _, d := range data {
		for _, s := range d["source"].([]map[string]interface{}) {

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

		var req = d["url"].(string) + ":" + d["port"].(string) + d["uriSecondary"].(string) + d["uriPrimary"].(string) + d["uriCover"].(string)
		var opt = &fetch.Config{Headers: headers, Params: params, Body: body}

		var res *fetch.Response // Declare res variable outside the switch statement
		switch d["method"].(string) {
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

		// //Hasil Data Akhir
		var object map[string]interface{}
		var hasil []map[string]interface{}

		//Pemangkasan Kedalaman Data
		if len(d["parsing"].([]map[string]interface{})) > 0 {

			if err := json.Unmarshal(res.Body, &object); err != nil {
				return output, err
			}

			for n, v := range d["parsing"].([]map[string]interface{}) {

				jsonString, err := json.Marshal(object[v["variable"].(string)])
				if err != nil {
					return output, err
				}

				if IsObject(string(jsonString)) {
					object = object[v["variable"].(string)].(map[string]interface{})
					if (n + 1) == len(d["parsing"].([]map[string]interface{})) {
						hasil = append(hasil, object)
					}
				} else {
					if err := json.Unmarshal([]byte(jsonString), &hasil); err != nil {
						return output, err
					}
				}
			}
		} else {
			if err := json.Unmarshal(res.Body, &object); err != nil {
				return output, err
			}
			hasil = append(hasil, object)
		}

		if len(d["manipulate"].([]map[string]interface{})) > 0 {
			for _, vm := range d["manipulate"].([]map[string]interface{}) {
				for _, item := range hasil {
					val, ok := item[vm["target"].(string)]
					val1, ok1 := item[vm["compare"].(string)]

					fmt.Println(item[vm[vm["patok"].(string)].(string)])
					fmt.Println(item[vm[vm["banding"].(string)].(string)])
					if vm["type"] == "var-convert" {
						//Algoritma 1
						//Merubah Variable yang sudah ada menjadi variable lain yang dirubah target menjadi menjadi variable yang dicompare
						if vm["algoritma"].(int) == 1 && ok {
							item[vm["compare"].(string)] = val
							delete(item, vm["target"].(string))
						}
					} else if vm["type"] == "var-delete" {
						//Algoritma 1
						//Mendelete Variable yang sudah ada menjadi tidak ada didalam data target
						if vm["algoritma"].(int) == 1 {
							delete(item, vm["target"].(string))
						}
					} else if vm["type"] == "var-add" {
						// Algoritma 1
						// Penambahan object dengan value ditambahkan melalui data value saja
						// Algoritma 2
						// Penambahan object dengan value ditambahkan melalui perbandingan value dan target + - * /
						// Algoritma 3
						// Penambahan object dengan value ditambahkan melalui
						if vm["algoritma"].(int) == 1 {
							item[vm["var_new"].(string)] = vm["value"].(string)
						}
					} else if vm["type"] == "val-convert" {
						//ALgoritma 1
						//data target dirubah dengan value
						//ALgoritma 2
						//data target dirubah dengan compare
						if vm["algoritma"].(int) == 1 {
							item[vm["target"].(string)] = vm["value"].(string)
						} else if vm["algoritma"].(int) == 2 && ok1 {
							item[vm["target"].(string)] = val1.(string)
						}

					} else if vm["type"] == "val-boolean" {
						// ALgoritma 1
						// Target adalah data yang akan dirubah dan dibandingkan dengan value sehingga menghasilkan data true atau false
						// ALgoritma 2
						// Target adalah data yang akan dirubah dan dibandingkan dengan data compare sehingga menghasilkan data true atau false
						// ALgoritma 3
						// Dua Kondisi yang digabungkan
						if ok {
							switch vm["condition"] {
							case "==":
								if vm["algoritma"].(int) == 1 {
									item[vm["target"].(string)] = (val.(string) == vm["value"].(string))
								} else if vm["algoritma"].(int) == 2 && ok1 {
									item[vm["target"].(string)] = (val.(string) == val1.(string))
								}
							case "!=":
								if vm["algoritma"].(int) == 1 {
									item[vm["target"].(string)] = (val.(string) != vm["value"].(string))
								} else if vm["algoritma"].(int) == 2 && ok1 {
									item[vm["target"].(string)] = (val.(string) != val1.(string))
								}
							case ">":
								if vm["algoritma"].(int) == 1 {
									item[vm["target"].(string)] = (val.(int) > vm["value"].(int))
								} else if vm["algoritma"].(int) == 2 && ok1 {
									item[vm["target"].(string)] = (val.(string) > val1.(string))
								}
							case ">=":
								if vm["algoritma"].(int) == 1 {
									item[vm["target"].(string)] = (val.(int) >= vm["value"].(int))
								} else if vm["algoritma"].(int) == 2 && ok1 {
									item[vm["target"].(string)] = (val.(string) >= val1.(string))
								}
							case "<":
								if vm["algoritma"].(int) == 1 {
									item[vm["target"].(string)] = (val.(int) < vm["value"].(int))
								} else if vm["algoritma"].(int) == 2 && ok1 {
									item[vm["target"].(string)] = (val.(string) < val1.(string))
								}
							case "<=":
								if vm["algoritma"].(int) == 1 {
									item[vm["target"].(string)] = (val.(int) <= vm["value"].(int))
								} else if vm["algoritma"].(int) == 2 && ok1 {
									item[vm["target"].(string)] = (val.(string) <= val1.(string))
								}
							}
						}
					} else if vm["type"] == "val-pembagian" {
						// ALgoritma 1
						// Target adalah data yang akan dirubah dan dibandingkan dengan value sehingga menghasilkan data true atau false
						// ALgoritma 2
						// Target adalah data yang akan dirubah dan dibandingkan dengan data compare sehingga menghasilkan data true atau false
						if ok {
							if vm["algoritma"].(int) == 1 {
								item[vm["target"].(string)] = val.(int) / vm["value"].(int)
							} else if vm["algoritma"].(int) == 2 && ok1 {
								item[vm["target"].(string)] = val.(int) / val1.(int)
							}
						}
					} else if vm["type"] == "val-perkalian" {
						//Algoritma 1
						//data target dikali dengan value
						//Algoritma 2
						//data target tambah dengan compare
						if ok {
							if vm["algoritma"].(int) == 1 {
								item[vm["target"].(string)] = val.(int) * vm["value"].(int)
							} else if vm["algoritma"].(int) == 2 && ok1 {
								item[vm["target"].(string)] = val.(int) * val1.(int)
							}
						}
					} else if vm["type"] == "val-pengurangan" {
						//Algoritma 1
						//data target dikurangi dengan value
						//Algoritma 2
						//data target tambah dengan compare
						if ok {
							if vm["algoritma"].(int) == 1 {
								item[vm["target"].(string)] = val.(int) - vm["value"].(int)
							} else if vm["algoritma"].(int) == 2 && ok1 {
								item[vm["target"].(string)] = val.(int) - val1.(int)
							}
						}
					} else if vm["type"] == "val-penambahan" {
						//Algoritma 1
						//data target tambah dengan value
						//Algoritma 2
						//data target tambah dengan compare
						if ok {
							if vm["algoritma"].(int) == 1 {
								item[vm["target"].(string)] = val.(int) + vm["value"].(int)
							} else if vm["algoritma"].(int) == 2 && ok1 {
								item[vm["target"].(string)] = val.(int) + val1.(int)
							}
						}
					}

				}
			}
		}
		output = append(output, hasil...)
	}
	return output, err
}
