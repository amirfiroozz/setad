package main

import (
	"fmt"
	"log"
	"os"
	"setad/api/routers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	subRouter := r.Group("/api")
	routers.CreateRoutes(subRouter)
	fmt.Printf("server is running on: %v\n", getServerURL())
	err := r.Run(getServerURL())
	if err != nil {
		panic(err)
	}

}

func init() {
	loadEnvFiles()
}

func loadEnvFiles() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}
}

func getServerPort() string {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "3000"
	}
	return port
}

func getServerIP() string {
	IP := os.Getenv("SERVER_IP")
	if IP == "" {
		IP = "localhost"
	}
	return IP
}

func getServerURL() string {
	return fmt.Sprintf("%v:%v", getServerIP(), getServerPort())
}
