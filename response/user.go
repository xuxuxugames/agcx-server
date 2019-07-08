package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserAuth 用户认证响应
func UserAuth(c *gin.Context, userID uint, token string) {
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"user_id": userID,
		"token":   token,
	})
}

// UserCreate 用户注册响应
func UserCreate(c *gin.Context, userID uint) {
	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"user_id": userID,
	})
}

// UserPassword 修改密码响应
func UserPassword(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
	})
}
