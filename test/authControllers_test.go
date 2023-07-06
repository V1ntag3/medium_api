package test

import (
	"bytes"
	"encoding/json"
	"medium_api/database"
	"medium_api/routers"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

type User struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func TestRegister(t *testing.T) {
	// Define a structure for specifying input and output data
	// of a single test case
	tests := []struct {
		method       string
		description  string // description of the test case
		route        string // route path to test
		expectedCode int    // expected HTTP status code
		body         User
	}{
		// First test case
		{
			method:       "POST",
			description:  "get HTTP status 200",
			route:        "/api/auth/register",
			expectedCode: 200,
			body: User{
				Name:     "Marcos",
				Surname:  "Vinicius",
				Email:    "marcos@gmail.com",
				Password: "Pass1302@",
			},
		},
		// Second test case
		{
			method:       "POST",
			description:  "get HTTP status 404, when route is not exists",
			route:        "/not-found",
			expectedCode: 404,
			body: User{
				Name:     "Marcos",
				Surname:  "Vinicius",
				Email:    "marcos@gmail.com",
				Password: "Pass1302@",
			},
		},
	}
	database.Connect("db_test.db")
	// Define Fiber app.
	app := fiber.New()

	routers.Setup(app)

	// Iterate through test single test cases
	for _, test := range tests {
		user, _ := json.Marshal(test.body)
		// Create a new http request with the route from the test case
		req := httptest.NewRequest(test.method, test.route, bytes.NewBufferString(string(user)))
		// Perform the request plain with the app,
		// the second argument is a request latency
		// (set to -1 for no latency)
		resp, _ := app.Test(req)
		// Verify, if the status code is as expected
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}
