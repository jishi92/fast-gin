package dao

import (
	"fast-gin/config"
	"fast-gin/model"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var db *gorm.DB

var dbc = config.Cfg.Database

func Setup() {
	// 构建 DSL
	DSL := fmt.Sprintf("%s:%s@%s(%s)/%s?charset=%s&parseTime=%s&loc=%s",
		dbc.User, dbc.Password, dbc.Protocol, dbc.Host, dbc.Name, dbc.Charset, dbc.ParseTime, dbc.Loc)

	// 连接到数据库
	var err error
	db, err = gorm.Open(dbc.Dialect, DSL)
	if err != nil {
		log.Fatalf("can't open database err: %v", err)
	}

	// 替换表名 Handler，设置表前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return dbc.TablePrefix + defaultTableName
	}

	// 全局禁用复数表名
	db.SingularTable(true)

	// 最大链接数
	db.DB().SetMaxIdleConns(10)
	// 最大打开链接
	db.DB().SetMaxOpenConns(100)

	// 自动迁移
	db.AutoMigrate(&model.User{})
}

func Close() {
	db.Close()
}
