package serializer

import "centralac/model"

// Admin 管理员序列化器
type Admin struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	CreatedAt int64  `json:"created_at"`
}

// BuildUser 序列化管理员
func BuildUser(admin model.Admin) Admin {
	return Admin{
		ID:        admin.ID,
		Username:  admin.Username,
		CreatedAt: admin.CreatedAt.Unix(),
	}
}

// BuildUserResponse 序列化管理员响应
func BuildUserResponse(admin model.Admin) Response {
	return Response{
		Data: BuildUser(admin),
	}
}
