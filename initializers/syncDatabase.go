package initializers

import(
	"github.com/Karthika-Rajagopal/jwt-go/models"
)

func SyncDatabase(){
	 // Migrate the schema
	DB.AutoMigrate(&models.User{})
}

