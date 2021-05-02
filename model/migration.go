package model

//database migration

func migration() {
	DB.AutoMigrate(&User{})
}
