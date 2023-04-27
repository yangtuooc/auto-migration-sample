package common

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Gorm() {
	m := Cfg.Mysql
	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Host + ")/" + m.Database + "?" + m.Option
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(fmt.Sprintf("MySQL启动异常: %v", err))
	}
	setupDB(db)
}
