package controller

import (
	"backend/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserRegister godoc
// @Summary user register
// @Description user register
// @Tags user API
// @ID /user/register
// @Accept  json
// @Produce  json
// @Param body body service.UserRegisterService true "body"
// @Success 200 {object} service.Response{data=service.User} "success"
// @Router /user/register [post]
func UserRegister(c *gin.Context) {
	var service service.UserRegisterService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Register()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserLogin godoc
// @Summary user login
// @Description user login
// @Tags user API
// @ID /user/login
// @Accept  json
// @Produce  json
// @Param body body service.UserLoginService true "body"
// @Success 200 {object} service.Response{data=service.User} "success"
// @Router /user/login [post]
func UserLogin(c *gin.Context) {
	var service service.UserLoginService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Login(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserInfo godoc
// @Summary user info
// @Description user info
// @Tags user API
// @ID /user/me
// @Accept  json
// @Produce  json
// @Success 200 {object} service.Response{data=service.User} "success"
// @Router /user/me [get]
func UserMe(c *gin.Context) {
	user := service.CurrentUser(c)
	res := service.BuildUserResponse(*user)
	c.JSON(200, res)
}

// UserLogout godoc
// @Summary user logout
// @Description user logout
// @Tags user API
// @ID /user/logout
// @Accept  json
// @Produce  json
// @Success 200 {object} service.Response{} "success"
// @Router /user/logout [delete]
func UserLogout(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	s.Save()
	c.JSON(200, service.Response{
		Code: 0,
		Msg:  "Logged out",
	})
}
