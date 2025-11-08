package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

type UserAuthRequest struct {
	User     string `json:"user" binding:"required,min=3,max=32"`
	Password string `json:"password" binding:"required,min=6,max=64"`
}

func LoginHandler(c *gin.Context) {
	var req UserAuthRequest

	// 解析 JSON 请求体
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "密码错误",
			"details": err.Error(),
		})
		return
	}

	// TODO: 实际应用中请替换为数据库查询或验证逻辑
	if req.User == "admin" && req.Password == "123456" {
		token, error := GenerateToken(req.User)
		if error != nil {
			c.JSON(200, gin.H{
				"error": "生成token失败",
			})
			return
		}
		c.SetCookie("whoami", token, 60*60*12, "", "", true, true)
		c.JSON(http.StatusOK, gin.H{
			"message": "登陆成功",
			"user":    req.User,
			"token":   token,
		})
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{
		"error": "密码错误",
	})
}

func RegisterHandler(c *gin.Context) {
	var req UserAuthRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request parameters",
			"details": err.Error(),
		})
		return
	}

	// TODO: 在这里实现数据库插入逻辑
	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user":    req.User,
	})
}
