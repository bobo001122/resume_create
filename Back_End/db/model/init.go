package model

import (
	"Back_End/db"
	"log"
)

func InitModel() {
	ok := db.DB.AutoMigrate(&User{}, &Class{})
	if ok != nil {
		log.Panicln("Database Error: ", ok)
	}
}
