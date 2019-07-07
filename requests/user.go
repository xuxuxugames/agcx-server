package requests

import (
	"errors"
	"github.com/agcx_server/models"
	"github.com/agcx_server/utils/database"
	"golang.org/x/crypto/bcrypt"
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
		return nil, errors.New("User not found")
	}

	// 验证密码是否匹配
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(r.Password))
	if err != nil {
		return nil, errors.New("Password invalid")
	}

	// 无误则返回空
	return user, nil
}
