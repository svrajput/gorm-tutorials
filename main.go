package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Id        int
	FirstName string
	LastName  string
	Email     string
}

func main() {

	db, err := gorm.Open(mysql.Open("root:root.123@/Test_MS"), &gorm.Config{})

	if err != nil {
		panic(" Unable to connect mysql db ")
	}

	db.AutoMigrate(&User{})

	// u := User{
	// 	FirstName: "John",
	// 	LastName:  "Doe",
	// 	Email:     "John.Doe@testmail.com",
	// }
	// db.Create(&u)

	// u := User{
	// 	Id:        1,
	// 	FirstName: "John 2",
	// 	LastName:  "Doe 2",
	// 	Email:     "John.Doe2@testmail.com",
	// }
	// db.Updates(&u)

	u := User{
		Id: 1,
	}

	db.Delete(&u)

}
