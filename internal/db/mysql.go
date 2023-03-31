package db

import (
	"fmt"
	"github.com/Nicrii/bookstore/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// InitGormMysql initialise new gorm mysql connection
func InitGormMysql() (*gorm.DB, error) {
	conf := config.GetConfig()

	conn, err := gorm.Open(
		mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			conf.User,
			conf.Password,
			conf.Host,
			conf.Port,
			conf.Name,
		),
		), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error in db connection: %w", err)
	}

	return conn, err
}
