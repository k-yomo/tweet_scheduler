package handler

import (
	"github.com/k-yomo/tweet_scheduler/helper"
	"github.com/k-yomo/tweet_scheduler/models"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func (h *Handler) Signup(c echo.Context) (err error) {
	u := &models.User{}
	if err = c.Bind(u); err != nil {
		return
	}

	// Validate
	if u.Email == "" || u.Password == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid email or password"}
	}

	// hashing
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: err}
	}
	u.PasswordHash = passwordHash

	// Save user
	result := h.DB.Create(u)
	if result.Error != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: result.Error}
	}
	u.Token, err = jwt_generator.GenerateJwt(u.ID)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: err}
	}

	// Not to include password in the response
	u.Password = ""
	return c.JSON(http.StatusCreated, u)
}

func (h *Handler) Login(c echo.Context) (err error) {
	// Bind
	u := new(models.User)
	if err = c.Bind(u); err != nil {
		return
	}

	result := h.DB.First(&u, "email = ?", u.Email)
	if result.Error != nil {
		if result.RecordNotFound() {
			return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid email or password"}
		}
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: result.Error}
	}

	if err = bcrypt.CompareHashAndPassword(u.PasswordHash, []byte(u.Password)); err != nil {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: err}
	}
	
	u.Token, err = jwt_generator.GenerateJwt(u.ID)
	if err != nil {
		return err
	}

	// Not to include password in the response
	u.Password = ""
	return c.JSON(http.StatusOK, u)
}
