package configs

import (
	"log"

	"github.com/joho/godotenv"
)

func init() {
	loadEnvFiles()
	DB_NAME = getDBName()
	MONGODB_FULL_URL = getMongoDBFullURL()
	MONGO_CLINET = getMongoClient()
	USER_ROLE = getUserRole()
	ADMIN_ROLE = getAdminRole()
	JWT_EXP = getJWTExpirationTime()
	JWT_SECRET = getJWTSecret()
}

func loadEnvFiles() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}
}
