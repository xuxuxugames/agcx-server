package models

// Score 分数模型
type Score struct {
	CommonFields
	UserID uint   `json:"user_id"`
	Game   string `json:"game" gorm:"status:enum('2048', 'pacman', 'snake')"`
	Score  uint   `json:"score"`
	IP     string `json:"ip"`
}
