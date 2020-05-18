package serializer

import "centralac/model"

// User 用户序列化器
type User struct {
	ID        uint   `json:"id"`
	RoomID    string `json:"room_id"`
	CreatedAt int64  `json:"created_at"`
}

// BuildUser 序列化用户
func BuildUser(user model.User) User {
	return User{
		ID:        user.ID,
		RoomID:    user.RoomID,
		CreatedAt: user.CreatedAt.Unix(),
	}
}

// BuildUserResponse 序列化用户响应
func BuildUserResponse(user model.User) Response {
	return Response{
		Data: BuildUser(user),
	}
}
