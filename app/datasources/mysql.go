package datasources

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"mutants/app/mutants/repositories"
	"os"
)

// Add gorm models to be auto migrated at the moment of build
func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&repositories.DnaDB{},
	)
}

// Format the string to add database init params
func dbParams(params ...string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", params[0], params[1], params[2], params[3], params[4])
}

// Connect with mysql database
func ConnectDb() *gorm.DB {
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	db, err := gorm.Open("mysql", dbParams(dbUsername,
		dbPassword,
		dbHost,
		dbPort,
		dbName))

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	return db
}
