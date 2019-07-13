package requests

import (
	"errors"
	"github.com/xuxuxugames/agcx-server/models"
	"github.com/xuxuxugames/agcx-server/utils/database"
	"github.com/gin-gonic/gin"
	"strconv"
)

// ScoreCreateRequest 分数保存请求
type ScoreCreateRequest struct {
	Game  string `binding:"required"`
	Score uint   `binding:"required"`
}

// Validate 验证 ScoreCreateRequest 请求中用户信息的有效性
func (r *ScoreCreateRequest) Validate(c *gin.Context) error {
	// 验证游戏名称的有效性
	if r.Game != "2048" && r.Game != "snake" && r.Game != "tetris" {
		return errors.New("游戏名称错误")
	}

	// 验证用户的存在性
	userID, _ := strconv.Atoi(c.Param("user_id"))
	user := models.User{}
	database.Connector.Where("id = ?", userID).First(&user)
	if user.ID <= 0 {
		return errors.New("用户不存在")
	}

	// 验证分数的合法性
	if r.Score < 0 {
		return errors.New("分数必须为不小于零的数字")
	}

	// 无误则返回空
	return nil
}
