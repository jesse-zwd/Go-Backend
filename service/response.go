package service

import (
	"github.com/gin-gonic/gin"
)

// Response basic serializer
type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Msg   string      `json:"msg"`
	Error string      `json:"error,omitempty"`
}

// TrackedErrorResponse
type TrackedErrorResponse struct {
	Response
	TrackID string `json:"track_id"`
}

// Error code
const (
	// CodeCheckLogin
	CodeCheckLogin = 401
	// CodeNoRightErr
	CodeNoRightErr = 403
	// CodeDBError database error
	CodeDBError = 50001
	// CodeEncryptError encrypting failed
	CodeEncryptError = 50002
	//CodeParamErr other param error
	CodeParamErr = 40001
)

// CheckLogin
func CheckLogin() Response {
	return Response{
		Code: CodeCheckLogin,
		Msg:  "not login",
	}
}

// Err handling common error
func Err(errCode int, msg string, err error) Response {
	res := Response{
		Code: errCode,
		Msg:  msg,
	}

	if err != nil && gin.Mode() != gin.ReleaseMode {
		res.Error = err.Error()
	}
	return res
}

// DBErr database error
func DBErr(msg string, err error) Response {
	if msg == "" {
		msg = "database error"
	}
	return Err(CodeDBError, msg, err)
}

// ParamErr other errors
func ParamErr(msg string, err error) Response {
	if msg == "" {
		msg = "param error"
	}
	return Err(CodeParamErr, msg, err)
}

type User struct {
	ID        uint   `json:"id"`
	UserName  string `json:"username"`
	Nickname  string `json:"nickname"`
	Status    string `json:"status"`
	Avatar    string `json:"avatar"`
	CreatedAt int64  `json:"created_at"`
}


