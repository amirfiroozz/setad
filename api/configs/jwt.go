package configs

import (
	"os"
	"strconv"
)

var (
	JWT_SECRET string
	JWT_EXP    int
)

func getJWTSecret() string {
	return os.Getenv("JWT_SECRET")
}
func getJWTExpirationTime() int {
	res, _ := strconv.Atoi(os.Getenv("JWT_EXP"))
	return res
}
