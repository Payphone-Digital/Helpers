package helpers

// import "fmt"

// func Manipulasi(data, manipulate []map[string]interface{}) []map[string]interface{} {
// 	// var hasil []map[string]interface{}

// 	// if len(manipulate) > 0 {
// 	for _, vm := range manipulate {
// 		for _, item := range data {
// 			val, ok := item[vm["target"].(string)]
// 			val1, ok1 := item[vm["compare"].(string)]

// 			fmt.Println(item[vm[vm["patok"].(string)].(string)])
// 			fmt.Println(item[vm[vm["banding"].(string)].(string)])
// 			if vm["type"] == "var-convert" {
// 				//Algoritma 1
// 				//Merubah Variable yang sudah ada menjadi variable lain yang dirubah target menjadi menjadi variable yang dicompare
// 				if vm["algoritma"].(int) == 1 && ok {
// 					item[vm["compare"].(string)] = val
// 					delete(item, vm["target"].(string))
// 				}
// 			} else if vm["type"] == "var-delete" {
// 				//Algoritma 1
// 				//Mendelete Variable yang sudah ada menjadi tidak ada didalam data target
// 				if vm["algoritma"].(int) == 1 {
// 					delete(item, vm["target"].(string))
// 				}
// 			} else if vm["type"] == "var-add" {
// 				// Algoritma 1
// 				// Penambahan object dengan value ditambahkan melalui data value saja
// 				// Algoritma 2
// 				// Penambahan object dengan value ditambahkan melalui perbandingan value dan target + - * /
// 				// Algoritma 3
// 				// Penambahan object dengan value ditambahkan melalui
// 				if vm["algoritma"].(int) == 1 {
// 					item[vm["var_new"].(string)] = vm["value"].(string)
// 				}
// 			} else if vm["type"] == "val-convert" {
// 				//ALgoritma 1
// 				//data target dirubah dengan value
// 				//ALgoritma 2
// 				//data target dirubah dengan compare
// 				if vm["algoritma"].(int) == 1 {
// 					item[vm["target"].(string)] = vm["value"].(string)
// 				} else if vm["algoritma"].(int) == 2 && ok1 {
// 					item[vm["target"].(string)] = val1.(string)
// 				}

// 			} else if vm["type"] == "val-boolean" {
// 				// ALgoritma 1
// 				// Target adalah data yang akan dirubah dan dibandingkan dengan value sehingga menghasilkan data true atau false
// 				// ALgoritma 2
// 				// Target adalah data yang akan dirubah dan dibandingkan dengan data compare sehingga menghasilkan data true atau false
// 				// ALgoritma 3
// 				// Dua Kondisi yang digabungkan
// 				if ok {
// 					switch vm["condition"] {
// 					case "==":
// 						if vm["algoritma"].(int) == 1 {
// 							item[vm["target"].(string)] = (val.(string) == vm["value"].(string))
// 						} else if vm["algoritma"].(int) == 2 && ok1 {
// 							item[vm["target"].(string)] = (val.(string) == val1.(string))
// 						}
// 					case "!=":
// 						if vm["algoritma"].(int) == 1 {
// 							item[vm["target"].(string)] = (val.(string) != vm["value"].(string))
// 						} else if vm["algoritma"].(int) == 2 && ok1 {
// 							item[vm["target"].(string)] = (val.(string) != val1.(string))
// 						}
// 					case ">":
// 						if vm["algoritma"].(int) == 1 {
// 							item[vm["target"].(string)] = (val.(int) > vm["value"].(int))
// 						} else if vm["algoritma"].(int) == 2 && ok1 {
// 							item[vm["target"].(string)] = (val.(string) > val1.(string))
// 						}
// 					case ">=":
// 						if vm["algoritma"].(int) == 1 {
// 							item[vm["target"].(string)] = (val.(int) >= vm["value"].(int))
// 						} else if vm["algoritma"].(int) == 2 && ok1 {
// 							item[vm["target"].(string)] = (val.(string) >= val1.(string))
// 						}
// 					case "<":
// 						if vm["algoritma"].(int) == 1 {
// 							item[vm["target"].(string)] = (val.(int) < vm["value"].(int))
// 						} else if vm["algoritma"].(int) == 2 && ok1 {
// 							item[vm["target"].(string)] = (val.(string) < val1.(string))
// 						}
// 					case "<=":
// 						if vm["algoritma"].(int) == 1 {
// 							item[vm["target"].(string)] = (val.(int) <= vm["value"].(int))
// 						} else if vm["algoritma"].(int) == 2 && ok1 {
// 							item[vm["target"].(string)] = (val.(string) <= val1.(string))
// 						}
// 					}
// 				}
// 			} else if vm["type"] == "val-pembagian" {
// 				// ALgoritma 1
// 				// Target adalah data yang akan dirubah dan dibandingkan dengan value sehingga menghasilkan data true atau false
// 				// ALgoritma 2
// 				// Target adalah data yang akan dirubah dan dibandingkan dengan data compare sehingga menghasilkan data true atau false
// 				if ok {
// 					if vm["algoritma"].(int) == 1 {
// 						item[vm["target"].(string)] = val.(int) / vm["value"].(int)
// 					} else if vm["algoritma"].(int) == 2 && ok1 {
// 						item[vm["target"].(string)] = val.(int) / val1.(int)
// 					}
// 				}
// 			} else if vm["type"] == "val-perkalian" {
// 				//Algoritma 1
// 				//data target dikali dengan value
// 				//Algoritma 2
// 				//data target tambah dengan compare
// 				if ok {
// 					if vm["algoritma"].(int) == 1 {
// 						item[vm["target"].(string)] = val.(int) * vm["value"].(int)
// 					} else if vm["algoritma"].(int) == 2 && ok1 {
// 						item[vm["target"].(string)] = val.(int) * val1.(int)
// 					}
// 				}
// 			} else if vm["type"] == "val-pengurangan" {
// 				//Algoritma 1
// 				//data target dikurangi dengan value
// 				//Algoritma 2
// 				//data target tambah dengan compare
// 				if ok {
// 					if vm["algoritma"].(int) == 1 {
// 						item[vm["target"].(string)] = val.(int) - vm["value"].(int)
// 					} else if vm["algoritma"].(int) == 2 && ok1 {
// 						item[vm["target"].(string)] = val.(int) - val1.(int)
// 					}
// 				}
// 			} else if vm["type"] == "val-penambahan" {
// 				//Algoritma 1
// 				//data target tambah dengan value
// 				//Algoritma 2
// 				//data target tambah dengan compare
// 				if ok {
// 					if vm["algoritma"].(int) == 1 {
// 						item[vm["target"].(string)] = val.(int) + vm["value"].(int)
// 					} else if vm["algoritma"].(int) == 2 && ok1 {
// 						item[vm["target"].(string)] = val.(int) + val1.(int)
// 					}
// 				}
// 			}

// 		}
// 	}
// 	// }
// 	return data
// }
