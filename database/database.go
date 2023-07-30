package database

import (
	"fmt"

	"github.com/joewilson27/rest-go-fiber-docker/database/db"
	"github.com/joewilson27/rest-go-fiber-docker/models"
)

func ConnectDb() {
	// dsn := fmt.Sprintf("host=db user=%s password=%s dbname=%s port=5432 sslmode=disable Timezone=Asia/Jakarta",
	// 	os.Getenv("DB_USER"),
	// 	os.Getenv("DB_PASSWORD"),
	// 	os.Getenv("DB_NAME"),
	// )

	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
	// 	Logger: logger.Default.LogMode(logger.Info),
	// })

	// if err != nil {
	// 	log.Fatal("Failed to connect to database. \n", err)
	// 	os.Exit(1)
	// }

	// log.Println("connecting to database")
	// db.Logger = logger.Default.LogMode(logger.Info)

	// log.Println("running migration")
	// db.AutoMigrate(&models.Fact{})

	// DB = Dbinstance{}
	// DB.Db = db

	if err := db.ConnectPg(); err == nil {

		db.DB.AutoMigrate(
			&models.Fact{},
		)
		fmt.Println("Success Migrate DB1")
	}

}
