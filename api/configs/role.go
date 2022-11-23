package configs

import "os"

var (
	USER_ROLE  string
	ADMIN_ROLE string
)

func getUserRole() string {
	return os.Getenv("USER_ROLE")
}
func getAdminRole() string {
	return os.Getenv("ADMIN_ROLE")
}
