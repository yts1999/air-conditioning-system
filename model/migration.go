package model

//执行数据迁移
func migration() {
	// 自动迁移模式
	DB.AutoMigrate(&Admin{})
	DB.AutoMigrate(&Room{})
	DB.AutoMigrate(&Guest{})
	DB.AutoMigrate(&Record{})
	DB.AutoMigrate(&Switch{})
	// 外键约束
	DB.Model(&Guest{}).AddForeignKey("room_id", "rooms(room_id)", "RESTRICT", "RESTRICT")
	DB.Model(&Record{}).AddForeignKey("room_id", "rooms(room_id)", "CASCADE", "CASCADE")
	DB.Model(&Room{}).AddForeignKey("current_record", "records(id)", "SET NULL", "RESTRICT")
	DB.Model(&Switch{}).AddForeignKey("room_id", "rooms(room_id)", "CASCADE", "CASCADE")
}
