package auto

import (
	"github.com/m00nreal/go-jwt-auth/api/database"
	"github.com/m00nreal/go-jwt-auth/api/models"
	"github.com/m00nreal/go-jwt-auth/api/utils/console"
	"log"
)

func Load() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Debug().DropTableIfExists(&models.User{}).Error
	if err != nil {
		log.Fatal(err)
	}
	err =  db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
		err = db.Debug().Model(&models.User{}).Create(&user).Error
		if err != nil {
			log.Fatal("Error creando usuario", err)
		}
		console.Pretty(user)
	}

}
