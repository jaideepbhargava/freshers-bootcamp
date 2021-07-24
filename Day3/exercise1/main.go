package main

import (
	"fmt"
	"freshers-bootcamp/Day3/exercise1/Config"
	"freshers-bootcamp/Day3/exercise1/Models"
	"freshers-bootcamp/Day3/exercise1/Routes"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var err error

func main(){
	Config.DB, err = gorm.Open("mysql",
		Config.DbURL(Config.BuildDBConfig()))

	if err != nil{
		fmt.Println("Status: ", err)
	}

	defer Config.DB.Close()

	Config.DB.AutoMigrate(&Models.User{})

	apiRunning := Routes.SetupRouter()
	apiRunning.Run()
}