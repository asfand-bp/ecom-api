package tests

import (
	"ecom/resources/auth"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	// User data
	userData := auth.User{
		Name:     "Test User",
		Email:    "testuser@gmail.com",
		Password: "test",
		Phone:    "23498285328",
		Address:  "4th Street somewhere in California",
		City:     "Los Angeles",
		State:    "CA",
		Country:  "USA",
		Role:     "basic",
	}

	// Make POST api request
	w := MakePostRequest(t, R, "/users", userData)

	var user auth.User
	json.Unmarshal(w.Body.Bytes(), &user)

	// Assertions
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.NotEmpty(t, user)

	// Assigning the user id for use in later tests
	DeleteUser(t, R, fmt.Sprintf("%v", user.ID))
}

func TestGetUsers(t *testing.T) {
	user_id := CreateUser(t, R).ID

	// Make GET api request
	w := MakeGetRequest(t, R, "/users")

	var users []auth.User
	json.Unmarshal(w.Body.Bytes(), &users)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, users)

	DeleteUser(t, R, fmt.Sprintf("%v", user_id))
}

func TestGetUser(t *testing.T) {
	user_id := CreateUser(t, R).ID

	// Make GET api request
	w := MakeGetRequest(t, R, fmt.Sprintf("/users/%v", user_id))

	var user auth.User
	json.Unmarshal(w.Body.Bytes(), &user)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, user)

	DeleteUser(t, R, fmt.Sprintf("%v", user_id))
}

func TestUpdateUser(t *testing.T) {
	user_id := CreateUser(t, R).ID

	NEW_NAME := "Test User New"
	NEW_ADDRESS := "Latifabad no 6, Hyderabad"
	body := auth.UserUpdateBody{
		Name:    NEW_NAME,
		Address: NEW_ADDRESS,
	}

	// Make PUT api request
	w := MakePutRequest(t, R, fmt.Sprintf("/users/%v", user_id), body)

	var user auth.User
	json.Unmarshal(w.Body.Bytes(), &user)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, user)
	assert.Equal(t, user.Name, NEW_NAME)
	assert.Equal(t, NEW_ADDRESS, user.Address)

	DeleteUser(t, R, fmt.Sprintf("%v", user_id))
}

func TestDeleteUser(t *testing.T) {
	user_id := CreateUser(t, R).ID

	// Make PUT api request
	w := MakeDeleteRequest(t, R, fmt.Sprintf("/users/%v", user_id))

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreateUserBadBody(t *testing.T) {
	// User data bad body
	userBadData := BadBodyType{
		BadField: "Bad",
	}

	// Make POST api request
	w := MakePostRequest(t, R, "/users", userBadData)

	var user auth.User
	json.Unmarshal(w.Body.Bytes(), &user)

	// Assertions
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Empty(t, user)
}

func TestUpdateUserBadBody(t *testing.T) {
	user_id := CreateUser(t, R).ID

	userBadData := BadBodyType{
		BadField: "Bad",
	}

	// Make PUT api request
	w := MakePutRequest(t, R, fmt.Sprintf("/users/%v", user_id), userBadData)

	var user auth.User
	json.Unmarshal(w.Body.Bytes(), &user)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Empty(t, user)

	DeleteUser(t, R, fmt.Sprintf("%v", user_id))
}

func TestGetUserByIDInvalidID(t *testing.T) {
	INVALID_USER_ID := "-1"

	// Make GET request
	w := MakeGetRequest(t, R, fmt.Sprintf("/users/%s", INVALID_USER_ID))

	// Assertions
	assert.Equal(t, http.StatusNotFound, w.Code)
}
