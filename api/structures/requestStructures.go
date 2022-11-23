package structures

type (
	LoginRequest struct {
		PhoneNumber string `json:"phoneNumber"`
		Password    string `json:"password"`
	}
)
