package config

import (
	"fmt"

	"github.com/satti999/todoapp/src/model"
	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
	SSLMode  string
}

func NewConnection(config *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}
	fmt.Println("database connected", db)
	return db, nil
}

func MigrateModels(db *gorm.DB) error {
	err := db.AutoMigrate(&model.User{}, &model.Todo{}, &model.SubTodo{})
	fmt.Println("All the tables are migrated")
	return err
}
