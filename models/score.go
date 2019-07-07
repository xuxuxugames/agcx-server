package models

// Score 分数模型
type Score struct {
	CommonFields
	UserID uint `json:"user_id"`
	GameID uint `json:"game_id"`
	Score  int  `json:"score"`
}
