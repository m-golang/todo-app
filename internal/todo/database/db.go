package database

import "database/sql"

// OpenDB establishes a connection to the database using the provided Data Source Name (DSN).
// It attempts to open the connection and then pings the database to ensure the connection is active.
// If successful, it returns a reference to the opened database connection. If any error occurs,
// such as an invalid DSN or inability to ping the database, it returns an error.
func OpenDB(dsn string) (*sql.DB, error) {
	// Open the database connection using the provided DSN
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Ping the database to check if the connection is established successfully
	if err := db.Ping(); err != nil {
		return nil, err
	}
	
	// Return the opened database connection if no errors occurred
	return db, nil
}
