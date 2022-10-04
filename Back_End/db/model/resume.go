package model

import (
	"Back_End/db"
	"fmt"

	"gorm.io/gorm/clause"
)

type ResumeInfo struct {
	UserID uint64 `json:"-"`
	Name   string `gorm:"not null" json:"name2" validate:"required"`

	TableID uint64 `gorm:"primaryKey" json:"TableId2"`

	PhoneNumber       string `json:"PhoneNumber2" validate:"required" gorm:"not null"`
	Web               string `json:"web2" validate:"required" gorm:"not null"`
	Mail              string `json:"mail2" validate:"required" gorm:"not null"`
	Workplace         string `json:"workplace2" validate:"required" gorm:"not null"`
	Position          string `json:"position2" validate:"required" gorm:"not null"`
	SelfIntroduce     string `json:"self-introduce2" validate:"required" gorm:"not null"`
	Education         string `json:"education2" validate:"required" gorm:"not null"`
	Skill             string `json:"skill2" validate:"required" gorm:"not null"`
	WorkExprience     string `json:"WorkExprience2" validate:"required" gorm:"not null"`
	ProjectExperience string `json:"ProjectExperience2" validate:"required" gorm:"not null"`
	MoreInformation   string `json:"MoreInformation2" validate:"required" gorm:"not null"`

	// ClassID uint64 `gorm:"primaryKey" json:"class_id"`
	// Name    string `gorm:"not null" json:"name" validate:"required"`
	// Teacher string `json:"teacher"`
	// Week    string `json:"week" validate:"required" gorm:"not null"`
	// Time    string `json:"time" validate:"required" gorm:"not null"`
	// WeekDay string `json:"week_day" validate:"required" gorm:"not null"`
	// Address string `json:"address" validate:"required" gorm:"not null"`
}

func AddResume(resume *ResumeInfo) error {
	fmt.Println(resume)
	return db.DB.Create(resume).Error
}

func DeleteResumeByID(ID uint64) error {
	c, err := GetResumeByID(ID)
	if err != nil {
		return err
	}
	return db.DB.Select(clause.Associations).Delete(c).Error
}

func UpdateResume(resume *ResumeInfo) error {
	var c ResumeInfo
	err := db.DB.Where("TableId2 = ?", resume.TableID).Take(&c).Error
	if err != nil {
		return err
	}
	resume.TableID = c.TableID
	return db.DB.Omit("TableID", "UserID").Updates(resume).Error
}

func GetResumeByID(ID uint64) (*ResumeInfo, error) {
	var resume ResumeInfo
	err := db.DB.Where("TableId2 = ?", ID).First(&resume).Error
	return &resume, err
}

func GetResumesByUserID(userID uint64) ([]ResumeInfo, error) {
	var resumes []ResumeInfo
	err := db.DB.Where("user_id = ?", userID).Find(&resumes).Error
	return resumes, err
}
