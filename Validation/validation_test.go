package customerValidation

import (
	"net/http"
	"testing"

	// "github.com/DATA-DOG/go-sqlmock" //for mocking the db calls
	mocks "github.com/sravan2509/Customer/mock_Dbconfig"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	Schema "github.com/sravan2509/Customer/Schema"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestSignupValidation(t *testing.T) {

	// mockDB, mock, err := sqlmock.New()
	// if err != nil {
	// 	t.Fatalf("Failed to create mock DB: %s", err)
	// }
	// defer mockDB.Close()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// Create a mock DB
	mockDB := mocks.NewMockDB(ctrl)

	newCustomer := Schema.Customer{
		Email:           "test@example.com",
		Password:        "Password123",
		ConformPassword: "Password123",
		PhoneNumber:     "1234567890",
		Name:            "John",
	}

	//testcase for valid user signup
	// mock.ExpectQuery("SELECT count(*) FROM customers WHERE Email=?").WithArgs(newCustomer.Email).
	// 	WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))//one way
	mockDB.EXPECT().IsCustomerExist(gomock.Any(), newCustomer.Email).Return(false) //another way
	status, err := SignupValidation(mockDB, newCustomer)
	assert.Equal(t, http.StatusCreated, status)
	assert.Nil(t, err, "Signup should be Successful!")

	//testcase for invalid signup
	//testcase invalid mail
	newCustomer.Email = "Invalid mail"
	status, err = SignupValidation(mockDB, newCustomer)
	assert.Equal(t, http.StatusBadRequest, status)
	assert.Error(t, err, "Error: Invalid Mail Address")
	//already Exists customer
	newCustomer.Email = "test@example.com"
	// mock.ExpectQuery("SELECT").WithArgs(newCustomer.Email).
	// 	WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))//one way
	mockDB.EXPECT().IsCustomerExist(gomock.Any(), newCustomer.Email).Return(false) //another way
	status, err = SignupValidation(mockDB, newCustomer)
	assert.Equal(t, http.StatusForbidden, status)
	assert.Error(t, err, "Error: User Already Exists")
	//password in valid
	newCustomer.Password = "test"
	status, err = SignupValidation(mockDB, newCustomer)
	assert.Equal(t, http.StatusBadRequest, status)
	assert.Error(t, err, "Error: Invalid Password")
	//Passwords Mismatch
	newCustomer.ConformPassword = "Password"
	status, err = SignupValidation(mockDB, newCustomer)
	assert.Equal(t, http.StatusBadRequest, status)
	assert.Error(t, err, "Error: Passwords mismatch")
	//Invalid phone
	newCustomer.PhoneNumber = "456456456"
	status, err = SignupValidation(mockDB, newCustomer)
	assert.Equal(t, http.StatusBadRequest, status)
	assert.Error(t, err, "Error: Invalid Phonenumber")

	// // Verify that the expected queries were executed
	// err = mock.ExpectationsWereMet()
	// if err != nil {
	// 	t.Fatalf("Failed to meet DB expectations: %s", err)
	// }

}

func TestLoginValidation(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to connect to database: %s", err)
	}
	defer db.Close()

	customer := Schema.LoginCustomer{
		Email:    "test@example.com",
		Password: "Password123",
	}

	mock.ExpectQuery("SELECT Password FROM customers WHERE Email = ?").
		WithArgs(customer.Email).
		WillReturnError(nil).
		WillReturnRows(sqlmock.NewRows([]string{"Password"}).
			AddRow(bcrypt.GenerateFromPassword([]byte("Password123"), bcrypt.DefaultCost)))
	status, err := LoginValidation(db, customer)
	assert.Equal(t, status, http.StatusOK)
	assert.Nil(t, err, "Login must succeed!")

}
