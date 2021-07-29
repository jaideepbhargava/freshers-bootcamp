package main

import (
	"fmt"
	"freshers-bootcamp/Day3/exercise2/Config"
	"freshers-bootcamp/Day3/exercise2/Models"
	"freshers-bootcamp/Day3/exercise2/Routes"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql",
		Config.DbURL(Config.BuildDB()))

	if err != nil{
		fmt.Println("Status: ", err)
	}

	defer Config.DB.Close()

	Config.DB.AutoMigrate(&Models.Student{})

	apiRunning := Routes.SetupRouter()

	apiRunning.Run()
}