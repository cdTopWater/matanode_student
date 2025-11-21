package mapper

import (
	"book-manager/internal/model"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	fmt.Println("mysql init start")
	DB, e := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/book_manager?charset=utf8mb4&parseTime=True&loc=Local"))
	if e != nil {
		fmt.Println(e.Error())
		return nil
	}
	fmt.Println("mysql init doing DB=", DB)
	sqlDb, _ := DB.DB()
	// 设置连接池参数
	sqlDb.SetMaxOpenConns(50)                  // 最大打开连接数
	sqlDb.SetMaxIdleConns(10)                  // 最大空闲连接数
	sqlDb.SetConnMaxLifetime(time.Hour)        // 连接最大存活时间
	sqlDb.SetConnMaxIdleTime(10 * time.Minute) // 连接最大空闲时间

	DB.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{})

	fmt.Println("mysql success")
	return DB
}
