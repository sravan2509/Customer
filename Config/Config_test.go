package Config

import "testing"

func testSecretKey(t *testing.T) {
	result := SecretKey()

	expectedOutput := "MyCustomerSecretKey"

	if string(result) != expectedOutput {
		t.Errorf("Result is %q, Expected %q", result, expectedOutput)
	}

}

func testDbLink(t *testing.T) {
	result := GetConnectionString()

	expectedOutput := "root:Sravan@2509@tcp(localhost:3306)/golang?charset=utf8mb4&parseTime=True&loc=Local"

	if string(result) != expectedOutput {
		t.Errorf("Result is %q, Expected %q", result, expectedOutput)
	}
}
