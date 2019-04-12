package db

import (
	"teashop/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

)

var Orm *gorm.DB

func init() {
	db, err := gorm.Open("postgres", 
		"host=" + config.DBHost + 
		" port=" + config.DBPort +
		" dbname=" + config.DBName +
		" user=" + config.DBUser + 
		" password=" + config.DBPassword +
		" sslmode="+ config.DBSSLmode)

	if err != nil {
	  panic(err)
	}
	
	Orm = db
}