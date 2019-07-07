package models

// User 用户模型
type User struct {
	CommonFields
	Name     string `json:"name"` // 用户名
	Password string `json:"-"`    // 用户密码
}
