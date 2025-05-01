package controller

import (
	"net/http"

	"go-middleware-primary/utils"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var json struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&json); err != nil || json.Username != "admin" || json.Password != "123456" {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "用户名或密码错误"})
		return
	}

	token, _ := utils.GenerateToken(json.Username)
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func Profile(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{"user": user, "msg": "欢迎访问个人中心"})
}

func GetData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "这是受保护的数据"})
}
