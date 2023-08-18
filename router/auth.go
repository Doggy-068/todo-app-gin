package router

import (
	"net/http"
	"todo-app/constant"
	"todo-app/dto"

	"github.com/golang-jwt/jwt/v5"

	"github.com/gin-gonic/gin"
)

// @summary login by user and password
// @tags auth
// @accept x-www-form-urlencoded
// @param payload formData dto.AuthLoginPostDto true "payload"
// @produce json
// @success 200 {object} dto.AuthLoginPostReturnDto
// @router /api/auth/login [post]
func postLogin(c *gin.Context) {
	form := dto.AuthLoginPostDto{}
	if err := c.ShouldBind(&form); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	if form.User == "admin" && form.Password == "123456" {
		s, err := jwt.New(jwt.SigningMethodHS256).SignedString([]byte(constant.JwtKey))
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
		authLoginPostReturnDto := dto.AuthLoginPostReturnDto{ApiKey: s}
		c.JSON(http.StatusOK, authLoginPostReturnDto)
		return
	}
	c.Status(http.StatusNotFound)
}

func LoadAuthRouter(e *gin.Engine) {
	router := e.Group("/api/auth")
	{
		router.POST("/login", postLogin)
	}
}
