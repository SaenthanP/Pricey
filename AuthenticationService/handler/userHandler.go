package handler

import(
	"github.com/gin-gonic/gin"
	"net/http"
	)

func RegisterUser(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message": "sucess",
	})
}

