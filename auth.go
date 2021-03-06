// auth
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func auth(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")
	pass := r.FormValue("password")

	fmt.Println("login: " + login + " pass: " + pass)

	answer := get_login(login, pass)
	PrintToScreen(w, answer)
}

//Для юнит-тестов
func get_login(login string, pass string) []byte {
	fmt.Println("AUTH_START1")
	//Поиск в бд
	rows, err := SelectDB("SELECT id FROM users WHERE phone_number='" + login + "' AND password='" + pass + "'")
	if err != nil {
		authAndRegFailed := FailAnswer{502, ERROR502}
		js, err := json.Marshal(authAndRegFailed)
		checkErr(err)
		return js
	}
	fmt.Println("AUTH_START12")
	for rows.Next() {
		var username int
		err := rows.Scan(&username)
		if err != nil {
			authAndRegFailed := FailAnswer{503, ERROR503}
			js, err := json.Marshal(authAndRegFailed)
			checkErr(err)
			return js
		}
		rows, err = SelectDB("SELECT id FROM users WHERE phone_number = '" + login + "'")
		if err != nil {
			authAndRegFailed := FailAnswer{502, ERROR502}
			js, err := json.Marshal(authAndRegFailed)
			checkErr(err)
			return js
		}
		uid := 0
		for rows.Next() {
			var user_id int
			err := rows.Scan(&user_id)
			if err != nil {
				authAndRegFailed := FailAnswer{503, ERROR503}
				js, err := json.Marshal(authAndRegFailed)
				checkErr(err)
				return js
			}
			uid = user_id
		}
		fmt.Println("AUTH_START11111")
		secret_key := start_session(strconv.Itoa(uid))
		authAndRegOK := AuthAndRegOK{200, secret_key, uid}
		js, err := json.Marshal(authAndRegOK)
		WriteLog("auth", "OK", secret_key, uid)
		if err != nil {
			authAndRegFailed := FailAnswer{500, ERROR500}
			js, err := json.Marshal(authAndRegFailed)
			checkErr(err)
			return js
		}
		return js
	}
	authAndRegFailed := FailAnswer{404, ERROR404}
	js, err := json.Marshal(authAndRegFailed)
	checkErr(err)
	return js
}
 