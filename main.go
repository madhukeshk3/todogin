package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/madhukeshk3/todogin/Config"
	"github.com/madhukeshk3/todogin/Models"
	"github.com/madhukeshk3/todogin/Routes"
)

var err error

func main(){
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		fmt.Println("Status: ",err)
	}

	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Todo{})

	r := Routes.SetupRouter()

	r.Run()
}