package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/szdx4/attsys-server/config"
	"github.com/szdx4/attsys-server/requests"
	"github.com/szdx4/attsys-server/response"
	"time"
)

// UserAuth 用户认证
func UserAuth(c *gin.Context) {
	var req requests.UserAuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Bad Request")
		c.Abort()
		return
	}

	// 验证提交数据的合法性
	user, err := req.Validate()
	if err != nil {
		response.Unauthorized(c, err.Error())
		c.Abort()
		return
	}

	// 生成新的 JWT Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
		//"role":       user.Role,
		"expired_at": time.Now().Add(time.Hour * time.Duration(config.App.TokenValid)).Format(time.RFC3339),
	})
	tokenString, err := token.SignedString([]byte(config.App.EncryptKey))
	if err != nil {
		response.InternalServerError(c, "Token sign error")
		c.Abort()
		return
	}

	// 发送响应
	response.UserAuth(c, user.ID, tokenString)
}
