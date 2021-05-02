package service

import (
	"backend/model"
	"backend/serializer"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserLoginService login service
type UserLoginService struct {
	UserName string `form:"username" json:"username" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
}

// setSession set session
func (service *UserLoginService) setSession(c *gin.Context, user model.User) {
	s := sessions.Default(c)
	s.Clear()
	s.Set("user_id", user.ID)
	s.Save()
}

// Login 
func (service *UserLoginService) Login(c *gin.Context) serializer.Response {
	var user model.User

	if err := model.DB.Where("user_name = ?", service.UserName).First(&user).Error; err != nil {
		return serializer.ParamErr("wrong username or password", nil)
	}

	if user.CheckPassword(service.Password) == false {
		return serializer.ParamErr("wrong username or password", nil)
	}

	// set session
	service.setSession(c, user)

	return serializer.BuildUserResponse(user)
}
