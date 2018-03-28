package main

import (
	"encoding/json"
	"net/http"
)

func get_list_product(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("secret")
	user_id := r.FormValue("user_id")
	category_id := r.FormValue("category_id")
	answer := get_list_product_unit(key, user_id, category_id)
	PrintToScreen(w, answer)
}

//Для юнит-тестов
func get_list_product_unit(secret string, user_id string, category_id string) []byte {
	if !check_session(secret, user_id) {
		authAndRegFailed := FailAnswer{403, "Неправильный ключ"}
		js, err := json.Marshal(authAndRegFailed)
		checkErr(err)
		return js
	}

	rows, err := SelectDB("SELECT info_product.id, name, description, photo_link, price, count " +
		"FROM info_product INNER JOIN cost_product " +
		"ON info_product.id = cost_product.id_product WHERE info_product.id_category = " + category_id)
	if err != nil {
		authAndRegFailed := FailAnswer{501, "Серверная ошибка"}
		js, err := json.Marshal(authAndRegFailed)
		checkErr(err)
		return js
	}
	var product_list ProductList
	i := 0
	for rows.Next() {
		var id int
		var name string
		var des string
		var link string
		var price float32
		var count float32
		err := rows.Scan(&id, &name, &des, &link, &price, &count)
		if err != nil {
			authAndRegFailed := FailAnswer{502, "Серверная ошибка"}
			js, err := json.Marshal(authAndRegFailed)
			checkErr(err)
			return js
		}
		product_list.Product_list = append(product_list.Product_list, Product{
			Id:          id,
			Name:        name,
			Description: des,
			Photo_link:  link,
			Price:       price,
			Count:       count,
		})
		i++
	}

	authAndRegOK := ProductList{200, i, product_list.Product_list}
	js, err := json.Marshal(authAndRegOK)
	if err != nil {
		authAndRegFailed := FailAnswer{503, "Серверная ошибка"}
		js, err := json.Marshal(authAndRegFailed)
		checkErr(err)
		return js
	}
	return js
}
