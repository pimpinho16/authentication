package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Configurations struct{
	Server     ServerConfigurations
	Database   DatabaseConfigurations
	Jwt Jwt
}

type ServerConfigurations struct{
	Port string
}

type DatabaseConfigurations struct{
	Host string
	Port string
	Username string
	Password string
	Dbname string
	Sslmode string
}

type Jwt struct {
	ExpireTime string
	Key	       string

}
func GetConfigs()  (Configurations,error){
	//set the file  of the configurations file
	viper.SetConfigName("application")

	//set the path to look for the configurations file
	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err!= nil{
		fmt.Println("Error reading config file,%s", err)
	}

	var config Configurations
	err := viper.Unmarshal(&config)

	if err != nil {
		fmt.Printf("unable to decode into struct, %v",err)
		return Configurations{},err
	}

	return config,nil

}

