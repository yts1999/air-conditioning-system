package model

//执行数据迁移
func migration() {
	// 自动迁移模式
	DB.AutoMigrate(&Admin{})
	DB.AutoMigrate(&Room{})
	DB.AutoMigrate(&Guest{})
	// 外键约束
	DB.Model(&Guest{}).AddForeignKey("room_id", "rooms(room_id)", "RESTRICT", "RESTRICT")
}
