package main

import (
	"gorm-tutorial/initD"

	"gorm.io/gorm"
)

type Book struct {
	Id    int
	Title string
}

type Address struct {
	gorm.Model
	UserId  int
	Address string
}

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
	Address   Address `gorm:"foreignKey:UserId"`
	Books     []Book  `gorm:"many2many:user_book"`
}

func init() {
	initD.ConnectDB()
}

func main() {

	initD.DB.AutoMigrate(&User{}, &Address{}, Book{})

	u := User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "John.Doe@testmail.com",
	}
	initD.DB.Create(&u)

	// u := User{
	// 	Id:        1,
	// 	FirstName: "John 2",
	// 	LastName:  "Doe 2",
	// 	Email:     "John.Doe2@testmail.com",
	// }
	// db.Updates(&u)

	// u := User{
	// 	Id: 1,
	// }
	// db.Delete(&u)

	// Refactor DB -Done

	// Create Multiple Users
	// user := []User{
	// 	User{FirstName: "John", LastName: "Doe", Email: "John.Doe@testmail.com"},
	// 	User{FirstName: "Sunil", LastName: "Rajput", Email: "Sunil.Rajput@testmail.com"},
	// 	User{FirstName: "Sam", LastName: "Adam", Email: "Sam.Adam@testmail.com"},
	// }
	// initD.DB.Create(&user)

	// Fetch First Record from users Table

	// user := User{}
	//initD.DB.First(&user)
	//fmt.Println(user)

	// Fetch Last Record from Users Table
	//initD.DB.Last(&user)
	//fmt.Println(user)

	// Fetch record using where condition
	//initD.DB.Where("last_name", "Rajput").First(&user)
	//fmt.Println(user)

	// Fetch Record using custome SQL statement
	//initD.DB.Raw("SELECT id, first_name, last_name, Email from users where id = ?", 2).Scan(&user)
	// fmt.Println(user)

	//gorm.Model

	// Constrains

}
