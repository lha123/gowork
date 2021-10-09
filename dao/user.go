package dao

import (
	"fmt"
	"github.com/lha123/gowork/public"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

var mysqlDb *gorm.DB

func init() {
	fmt.Println("===================================")
	dsn := "comma-admin:comma-admin@tcp(139.224.1.155:3306)/todo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	if err != nil {
		log.Fatalf("db open is error:%v", err)
		return
	}
	mysqlDb = db
}

type UserInfo struct {
	ID         int    `gorm:"primaryKey;autoIncrement;column:id"`
	Name       string `gorm:"column:name"`
	Age        int    `gorm:"column:age"`
	Email      string `gorm:"column:email"`
	CreateTime time.Time
	UpdateTime time.Time
	Del        int
}

func AddUser(u *UserInfo) bool {
	return mysqlDb.Create(u).RowsAffected > 0
}

func UpdateUser(u *UserInfo) bool {
	return mysqlDb.Model(u).Updates(UserInfo{Name: u.Name, Age: u.Age, Email: u.Email}).RowsAffected > 0
}

func Del(id int) bool {
	return mysqlDb.Unscoped().Delete(&UserInfo{}, id).RowsAffected > 0
}

func QueryUserInfoAll() []UserInfo {
	var a []UserInfo
	mysqlDb.Find(&a)
	return a
}

func FindUserInfoById(id int) *UserInfo {
	a := new(UserInfo)
	mysqlDb.Find(a, id)
	return a
}

func FindUserInfoByIdIn(ids []int) []UserInfo {
	var a []UserInfo
	mysqlDb.Find(&a, ids)
	return a
}

func FindUserInfoByCondition(u *UserInfo) []UserInfo {
	var r []UserInfo
	mysqlDb.Where(u).Find(&r)
	return r
}

func FinUserInfoPage(p *public.Page) []UserInfo {
	var r []UserInfo
	mysqlDb.Offset((p.Page - 1) * p.Size).Limit(p.Size).Find(&r)
	return r
}
