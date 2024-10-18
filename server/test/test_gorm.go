package main

import (
	"gincrm/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	uri := "root:123456@tcp(127.0.0.1:3307)/gincrm?charset=utf8&parseTime=True&loc=Local&multiStatements=true"

	db, err := gorm.Open(mysql.Open(uri), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	//_ = db.AutoMigrate(&models.Message{})
	//_ = db.AutoMigrate(&models.User{})
	//_ = db.AutoMigrate(&models.Notice{})
	//_ = db.AutoMigrate(&models.Subscribe{})
	//_ = db.AutoMigrate(&models.Customer{})
	//_ = db.AutoMigrate(&models.MailConfig{})
	//_ = db.AutoMigrate(&models.Contract{})
	_ = db.AutoMigrate(&models.Product{})
	//_ = db.AutoMigrate(&models.Contact{})
	//_ = db.AutoMigrate(&models.Group{})
	//_ = db.AutoMigrate(&models.Community{})
	//
	//// Create
	//db.Create(&models.User{Name: "abc", Password: "555"})
	//
	//// Read
	//var user models.User
	//db.First(&user, "name = ?", "abc")
	//fmt.Println(user)
	//// Update
	//db.Model(&user).Update("Password", "321")
	//// Update - update multiple fields
	//db.Model(&user).Updates(models.User{Name: "ydc"})
	//db.Model(&user).Updates(map[string]interface{}{"Name": "aaa", "Password": "F42"})
	//
	//// Delete - delete user
	//db.Delete(&user, "name = ?", "aaa")
}
