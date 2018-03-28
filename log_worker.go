package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"
)

func WriteLog(name string, answer string, secret_key string, code int) {
	path := "./Session/" + secret_key + ".txt"
	_, err := os.Stat(path)
	file, err := os.Open(path)
	checkErr(err)

	b, err := ioutil.ReadFile(path)
	session_data := SessionData{}
	json.Unmarshal(b, &session_data)
	session_data.LogArray = append(session_data.LogArray, LogData{
		Name:   name,
		Code:   code,
		Status: answer,
		Date:   time.Now().Format("02.01.2006 15:04:05"),
	})
	js, err := json.Marshal(session_data)
	checkErr(err)

	err = ioutil.WriteFile(path, js, 0644)
	checkErr(err)

	defer file.Close()

	//if check_session(secret_key, id) {
	//	file, err := os.Open("./Session/" + secret_key + ".txt")
	//	checkErr(err)

	//}

	//fmt.Fprintf(file, id)
}
