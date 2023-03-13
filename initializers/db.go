package initializers

import (
	"fmt"

	"github.com/theronj60/project-horizon/helpers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error

	env := helpers.GetEnvData()
	dsn := "host=" + env.Host + " user=" + env.User + " password=" + env.Password + " dbname=" + env.Name + " port=" + env.Port

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed to connect to database.")
	}
}

// func SyncDB() {
// 	DB.AutoMigrate(&models.User{})
// }
