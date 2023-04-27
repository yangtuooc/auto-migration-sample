package common

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	v  *viper.Viper
	db *gorm.DB
)

func UseViper() *viper.Viper {
	return v
}

func setupViper(vpr *viper.Viper) {
	v = vpr
}

func UseDB() *gorm.DB {
	return db
}

func setupDB(dbConn *gorm.DB) {
	db = dbConn
}
