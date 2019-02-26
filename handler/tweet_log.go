package handler

import (
	"github.com/k-yomo/tweet_scheduler/models"
	"github.com/labstack/echo"
	"net/http"
)

func (h *Handler) GetTweetLogs(c echo.Context) (err error) {
	tweetLogs := make([]models.TweetLog, 0)
	result := h.DB.Order("created_at desc").Find(&tweetLogs)
	if result.Error != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: result.Error}
	}

	return c.JSON(http.StatusOK, map[string][]models.TweetLog{
		"tweetLogs": tweetLogs,
	})
}
