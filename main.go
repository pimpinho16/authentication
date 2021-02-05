package main

import (
	userRepo "authentication/repo/users"
	"authentication/utils/config"
	jwt "authentication/utils/jwt"
	//"authentication/model"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"log"
)

func main(){
	fmt.Println("Hello World")

	var log log.Logger
	//obteniendo configuraciones
	config,err := config.GetConfigs()

	if err != nil{
		fmt.Printf("Error reading config file")
	}

	//connection to database
	dsn := getDBConnection(config)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	userDB := userRepo.NewUserDataBase(db)

	result,err := userDB.IsUser("aandrade","12345")

	if err != nil {
		log.Printf("Error trying to find user %s  ","andrade")
	} else{
		tokenString,err := jwt.GenerateToken(int(result.ID),"12345",config.Jwt.ExpireTime,config.Jwt.Key)

		if err == nil{
			fmt.Println(tokenString)
		}
	}



	http.ListenAndServe(config.Server.Port,nil)
}


func getDBConnection(config config.Configurations) string{
	dns:= fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=%s",
		config.Database.Host,
		config.Database.Port,
		config.Database.Username,
		config.Database.Password,
		config.Database.Dbname,
		config.Database.Sslmode)

	return dns
}
