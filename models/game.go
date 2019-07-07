package models

// Game 游戏模型
type Game struct {
	CommonFields
	Name string `json:"name"`
}
