package controllers

import (
	"github.com/xuxuxugames/agcx-server/config"
	"github.com/xuxuxugames/agcx-server/models"
	"github.com/xuxuxugames/agcx-server/requests"
	"github.com/xuxuxugames/agcx-server/response"
	"github.com/xuxuxugames/agcx-server/utils/common"
	"github.com/xuxuxugames/agcx-server/utils/database"
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
		IP:     c.ClientIP(),
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

// ScoreList
func ScoreList(c *gin.Context) {
	// 初始化条件查询模型
	scores := []models.Score{}
	db := database.Connector

	// 检测游戏
	if game, isExist := c.GetQuery("game"); isExist {
		if game != "2048" && game != "snake" && game != "tetris" {
			response.BadRequest(c, "游戏筛选请求错误")
			c.Abort()
			return
		} else {
			db = db.Where("game = ?", game)
		}
	}

	// 检测开始时间
	if startAt, isExist := c.GetQuery("start_at"); isExist {
		startAt, err := common.ParseTime(startAt)
		if err != nil {
			response.BadRequest(c, "开始时间格式错误")
			c.Abort()
			return
		}
		db = db.Where("created_at >= ?", startAt)
	}

	// 检测结束时间
	if endAt, isExist := c.GetQuery("end_at"); isExist {
		endAt, err := common.ParseTime(endAt)
		if err != nil {
			response.BadRequest(c, "结束时间格式错误")
			c.Abort()
			return
		}
		db = db.Where("created_at <= ?", endAt)
	}

	// 检测分页
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}
	if page < 1 {
		page = 1
	}
	perPage := config.App.ItemsPerPage
	total := 0

	// 执行查询
	db.Preload("User").Limit(perPage).Offset((page - 1) * perPage).Order("score desc").Find(&scores)
	db.Model(&models.Score{}).Count(&total)

	// 判断当前页是否为空
	if (page-1)*perPage >= total {
		response.NoContent(c)
		c.Abort()
		return
	}

	// 发送响应
	response.ScoreList(c, total, page, scores)
}
