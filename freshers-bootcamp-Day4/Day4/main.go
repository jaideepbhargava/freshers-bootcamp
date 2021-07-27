package main

import (
	"fmt"
	"freshers-bootcamp/Day4/Models"
	"freshers-bootcamp/Day4/Routes"
	"gorm.io/driver/mysql"
	"freshers-bootcamp/Day4/db"
	"gorm.io/gorm"
)
var err error
func main() {

	dsn := "root:temppass@tcp(127.0.0.1:3306)/service?charset=utf8mb4&parseTime=True&loc=Local"
	database.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil{
		fmt.Println("Status: ", err)
	}

	database.DB.AutoMigrate(&Models.Product{}, &Models.Customer{},&Models.Orders{})


	apiRunning := Routes.SetupRouter()

	apiRunning.Run()
}