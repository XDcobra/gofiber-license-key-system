package main

import (
	"fmt"
	dummycontroller "github.com/XDcobra/gofiber-license-key-system/controller/DummyController"
	mysqlcontroller "github.com/XDcobra/gofiber-license-key-system/controller/MySQLController"
	rediscontroller "github.com/XDcobra/gofiber-license-key-system/controller/RedisController"
	dbmanagementMySQL "github.com/XDcobra/gofiber-license-key-system/database/MySQL"
	dbmanagementRedis "github.com/XDcobra/gofiber-license-key-system/database/Redis"
	router "github.com/XDcobra/gofiber-license-key-system/router"
	"log"
)

// Main function
func main() {
	// create connection to redis database
	redisClient := dbmanagementRedis.ConnectionRedisDB()

	// create connection to mysql database using GORM
	mysqlClient := dbmanagementMySQL.ConnectionMySQLDB()

	// do mysql model automigration
	err := dbmanagementMySQL.Automigration(mysqlClient)
	if err != nil {
		log.Fatalf("Error while migrating MySQL Database: %v", err)
	}

	// create server / register middlewares
	app := router.CreateServer()

	// create controllers
	dummyController := dummycontroller.NewDummyController()
	redisController := rediscontroller.NewRedisController(redisClient)
	mysqlController := mysqlcontroller.NewMySQLController(mysqlClient)

	// register routes
	routerClient := router.RegisterRoutes(app, redisController, dummyController, mysqlController)

	// start http server
	err = routerClient.Listen(":8000")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Listening on 8000")
	}

}
