package controller

import (
	"backend/service"
	"encoding/json"
)

// ErrorResponse get error
func ErrorResponse(err error) service.Response {
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return service.ParamErr("JSON types don't match", err)
	}

	return service.ParamErr("param error", err)
}
