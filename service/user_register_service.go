package service

import (
	"backend/model"
	"backend/serializer"
)

// UserRegisterService register service
type UserRegisterService struct {
	Nickname        string `form:"nickname" json:"nickname" binding:"required,min=2,max=30"`
	UserName        string `form:"username" json:"username" binding:"required,min=5,max=30"`
	Password        string `form:"password" json:"password" binding:"required,min=8,max=40"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=8,max=40"`
}

// valid form
func (service *UserRegisterService) valid() *serializer.Response {
	if service.PasswordConfirm != service.Password {
		return &serializer.Response{
			Code: 40001,
			Msg:  "passwords don't match",
		}
	}

	count := 0
	model.DB.Model(&model.User{}).Where("nickname = ?", service.Nickname).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Code: 40001,
			Msg:  "nickname not available",
		}
	}

	count = 0
	model.DB.Model(&model.User{}).Where("user_name = ?", service.UserName).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Code: 40001,
			Msg:  "username not available",
		}
	}

	return nil
}

// Register 
func (service *UserRegisterService) Register() serializer.Response {
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
		return serializer.Err(
			serializer.CodeEncryptError,
			"encryption failed",
			err,
		)
	}

	// Create user
	if err := model.DB.Create(&user).Error; err != nil {
		return serializer.ParamErr("register failed", err)
	}

	return serializer.BuildUserResponse(user)
}
