package controllers

import (
	"github.com/xuxuxugames/agcx-server/config"
	"github.com/xuxuxugames/agcx-server/models"
	"github.com/xuxuxugames/agcx-server/requests"
	"github.com/xuxuxugames/agcx-server/response"
	"github.com/xuxuxugames/agcx-server/utils/database"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

// UserAuth 用户认证
func UserAuth(c *gin.Context) {
	var req requests.UserAuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "错误请求")
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
		"id":    user.ID,
		"email": user.Email,
		"name":  user.Name,
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

// UserCreate 用户注册
func UserCreate(c *gin.Context) {
	var req requests.UserCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "错误请求")
		c.Abort()
		return
	}

	// 验证提交数据的合法性
	if err := req.Validate(); err != nil {
		response.BadRequest(c, err.Error())
		c.Abort()
		return
	}

	// 生成密码的 bcrypt hash
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		response.InternalServerError(c, "密码生成出现异常")
		c.Abort()
		return
	}

	// 创建用户
	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hash),
	}
	database.Connector.Create(&user)
	if user.ID < 1 {
		response.InternalServerError(c, "数据库异常")
		c.Abort()
		return
	}

	// 发送响应
	response.UserCreate(c, user.ID)
}

// UserPassword 修改密码
func UserPassword(c *gin.Context) {
	// 获取认证用户信息
	authID, exist := c.Get("user_id")
	if !exist {
		response.BadRequest(c, "slkfjdklsf")
		c.Abort()
		return
	}

	// 获取 URL 中用户 ID
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		response.BadRequest(c, "UserID获取错误")
		c.Abort()
		return
	}

	var req requests.UserPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "错误请求")
		c.Abort()
		return
	}

	// 验证提交数据的合法性
	pwdHash, err := req.Validate(authID.(int), int(userID))
	if err != nil {
		response.BadRequest(c, err.Error())
		c.Abort()
		return
	}

	// 执行修改密码操作
	user := models.User{}
	database.Connector.First(&user, userID)
	user.Password = pwdHash
	database.Connector.Save(&user)

	// 发送响应
	response.UserPassword(c)
}
