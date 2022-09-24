package database

import "task-5-vix-btpns-Yoga_Cahya_Romadhon/models"

func MigrateDb() {
	db := ConnDb()
	db.AutoMigrate(&models.User{}, &models.Photo{})
}
