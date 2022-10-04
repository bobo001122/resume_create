// 处理前端请求
package controller

import (
	"Back_End/db/model"
	"Back_End/utility"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddResume(c *gin.Context) {
	user_id, _ := c.Get("user_id")
	var data model.ResumeInfo
	err := c.ShouldBindJSON(&data)
	if err != nil {
		log.Println(err)
		utility.ResponseBadRequest(c)
		return
	}
	data.UserID = user_id.(uint64)
	err = model.AddResume(&data)
	if err != nil {
		log.Println(err)
		utility.ResponseInternalServerError(c)
		return
	}
	utility.ResponseOK(c, nil)
}

func DeleteResume(c *gin.Context) {
	resumeID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		utility.ResponseBadRequest(c)
	}
	err = model.DeleteResumeByID(uint64(resumeID))
	if err == gorm.ErrRecordNotFound {
		utility.Response(http.StatusNotFound, "Resume not found", nil, c)
		return
	} else if err != nil {
		utility.ResponseInternalServerError(c)
		return
	}
	utility.ResponseOK(c, nil)
}

func UpdateResume(c *gin.Context) {
	var resume model.ResumeInfo
	err := c.ShouldBindJSON(&resume)
	fmt.Println(resume)
	if err != nil {
		utility.ResponseBadRequest(c)
		return
	}
	err = model.UpdateResume(&resume)
	if err == gorm.ErrRecordNotFound {
		utility.Response(http.StatusNotFound, "Resume not found", nil, c)
		return
	} else if err != nil {
		utility.ResponseInternalServerError(c)
		return
	}
	utility.ResponseOK(c, nil)
}

func GetResumes(c *gin.Context) {
	user_id, _ := c.Get("user_id")
	resumes, err := model.GetResumesByUserID(user_id.(uint64))
	if err != nil {
		utility.ResponseInternalServerError(c)
		return
	}
	utility.ResponseOK(c, gin.H{
		"resume": resumes,
	})
}
