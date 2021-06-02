package service

import (
	"backend/global"
	"backend/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// setSession set session
func (service *UserLoginService) setSession(c *gin.Context, user model.User) {
	s := sessions.Default(c)
	s.Clear()
	s.Set("user_id", user.ID)
	s.Save()
}

// Login
func (service *UserLoginService) Login(c *gin.Context) Response {
	var user model.User

	if err := global.GORM_DB.Where("user_name = ?", service.UserName).First(&user).Error; err != nil {
		return ParamErr("wrong username or password", nil)
	}

	if !user.CheckPassword(service.Password) {
		return ParamErr("wrong username or password", nil)
	}

	// set session
	service.setSession(c, user)

	return BuildUserResponse(user)
}

// valid form
func (service *UserRegisterService) valid() *Response {
	if service.PasswordConfirm != service.Password {
		return &Response{
			Code: 40001,
			Msg:  "passwords don't match",
		}
	}

	var count int64
	global.GORM_DB.Model(&model.User{}).Where("nickname = ?", service.Nickname).Count(&count)
	if count > 0 {
		return &Response{
			Code: 40001,
			Msg:  "nickname not available",
		}
	}

	count = 0
	global.GORM_DB.Model(&model.User{}).Where("user_name = ?", service.UserName).Count(&count)
	if count > 0 {
		return &Response{
			Code: 40001,
			Msg:  "username not available",
		}
	}

	return nil
}

// Register
func (service *UserRegisterService) Register() Response {
	user := model.User{
		Nickname: service.Nickname,
		UserName: service.UserName,
		Status:   model.Active,
	}

	// validate
	if err := service.valid(); err != nil {
		return *err
	}

	// encrypt
	if err := user.SetPassword(service.Password); err != nil {
		return Err(
			CodeEncryptError,
			"encryption failed",
			err,
		)
	}

	// Create user
	if err := global.GORM_DB.Create(&user).Error; err != nil {
		return ParamErr("register failed", err)
	}

	return BuildUserResponse(user)
}

// BuildUser 
func BuildUser(user model.User) User {
	return User{
		ID:        user.ID,
		UserName:  user.UserName,
		Nickname:  user.Nickname,
		Status:    user.Status,
		Avatar:    user.Avatar,
		CreatedAt: user.CreatedAt.Unix(),
	}
}

// BuildUserResponse Response of User
func BuildUserResponse(user model.User) Response {
	return Response{
		Data: BuildUser(user),
	}
}

// CurrentUser get current user
func CurrentUser(c *gin.Context) *model.User {
	if user, _ := c.Get("user"); user != nil {
		if u, ok := user.(*model.User); ok {
			return u
		}
	}
	return nil
}
