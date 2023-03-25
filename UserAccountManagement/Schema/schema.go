package CustomerSchema

type Customer struct {
	Name        string `json:"Name"`
	PhoneNumber string `json:"PhoneNumber"`
	Password    string `json:"Password"`
	Email       string `json:"Email"`
	Address     string `json:"Address"`
}

type LoginCustomer struct {
	Password string `json:"Password"`
	Email    string `json:"Email"`
}

type ChangeLogin struct {
	Email       string `json:"Email"`
	NewPassword string `json:"NewPassword"`
	OldPassword string `json:"OldPassword"`
}
