package models

type (
	SignupResponse struct {
		Message string      `json:"message"`
		UserID  interface{} `json:"userId"`
		Code    int         `json:"code"`
	}
	LoginResponse struct {
		Message string `json:"message"`
		JWT     string `json:"jwt"`
		Code    int    `json:"code"`
	}
)

func NewSignupResponse(message string, userID interface{}, code int) SignupResponse {
	return SignupResponse{message, userID, code}
}
func NewLoginResponse(message string, jwt string, code int) LoginResponse {
	return LoginResponse{message, jwt, code}
}
