package config

var DBHost string
var DBPort string
var DBName string
var DBUser string
var DBPassword string
var	DBSSLmode string



func init() {
	// database configuration
	DBHost = "localhost"
	DBPort = "5432"
	DBName = "teashop"
	DBUser = "postgres"
	DBPassword = "postgres"
	DBSSLmode= "disable"
}