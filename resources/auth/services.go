package auth

import (
	"ecom/db"
	"ecom/utils"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

type Service struct {
	DB *gorm.DB
}

func NewService() *Service {
	return &Service{
		DB: db.DB,
	}
}

func (r *Service) GetAllUsers() ([]User, *utils.Error) {
	var users []User

	result := r.DB.Find(users)
	if result.Error != nil {
		return nil, utils.InternalServerError("Something went wrong")
	}

	return users, nil
}

func (r *Service) GetUserByID(id string) (*User, *utils.Error) {
	var data User

	result := r.DB.First(&data, id)
	if result.Error != nil {

		//if record not found
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, utils.NotFoundError("Invalid ID. Record not found")

		} else { //any other error
			return nil, utils.InternalServerError(result.Error.Error())
		}
	}

	return &data, nil
}

func (r *Service) CreateUser(body User) (*User, *utils.Error) {
	// Check if email already exists
	var existingUser User
	emailCheck := r.DB.Where("email = ?", body.Email).First(&existingUser)
	if emailCheck.Error == nil {
		return nil, utils.BadRequestError("Email already exists")
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(body.Password)
	if err != nil {
		return nil, utils.InternalServerError("Failed to hash password")
	}

	// Create a new User object and populate its fields based on the provided body
	user := User{
		Name:     body.Name,
		Email:    body.Email,
		Password: hashedPassword,
		Phone:    body.Phone,
		Address:  body.Address,
		City:     body.City,
		State:    body.State,
		Country:  body.Country,
		Role:     body.Role,
	}

	// Save the new User object to the database
	result := r.DB.Create(&user)
	if result.Error != nil {
		return nil, utils.InternalServerError(result.Error.Error())
	}

	return &user, nil
}

func (r *Service) UpdateUser(id string, body UserUpdateBody) (*User, *utils.Error) {

	var data User
	dataResult := r.DB.First(&data, "id = ?", id)

	if dataResult.Error == gorm.ErrRecordNotFound {
		return nil, &utils.Error{Status: http.StatusNotFound, Message: "Couldn't find the user. Try again."}
	}

	if dataResult.Error != nil {
		return nil, &utils.Error{Status: http.StatusInternalServerError, Message: dataResult.Error.Error()}
	}

	if (body.Name != data.Name) && (body.Name != "") {
		data.Name = body.Name
	}

	if body.Email != data.Email {

		var existingUser User

		emailCheck := r.DB.Where("email = ?", body.Email).First(&existingUser)

		if emailCheck.Error == nil {
			return nil, utils.BadRequestError("Email already exists, please choose a different email.")
		}
	}

	if body.Email != "" {
		data.Email = body.Email
	}

	if body.Password != "" {
		data.Password = body.Password
	}

	if body.Phone != "" {
		data.Phone = body.Phone
	}

	if body.Address != "" {
		data.Address = body.Address
	}

	if body.City != "" {
		data.City = body.City
	}

	if body.State != "" {
		data.State = body.State
	}

	if body.Country != "" {
		data.Country = body.Country
	}

	if body.Role != "" {
		data.Role = body.Role
	}

	// if body.Role != "" {
	// 	var roleUser User
	// 	roleResult := r.DB.First(&roleUser, "id = ?", body.Role)

	// 	if roleResult.Error != nil {
	// 		// checking for role if it exists
	// 		if roleResult.Error == gorm.ErrRecordNotFound {
	// 			return nil, utils.BadRequestError("Role not found.")

	// 		} else {
	// 			// Error querying the role, return an error
	// 			return nil, &utils.Error{Status: http.StatusInternalServerError, Message: roleResult.Error.Error()}
	// 		}
	// 	}

	// 	data.Role = body.Role
	// }

	// Save to the database
	result := r.DB.Save(&data)

	if result.Error != nil {
		return nil, &utils.Error{Status: http.StatusInternalServerError, Message: result.Error.Error()}
	}

	return &data, nil
}

func (r *Service) DeleteUser(id string) (map[string]string, *utils.Error) {
	var user User
	res := db.DB.First(&user, id)

	// If any error
	if res.Error != nil {
		// If record is not found with provided id
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, utils.NotFoundError("User not found with provided id")
		} else { // Else any other error
			return nil, utils.InternalServerError("Something went wrong")
		}
	}
	// All good, delete the user
	res = db.DB.Delete(&user, id)
	if res.Error != nil {
		return nil, utils.InternalServerError("Something went wrong")
	}

	return map[string]string{"message": "Successfully deleted the user"}, nil
}
