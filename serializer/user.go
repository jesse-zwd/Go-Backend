package serializer

import "backend/model"

// User serializer
type User struct {
	ID        uint   `json:"id"`
	UserName  string `json:"username"`
	Nickname  string `json:"nickname"`
	Status    string `json:"status"`
	Avatar    string `json:"avatar"`
	CreatedAt int64  `json:"createdAt"`
}

// BuildUser serialize User
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

// BuildUserResponse Response of User serializer
func BuildUserResponse(user model.User) Response {
	return Response{
		Data: BuildUser(user),
	}
}
