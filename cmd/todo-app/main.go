package main

import (
	"flag"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/m-golang/todo-app/internal/todo/config"
	"github.com/m-golang/todo-app/internal/todo/database"
	"github.com/m-golang/todo-app/internal/todo/routes"
	repository "github.com/m-golang/todo-app/internal/todo/repository"
)

// Environment variable keys for MySQL configuration
const (
	pathEnvFile          = ".env"
	mysqlUserNameEnv     = "MYSQL_USER_NAME"     // MySQL username environment variable
	mysqlUserPasswordEnv = "MYSQL_USER_PASSWORD" // MySQL password environment variable
	mysqlDBNameEnv       = "MYSQL_DB_NAME"       // MySQL database name environment variable
)

func main() {
	// Load environment variables from a file (e.g., .env) to configure the application
	if err := godotenv.Load(pathEnvFile); err != nil {
		log.Fatal(err)
	}
	// Define a slice of environment variable keys needed for MySQL configuration
	keys := []string{mysqlUserNameEnv, mysqlUserPasswordEnv, mysqlDBNameEnv}

	// Load the values of the environment variables specified in the 'keys' slice
	dbConfigs, err := config.LoadEnvVars(keys)
	if err != nil {
		log.Fatal(err)
	}

	// Format the MySQL Data Source Name (DSN) using the loaded variables
	dsn := fmt.Sprintf("%s:%s@/%s?parseTime=true", dbConfigs[0], dbConfigs[1], dbConfigs[2])

	// Define a flag for the DSN (Data Source Name) to allow overriding through command-line arguments
	flag.StringVar(&dsn, "dsn", dsn, "MySQL data source name")

	// Parse the command-line flags
	flag.Parse()

	// Open a connection to the database using the DSN, and store the connection in repository.DB
	db, err := database.OpenDB(dsn)
	if err != nil {
		log.Fatal(err)
	}

	// Defer the closure of the database connection to ensure it is closed when the function completes.
	defer db.Close()

	// Set the global repository.DB variable to the database connection `db`.
	// This makes the database connection accessible in other parts of the application
	// through the `repository.DB` variable.
	repository.DB = db

	// Register the application's routes using the routes package
	r := routes.Router()

	// Start the Gin web server and listen for incoming requests
	r.Run()

}
