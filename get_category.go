package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func get_category(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("secret")
	user_id := r.FormValue("user_id")
	answer := get_category_unit(key, user_id)
	PrintToScreen(w, answer)
}

//Для юнит-тестов
//Возвращать id категории
func get_category_unit(secret string, user_id string) []byte {
	if !check_session(secret, user_id) {
		authAndRegFailed := FailAnswer{403, "Неправильный ключ"}
		js, err := json.Marshal(authAndRegFailed)
		checkErr(err)
		return js
	}
	rows, err := SelectDB("SELECT category from category;")
	if err != nil {
		authAndRegFailed := FailAnswer{501, "Серверная ошибка"}
		js, err := json.Marshal(authAndRegFailed)
		checkErr(err)
		return js
	}
	var category_list CategoryList
	i := 0
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			authAndRegFailed := FailAnswer{502, "Серверная ошибка"}
			js, err := json.Marshal(authAndRegFailed)
			checkErr(err)
			return js
		}
		category_list.Category = append(category_list.Category, CategoryName{
			Name: name,
		})
		i++
	}

	authAndRegOK := CategoryList{200, i, category_list.Category}
	code, err := strconv.Atoi(user_id)
	checkErr(err)
	WriteLog("get_category", "OK", secret, code)
	js, err := json.Marshal(authAndRegOK)
	if err != nil {
		authAndRegFailed := FailAnswer{503, "Серверная ошибка"}
		js, err := json.Marshal(authAndRegFailed)
		checkErr(err)
		return js
	}
	return js
}
                                                                                                                               