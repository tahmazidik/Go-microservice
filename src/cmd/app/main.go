package main

import (
	"fmt"
	"github.com/tahmazidik/Go-microservice/internal/config"
)

func mian() {
	config.LoadEnviroment()

	appConfig := config.NewConfig()

	fmt.Println(fmt.Sprintf("User: %s", appConfig.Database.User))
	fmt.Println(fmt.Sprintf("Host: %s", appConfig.Database.Host))
	fmt.Println(fmt.Sprintf("Password: %s", appConfig.Database.Password))
	fmt.Println(fmt.Sprintf("DbName: %s", appConfig.Database.Dbname))
	fmt.Println(fmt.Sprintf("Port: %d", appConfig.Database.Port))
}
