package model

import (
	"Back_End/db"
	"log"
)

func InitModel() {
	err := db.DB.AutoMigrate(&User{}, &ResumeInfo{})
	if err != nil {
		log.Panicln("Database Error: ", err)
	}
}
