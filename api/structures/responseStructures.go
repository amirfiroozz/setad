package structures

type (
	SignupResponse struct {
		Message string      `json:"message"`
		UserID  interface{} `json:"userId"`
		Code    int         `json:"code"`
	}
)

func NewSignupResponse(message string, userID interface{}, code int) SignupResponse {
	return SignupResponse{message, userID, code}
}
