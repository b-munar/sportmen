package database

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() {
	var err error

	viper.SetConfigFile(".env")

	viper.ReadInConfig()

	viper.AutomaticEnv()

	db_user, _ := viper.Get("SPORTMEN_DB_USER").(string)
	db_host, _ := viper.Get("SPORTMEN_DB_HOST").(string)
	db_name, _ := viper.Get("SPORTMEN_DB_NAME").(string)
	db_password, _ := viper.Get("SPORTMEN_DB_PASSWORD").(string)
	db_port, _ := viper.Get("SPORTMEN_DB_PORT").(string)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		db_host, db_user, db_password, db_name, db_port)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
}
