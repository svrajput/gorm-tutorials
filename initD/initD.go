package initD

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	DB, err = gorm.Open(mysql.Open("root:root.123@/Test_MS"), &gorm.Config{})

	if err != nil {
		panic(" Could not connect mysql DB ")
	}

}
