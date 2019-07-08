package response

import (
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
