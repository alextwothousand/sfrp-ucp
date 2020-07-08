package config

type database struct {
	Username string
	Password string
	Database string
	Port     int
}

// Database contains the configuration for the database.
var Database database = database{
	Username: "root",
	Password: "",
	Database: "sfrp",
	Port:     3306,
}
