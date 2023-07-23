package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joewilson27/rest-go-fiber-docker/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDb() {
	dsn := fmt.Sprintf("host=db user=%s password=%s dbname=%s port=5432 sslmode=disable Timezone=Asia/Jakarta",
										os.Getenv("DB_USER"),
										os.Getenv("DB_PASSWORD"),
									  os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
	})

	if err!= nil {
    log.Fatal("Failed to connect to database. \n", err)
		os.Exit(1)
  }

	log.Println("connecting to database")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("running migration")
	db.AutoMigrate(&models.Fact{})

	DB = Dbinstance{}
	DB.Db = db

}