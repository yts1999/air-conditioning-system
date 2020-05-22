package serializer

import "centralac/model"

// Admin 管理员序列化器
type Admin struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	CreatedAt int64  `json:"created_at"`
}

// BuildAdmin 序列化管理员
func BuildAdmin(admin model.Admin) Admin {
	return Admin{
		ID:        admin.ID,
		Username:  admin.Username,
		CreatedAt: admin.CreatedAt.Unix(),
	}
}

// BuildAdminResponse 序列化管理员响应
func BuildAdminResponse(admin model.Admin) Response {
	return Response{
		Data: BuildAdmin(admin),
	}
}
