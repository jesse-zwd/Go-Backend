package controller

import (
	"backend/initialize"
	"backend/model"
	"backend/service"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	validator "gopkg.in/go-playground/validator.v8"
)

// CurrentUser get current user
func CurrentUser(c *gin.Context) *model.User {
	if user, _ := c.Get("user"); user != nil {
		if u, ok := user.(*model.User); ok {
			return u
		}
	}
	return nil
}

// ErrorResponse get error
func ErrorResponse(err error) service.Response {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			field := initialize.T(fmt.Sprintf("Field.%s", e.Field))
			tag := initialize.T(fmt.Sprintf("Tag.Valid.%s", e.Tag))
			return service.ParamErr(
				fmt.Sprintf("%s%s", field, tag),
				err,
			)
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return service.ParamErr("JSON types don't match", err)
	}

	return service.ParamErr("param error", err)
}
