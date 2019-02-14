package handler

import (
	"github.com/k-yomo/tweet_scheduler/models"
	"github.com/labstack/echo"
	"net/http"
)


func (h *Handler) CreateTweet(c echo.Context) (err error) {
	t := &models.Tweet{}
	if err = c.Bind(t); err != nil {
		return
	}

	result := h.DB.Create(t)
	if result.Error != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: result.Error}
	}

	return c.JSON(http.StatusOK, t)
}

func (h *Handler) GetTweets(c echo.Context) (err error) {
	tweets := make([]models.Tweet, 0)
	result := h.DB.Find(&tweets)
	if result.Error != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: result.Error}
	}

	return c.JSON(http.StatusOK, map[string][]models.Tweet{
		"tweets": tweets,
	})
}

func (h *Handler) UpdateTweet(c echo.Context) (err error) {
	id := c.Param("id")
	t := &models.Tweet{}
	h.DB.First(t, "id = ?", id)

	if err = c.Bind(t); err != nil {
		return
	}
	result := h.DB.Save(&t)
	if result.Error != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: result.Error}
	}

	return c.JSON(http.StatusOK, t)
}

func (h *Handler) DeleteTweet(c echo.Context) (err error) {
	id := c.Param("id")
	result := h.DB.Delete(&models.Tweet{}, "id = ?", id)
	if result.Error != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: result.Error}
	}

	return c.JSON(http.StatusOK, "")
}
