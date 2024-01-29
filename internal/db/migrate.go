package db

import (
	"github.com/LamichhaneBibek/go-template/internal/api/models"
	"github.com/LamichhaneBibek/go-template/internal/config"
	"github.com/LamichhaneBibek/go-template/internal/constants"
	"github.com/LamichhaneBibek/go-template/internal/utils/logger"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var loge = logger.NewLogger(config.GetConfig())

func Up() {
	database := GetDB()
	createTables(database)
	createDefaultUserInformation(database)
	loge.Info(logger.Postgres, logger.Migration, "MigrationUp", nil)

}

func addNewTable(db *gorm.DB, model interface{}, tables []interface{}) []interface{} {
	if !db.Migrator().HasTable(model) {
		tables = append(tables, model)
	}
	return tables
}

func createTables(db *gorm.DB) {
	tables := []interface{}{}

	tables = addNewTable(db, models.User{}, tables)

	err := db.AutoMigrate(tables...)
	if err != nil {
		loge.Error(logger.Postgres, logger.Migration, err.Error(), nil)
	}
	loge.Info(logger.Postgres, logger.Migration, "tables created", nil)
}

func createDefaultUserInformation(db *gorm.DB) {
	user := models.User{Username: constants.DefaultUserName, Email: constants.DefaultUserEmail}
	pass := constants.DefaultUserPassword
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	createAdminUserIfNotExists(db, &user)
}

func createAdminUserIfNotExists(db *gorm.DB, user *models.User) {
	var exists int64 = 0
	db.Model(&models.User{}).
		Select("1").
		Where("username = ?", user.Username).
		First(&exists)
	if exists == 0 {
		db.Create(user)
	}
}
