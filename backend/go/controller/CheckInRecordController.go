package controller

import (
	//需要用到的结构体
	"backend/go/entity"
	//gin框架的依赖
	"github.com/gin-gonic/gin"
	//http连接包
	"net/http"
	//service层方法
	"backend/go/service"
)

func CreateCheckInRecord(c *gin.Context) {
	var checkInRecord entity.CheckInRecord
	c.BindJSON(&checkInRecord)
	err := service.CreateCheckInRecord(&checkInRecord)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
		})
	}
}

func GetCheckInRecordList(c *gin.Context) {
	checkInRecordList, err := service.GetAllCheckInRecord()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": checkInRecordList,
		})
	}
}

func DeleteCheckInRecord(c *gin.Context) {
	var checkInRecord entity.CheckInRecord
	c.BindJSON(&checkInRecord)
	err := service.DeleteStudentById(checkInRecord.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
		})
	}
}
