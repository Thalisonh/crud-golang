package migrations

import (
	"github.com/Thalisonh/crud-golang/database/entity"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(entity.Book{})
	db.AutoMigrate(entity.User{})
}
