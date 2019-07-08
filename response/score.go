package response

import (
	"github.com/agcx_server/config"
	"github.com/agcx_server/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ScoreCreate 分数保存响应
func ScoreCreate(c *gin.Context, scoreID uint) {
	c.JSON(http.StatusCreated, gin.H{
		"status":   http.StatusCreated,
		"score_id": scoreID,
	})
}

// ScoreList 分数列表响应
func ScoreList(c *gin.Context, total, page int, data []models.Score) {
	c.JSON(http.StatusOK, gin.H{
		"status":       http.StatusOK,
		"total":        total,
		"current_page": page,
		"per_page":     config.App.ItemsPerPage,
		"data":         data,
	})
}
