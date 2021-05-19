package initialize

import (
	"backend/model"
	"backend/global"
)
//database migration

func migration() {
	global.GORM_DB.AutoMigrate(&model.User{})
}
