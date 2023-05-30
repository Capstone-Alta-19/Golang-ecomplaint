package usecase

import (
	"capstone/middleware"
	"capstone/model"
	"capstone/model/payload"
	"capstone/repository/database"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func LoginUser(usernameOrEmail, password string) (*payload.LoginUserResponse, error) {
	// check to db email and password
	user, err := database.GetUserByUsernameOrEmail(usernameOrEmail)
	if err != nil {
		fmt.Println("LoginUser: Error getting user from the database")
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, errors.New("wrong password")
	}

	// generate jwt
	token, err := middleware.CreateToken(user.ID, user.Role)
	if err != nil {
		fmt.Println("LoginUser: Error generating token")
		return nil, err
	}

	user.Token = token
	res := payload.LoginUserResponse{
		Token: token,
	}
	return &res, nil
}

func CreateUser(req *payload.CreateUserRequest) (*payload.CreateUserResponse, error) {

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	if _, err := database.GetUserByEmail(req.Email); err == nil {
		return nil, errors.New("email already registered")
	}

	newUser := &model.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(passwordHash),
		Phone:    req.Phone,
		Role:     "USER",
	}

	err = database.CreateUser(newUser)
	if err != nil {
		return nil, err
	}

	// generate jwt
	token, err := middleware.CreateToken(newUser.ID, newUser.Role)
	if err != nil {
		fmt.Println("GetUser : Error Generate token")
		return nil, err
	}

	newUser.Token = token
	err = database.UpdateUser(newUser)
	if err != nil {
		fmt.Println("UpdateUser : Error Update user")
		return nil, err
	}
	result := payload.CreateUserResponse{
		Token: newUser.Token,
	}
	return &result, nil
}

func GetListUsers() (users []model.User, err error) {
	users, err = database.GetUsers()
	if err != nil {
		fmt.Println("GetListUsers: Error getting users from the database")
		return
	}
	return
}

func UpdateUser(user *model.User) (err error) {
	err = database.UpdateUser(user)
	if err != nil {
		fmt.Println("UpdateUser: Error updating user, err:", err)
		return
	}
	return
}
