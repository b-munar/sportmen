package database

import (
	"sportmen/model"
)

func Migrate() {
	ConnectDB()

	db := DB

	db.AutoMigrate(&model.Sportmen{}, &model.SportmenSport{})

}
