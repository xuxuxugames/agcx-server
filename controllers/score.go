package controllers

import (
	"github.com/agcx_server/models"
	"github.com/agcx_server/requests"
	"github.com/agcx_server/response"
	"github.com/agcx_server/utils/database"
	"github.com/gin-gonic/gin"
	"strconv"
)

// ScoreCreate 保存分数
func ScoreCreate(c *gin.Context) {
	var req requests.ScoreCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "错误请求")
		c.Abort()
		return
	}

	// 验证提交数据的合法性
	if err := req.Validate(c); err != nil {
		response.BadRequest(c, err.Error())
		c.Abort()
		return
	}

	// 创建分数
	userID, _ := strconv.Atoi(c.Param("user_id"))
	score := models.Score{
		UserID: uint(userID),
		Game:   req.Game,
		Score:  req.Score,
	}
	database.Connector.Create(&score)
	if score.ID < 1 {
		response.InternalServerError(c, "数据库异常")
		c.Abort()
		return
	}

	// 发送响应
	response.ScoreCreate(c, score.ID)
}
