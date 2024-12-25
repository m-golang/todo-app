package main

import (
	"database/sql"
	"flag"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/m-golang/todo-app/services"
)

// Define the application environment structure
type appEnv struct {
	services *services.SerEnv
}

func main() {
	// Parse the Data Source Name (DSN) flag for MySQL connection string
	dsn := flag.String("dsn", "dbusername:dbpassword@/tododb?parseTime=true", "MySQL data source name")

	flag.Parse()


	// Open DB connection
	db, err := openDB(*dsn)
	if err != nil {
		log.Fatal(err)
	}

	// Create a Gin router
	r := gin.Default()

	// Load HTML templates
	r.LoadHTMLGlob("ui/html/*")
	
	// Serving static files
	r.Static("/static","./ui/static")

	// Set up app environment with DB connection
	appEnv := &appEnv{
		services: &services.SerEnv{DB: db},
	}

	// Register routes 
	appEnv.Router(r)

	// Start the Gin server
	r.Run()

}

// Open DB connection and ping to ensure it's active
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
