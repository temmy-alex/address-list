package models

type Address struct {
	ID     int    `json:id`
	Street string `json:street`
	City   string `json:city`
	Zip    string `json:zip`
	UserID int    `json:userid`
}

type User struct {
	ID    int    `json:id`
	Name  string `json:name`
	Email string `json:email`
}
