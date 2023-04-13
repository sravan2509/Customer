package CustomerSchema

type Customer struct {
	Name        string `json:"Name"`
	PhoneNumber string `json:"PhoneNumber"`
	Password    string `json:"Password"`
	Email       string `json:"Email"`
	Address     string `json:"Address"`
	// NewPassword     string `json:"NewPassword"`
	ConformPassword string `json:"ConformPassword"`
}

type Response struct {
	StatusCode int         `json:"StatusCode"`
	Message    string      `json:"Message"`
	Data       interface{} `json:"Data"`
}

type LoginCustomer struct {
	Password string `json:"Password"`
	Email    string `json:"Email"`
}

type ChangeLoginPassword struct {
	Email       string `json:"Email"`
	NewPassword string `json:"NewPassword"`
	Password    string `json:"OldPassword"`
}
