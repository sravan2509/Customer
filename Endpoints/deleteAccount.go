package endpoint

import (
	"net/http"

	Dbconfig "github.com/sravan2509/Customer/Dbconfig"
	Validation "github.com/sravan2509/Customer/Validation"
)

func DeleteCustomerHandler(w http.ResponseWriter, r *http.Request) {
	// Get the email address from the request context
	TokenEmail, ok := r.Context().Value("Email").(string)
	if !ok {
		ResponseFormat(w, "Email not found in context", http.StatusInternalServerError, nil)
		return
	}
	if r.Method != "DELETE" {
		ResponseFormat(w, "Method not allowed", http.StatusMethodNotAllowed, nil)
	}

	// db connection
	db, err := Dbconfig.DBConnection()
	if err != nil {
		ResponseFormat(w, "DB Connection Failed", http.StatusInternalServerError, nil)
		return
	}
	defer db.Close()

	//reading the Email from query
	Email := r.URL.Query().Get("Email")

	//validate the email
	if statusCode, err := Validation.DeleteCustomerValidation(db, Email, TokenEmail); err != nil {
		ResponseFormat(w, err.Error(), statusCode, nil)
		return
	}

	//Deleting the Customer
	if statusCode, err := Dbconfig.DeleteCustomer(db, Email); err != nil {
		ResponseFormat(w, err.Error(), statusCode, nil)
		return
	}

	ResponseFormat(w, "Customer Deleted Successfully!", http.StatusOK, nil)

}
