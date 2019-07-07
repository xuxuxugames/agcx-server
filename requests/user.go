package requests

import (
	"errors"
	"github.com/agcx_server/config"
	"github.com/agcx_server/models"
	"github.com/agcx_server/utils/database"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

// UserAuthRequest 用户认证请求
type UserAuthRequest struct {
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

// Validate 验证 UserAuthRequest 请求中用户信息的有效性
func (r *UserAuthRequest) Validate() (*models.User, error) {
	// 验证用户是否存在
	user := &models.User{}
	database.Connector.Where("email = ?", r.Email).First(&user)
	if user.ID == 0 {
		return nil, errors.New("未找到用户")
	}

	// 验证密码是否匹配
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(r.Password))
	if err != nil {
		return nil, errors.New("密码不正确")
	}

	// 无误则返回空
	return user, nil
}

// UserCreateRequest 用户注册请求
type UserCreateRequest struct {
	Email    string `binding:"required"`
	Password string `binding:"required"`
	Name     string `binding:"required"`
}

// Validate 验证 UserCreateRequest 请求中用户信息的有效性
func (r *UserCreateRequest) Validate() error {
	// 验证用户名长度
	if len(r.Name) < config.App.MinUserLength {
		return errors.New("用户的名字必须长于" + strconv.Itoa(config.App.MinUserLength))
	}

	// 验证登陆邮箱的存在性
	user := models.User{}
	database.Connector.Where("email = ?", r.Email).First(&user)
	if user.ID > 0 {
		return errors.New("用户已经存在")
	}

	// 验证密码长度
	if len(r.Password) < config.App.MinPwdLength {
		return errors.New("密码必须长于" + strconv.Itoa(config.App.MinPwdLength))
	}

	// 无误则返回空
	return nil
}

// UserPasswordRequest 修改密码请求
type UserPasswordRequest struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password" binding:"required"`
}

// Validate 验证修改密码请求的合法性
func (r *UserPasswordRequest) Validate(role string, authID, userID int) (string, error) {
	// 用户必须输入原密码
	if role == "user" && r.OldPassword == "" {
		return "", errors.New("原密码缺失")
	}

	// 用户只能修改自己的密码
	if role == "user" && authID != userID {
		return "", errors.New("您只能修改自己的密码")
	}

	// 验证用户是否存在
	user := models.User{}
	database.Connector.First(&user, userID)
	if user.ID == 0 {
		return "", errors.New("未找到用户")
	}

	// 验证用户的原密码
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(r.OldPassword))
	if err != nil {
		return "", errors.New("密码不正确")
	}

	// 验证新密码的长度
	if len(r.NewPassword) < config.App.MinPwdLength {
		return "", errors.New("密码必须长于" + strconv.Itoa(config.App.MinPwdLength))
	}

	// 生成新密码的 bcrypt hash
	hash, err := bcrypt.GenerateFromPassword([]byte(r.NewPassword), 10)
	if err != nil {
		return "", errors.New("新密码生成错误")
	}

	// 无误则返回空
	return string(hash), nil
}
