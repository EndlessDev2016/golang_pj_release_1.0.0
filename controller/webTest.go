package controller

import (
	"go-gin-pj/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// WebTest - Web Test API
func WebTest(c *gin.Context) {
	// c.JSON(http.StatusOK, gin.H{"message": "pong",})
	// if err != nil {
	// 	// ...
	// }
	testService := service.WebTestService{}
	test := testService.GetList()
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    test,
	})
}
