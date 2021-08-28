package mysql

import (
	"assessment/Models"
	"assessment/settings"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// Init 初始化MySQL连接
func Init(cfg *settings.MySQLConfig) (err error) {
	// "user:password@tcp(host:port)/dbname"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	db, err = gorm.Open("mysql", dsn)
	db.AutoMigrate(&Models.Post{}, &Models.Community{}, &Models.User{}, &Models.Comment{})
	if err != nil {
		return
	}
	return
}

// Close 关闭MySQL连接
func Close() {
	_ = db.Close()
}
