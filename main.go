package main

import (
	"authentication/config"
	jwt "authentication/jwt"
	userRepo "authentication/repo/user"
	//"authentication/model"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
)

func main(){
	//var claims model.MyClaims
	fmt.Println("Hello World")

	//obteniendo configuraciones
	config,err := config.GetConfigs()

	if err != nil{
		fmt.Printf("Error reading config file")
	}

	//connection to database
	dsn := getDBConnection(config)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	userDB := userRepo.NewUserDataBase(db)

	result := userDB.IsUser("aandrade","12345")
	if result {
		tokenString,err := jwt.GenerateToken("aandrade","12345",config.ExpireTime)

		if err!= nil{
			fmt.Println(tokenString)
		}
	}



	http.ListenAndServe(config.Server.Port,nil)
}


func getDBConnection(config  config.Configurations) string{
	dns:= fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=%s",
		config.Database.Host,
		config.Database.Port,
		config.Database.User,
		config.Database.Password,
		config.Database.Dbname,
		config.Database.Sslmode)

	return dns
}
