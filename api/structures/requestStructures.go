package structures

type (
	LoginRequest struct {
		PhoneNumber string `json:"phoneNumber"`
		Password    string `json:"password"`
	}
	SignupRequest struct {
		FirstName   string `json:"firstName"`
		LastName    string `json:"lastName"`
		Password    string `bson:"password" json:"password"`
		PhoneNumber string `json:"phoneNumber"`
	}
)

func NewSignupResuest() SignupRequest {
	return SignupRequest{}
}
func NewLoginResuest() LoginRequest {
	return LoginRequest{}
}
