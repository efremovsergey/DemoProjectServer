package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func Inc_Test(t int) int {
	return t + 1
}

func Inc_Failed_Test(f int) int {
	return f + 1
}

func testing(w http.ResponseWriter, r *http.Request) {
	failed_test_count := 0
	total_test := 0

	//1 авторизация не существующего пользователя
	fmt.Println("TEST1")
	auth_data := AuthAndRegOK{}
	answer := get_login("SergoGarage", "ssdd")
	bytes := []byte(answer)
	json.Unmarshal(bytes, &auth_data)
	total_test = Inc_Test(total_test)
	if auth_data.Code == 200 {
		failed_test_count = Inc_Failed_Test(failed_test_count)
		w.Write([]byte("1 Failed\n"))
	}
	fmt.Println("TEST2")
	//2 авторизация существующего пользователя
	auth_data_true := AuthAndRegOK{}
	answer = get_login("8(902) 107-5295", "12345678")
	bytes = []byte(answer)
	json.Unmarshal(bytes, &auth_data_true)
	total_test = Inc_Test(total_test)
	if auth_data_true.Code != 200 {
		failed_test_count = Inc_Failed_Test(failed_test_count)
		w.Write([]byte("2 Failed\n"))
	}
	//Авторизация уже вошедшего в аккаунт пользователя
	//3 авторизация неправильный пароль
	fmt.Println("TEST3")
	auth_data = AuthAndRegOK{}
	answer = get_login("8(902) 107-52-95", "1234")
	bytes = []byte(answer)
	json.Unmarshal(bytes, &auth_data)
	total_test = Inc_Test(total_test)
	if auth_data.Code != 401 {
		failed_test_count = Inc_Failed_Test(failed_test_count)
		w.Write([]byte("3 Failed\n"))
	}
	//4 регистрируем существющего пользователя
	fmt.Println("TEST4")
	reg_data := AuthAndRegOK{}
	answer = get_reg("8(111) 111-11-11", "Test", "123456", "yo", "123")
	bytes = []byte(answer)
	json.Unmarshal(bytes, &reg_data)
	total_test = Inc_Test(total_test)
	if reg_data.Code == 200 {
		failed_test_count = Inc_Failed_Test(failed_test_count)
		w.Write([]byte("4 Failed\n"))
	}
	//5 регистрация с некорректными полями города
	fmt.Println("TEST5")
	reg_data = AuthAndRegOK{}
	answer = get_reg("8(111) 111-11-12", "Test", "123456", "2", "123")
	bytes = []byte(answer)
	json.Unmarshal(bytes, &reg_data)
	total_test = Inc_Test(total_test)
	if reg_data.Code != 404 {
		failed_test_count = Inc_Failed_Test(failed_test_count)
		w.Write([]byte("5 Failed\n"))
	}
	//6 регистрация с некорректными полями склада
	fmt.Println("TEST6")
	reg_data = AuthAndRegOK{}
	answer = get_reg("8(111) 111-11-13", "Test", "123456", "yo", "werwe")
	bytes = []byte(answer)
	json.Unmarshal(bytes, &reg_data)
	total_test = Inc_Test(total_test)
	if reg_data.Code != 402 {
		failed_test_count = Inc_Failed_Test(failed_test_count)
		w.Write([]byte("6 Failed \n"))
	}
	//7 регистрация номера в белом списке
	/*reg_data = AuthAndRegOK{}
	answer = get_reg("8(902) 222-22-22", "Test", "123456", "yo", "123")
	bytes = []byte(answer)
	json.Unmarshal(bytes, &reg_data)
	total_test = Inc_Test(total_test)
	w.Write([]byte(strconv.Itoa(reg_data.Code)))
	if reg_data.Code != 200 {
		failed_test_count = Inc_Failed_Test(failed_test_count)
		w.Write([]byte("7 Failed \n"))
	}*/
	//8 регистрация номера не в белом списке
	fmt.Println("TEST7")
	reg_data = AuthAndRegOK{}
	answer = get_reg("8(902) 222-22-32", "Test", "123456", "yo", "123")
	bytes = []byte(answer)
	json.Unmarshal(bytes, &reg_data)
	total_test = Inc_Test(total_test)
	if reg_data.Code != 402 {
		failed_test_count = Inc_Failed_Test(failed_test_count)
		w.Write([]byte("8 Failed \n"))
	}
	//9 получить список категорий
	fmt.Println("TEST9")
	categorys := CategoryList{}
	answer = get_category_unit(auth_data_true.SecretCode, "1")
	bytes = []byte(answer)
	json.Unmarshal(bytes, &categorys)
	total_test = Inc_Test(total_test)
	if categorys.Code != 200 {
		failed_test_count = Inc_Failed_Test(failed_test_count)
		w.Write([]byte("9 Failed \n"))
	}
	//10 проверка длины массива
	fmt.Println(auth_data_true.SecretCode)
	total_test = Inc_Test(total_test)
	if categorys.Length != 4 {
		failed_test_count = Inc_Failed_Test(failed_test_count)
		fmt.Println("FAIL")
		w.Write([]byte("10 Failed \n"))
	}
	//11 проверка первого элемента категорий
	total_test = Inc_Test(total_test)
	fmt.Println(strconv.Itoa(categorys.Length))
	if categorys.Category[0].Name != "Заготовки" {
		failed_test_count = Inc_Failed_Test(failed_test_count)
		w.Write([]byte("11 Failed \n"))
	}
	//11 проверка четвертого элемента категорий
	total_test = Inc_Test(total_test)
	if categorys.Category[3].Name != "Пульты" {
		failed_test_count = Inc_Failed_Test(failed_test_count)
		w.Write([]byte("11 Failed \n"))
	}
	//12 получить список в категории 1
	products := ProductList{}
	answer = get_list_product_unit(auth_data_true.SecretCode, "1", "1")
	bytes = []byte(answer)
	json.Unmarshal(bytes, &products)
	total_test = Inc_Test(total_test)
	if products.Code != 200 {
		failed_test_count = Inc_Failed_Test(failed_test_count)
		w.Write([]byte("12 Failed \n"))
	}
	//13 проверка длины массива
	total_test = Inc_Test(total_test)
	if products.Length != 1 {
		failed_test_count = Inc_Failed_Test(failed_test_count)
		w.Write([]byte("13 Failed \n"))
	}
	//14 проверка получаемой информации о товаре
	productInfo := ProductInfo{}
	answer = get_product_info_unit(auth_data_true.SecretCode, "1", strconv.Itoa(products.Product_list[0].Id))
	bytes = []byte(answer)
	json.Unmarshal(bytes, &productInfo)
	total_test = Inc_Test(total_test)
	if productInfo.Code != 200 {
		failed_test_count = Inc_Failed_Test(failed_test_count)
		w.Write([]byte("14 Failed \n"))
	}
	total_test = Inc_Test(total_test)
	//15 стоимость доставки товара - 500
	if productInfo.Devilery != 500 {
		failed_test_count = Inc_Failed_Test(failed_test_count)
		w.Write([]byte("15 Failed \n"))
	}
	total_test = Inc_Test(total_test)
	//16 Имя товара - RW1990s
	if productInfo.Name != "RW1990s" {
		failed_test_count = Inc_Failed_Test(failed_test_count)
		w.Write([]byte("15 Failed \n"))
	}
	//17 получить инфо по пользователю
	userInfo := UserInfo{}
	answer = get_user_info_unit(auth_data_true.SecretCode, "1")
	bytes = []byte(answer)
	json.Unmarshal(bytes, &userInfo)
	total_test = Inc_Test(total_test)
	if userInfo.Code != 200 {
		failed_test_count = Inc_Failed_Test(failed_test_count)
		w.Write([]byte("17 Failed \n"))
	}
	//18 имя - Sergey
	total_test = Inc_Test(total_test)
	if userInfo.Name != "Sergey" {
		failed_test_count = Inc_Failed_Test(failed_test_count)
		w.Write([]byte("18 Failed \n"))
	}
	w.Write([]byte("Total test - " + strconv.Itoa(total_test) + ". Failed test - " + strconv.Itoa(failed_test_count)))
}
