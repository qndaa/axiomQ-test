package db

import (
	"github.com/qndaa/backend/src/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var DB *gorm.DB

func Connect() {
	var err error
	for i := 0; i < 5; i++ {
		DB, err = gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/axiomQ?parseTime=true"), &gorm.Config{})
		if err != nil && i < 5 {
			time.Sleep(2 * time.Second)
			log.Println("Reconnection to database...")
			continue
		}
		break
	}
}

func AutoMigration() {
	err := DB.Migrator().DropTable(model.Book{}, model.User{})
	if err != nil {
		panic("Error while drop tables!")
	}
	err = DB.AutoMigrate(model.Book{}, model.User{})
	if err != nil {
		panic("Error while creating tables!")
	}
	DB.Create(Books)
	DB.Create(Users)
}
