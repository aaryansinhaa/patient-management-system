package main

import (
	"fmt"

	"github.com/aaryansinhaa/patient-management-system/internals/config"
	"github.com/aaryansinhaa/patient-management-system/internals/database"
)

func main() {
	config := config.MustLoadConfig()
	fmt.Printf("Environment: %s\n", config.Env)
	fmt.Printf("Description: %s\n", config.Description)
	fmt.Printf("HTTP Server Host: http://%s\n", config.HTTPServerConfig.Host)

	connection, err := database.LoadPSqlDb(&config.DatabaseConfig)
	if err != nil {
		fmt.Printf("Failed to connect to the database: %v\n", err)
		return
	}
	defer connection.Connection.Close()
	fmt.Printf("Database connection established successfully. %v", connection.Connection.Stats().OpenConnections)

}
