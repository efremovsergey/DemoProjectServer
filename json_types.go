package main

type AuthAndRegOK struct {
	Code       int
	SecretCode string
	User_id    int
}

type FailAnswer struct {
	Code        int
	Description string
}

type CategoryName struct {
	Name string
}

type CategoryList struct {
	Code     int
	Length   int
	Category []CategoryName
}

type Product struct {
	Id          int
	Name        string
	Description string
	Photo_link  string
	Price       float32
	Count       float32
}

type ProductList struct {
	Code         int
	Length       int
	Product_list []Product
}

type ProductInfo struct {
	Code        int
	Name        string
	Description string
	Photo_link  string
	Price       float32
	Count       float32
	Devilery    float32
}

type UserInfo struct {
	Code  int
	Phone string
	Name  string
	City  string
	Stock string
}

type Sucsses struct {
	Code int
}

type LogData struct {
	Name   string
	Code   int
	Status string
	Date   string
}

type SessionData struct {
	UserID    string
	SecretKey string
	StartDate string
	LogArray  []LogData
}
