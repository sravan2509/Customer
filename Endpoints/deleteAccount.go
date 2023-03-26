package endpoint

import (
	"fmt"
	"net/http"

	Dbconfig "github.com/sravan2509/Customer/Dbconfig"
	validation "github.com/sravan2509/Customer/Validation"
)

func deleteCustomerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	db, err := Dbconfig.DBConnection()
	defer db.Close()
	Email := r.URL.Query().Get("Email")
	if !validation.IsCustomerExist(Email) {
		http.Error(w, "Customer not found", http.StatusBadRequest)
		return
	}

	_, err = db.Query(`DELETE FROM customers WHERE Email = ?`, Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Customer Deleted Successfully!")
	w.WriteHeader(http.StatusOK)

}
