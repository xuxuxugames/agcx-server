package models

// User 用户模型
type User struct {
	CommonFields
	Email    string `json:"email"`    // 登陆邮箱
	Password string `json:"password"` // 用户密码
	Name     string `json:"name"`     // 用户名
}
