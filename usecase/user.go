package usecase

import (
	"capstone/constant"
	"capstone/middleware"
	"capstone/model"
	"capstone/model/payload"
	"capstone/repository/database"
	"capstone/utils"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func LoginUser(usernameOrEmail, password string) (*payload.LoginUserResponse, error) {
	// check to db email and password
	user, err := database.GetUserByUsernameOrEmail(usernameOrEmail)
	if err != nil {
		fmt.Println("LoginUser: Error getting user from the database")
		return nil, errors.New("username or email not found")
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
	if _, err := database.GetUserByUsername(req.Username); err == nil {
		return nil, errors.New("username already registered")
	}

	newUser := &model.User{
		FullName: req.Username,
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

func UpdateUser(user *model.User) (err error) {
	err = database.UpdateUser(user)
	if err != nil {
		fmt.Println("UpdateUser: Error updating user, err:", err)
		return
	}
	return
}

func ChangePasswordUser(UserID uint, payload *payload.ChangePasswordRequest) (err error) {
	user, err := database.GetUserByID(UserID)
	if err != nil {
		fmt.Println("GetUserByID: Error Get user, err:", err)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.OldPassword))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return errors.New("wrong password")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(payload.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(passwordHash)
	err = database.UpdateUser(user)
	if err != nil {
		fmt.Println("ChangePassword: Error Change Password, err:", err)
		return
	}
	return
}

func GetUserProfile(userID uint) (*payload.GetUserProfileResponse, error) {
	user, err := database.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	complaints, err := database.GetComplaintsByUserID(user.ID, constant.StatusAll)
	if err != nil {
		return nil, err
	}
	var laporan, pending, proccess, resolved uint
	for _, complaint := range complaints {
		switch complaint.Status {
		case constant.StatusPending:
			pending++
		case constant.StatusProccess:
			proccess++
		case constant.StatusResolved:
			resolved++
		}
		laporan++
	}
	resp := payload.GetUserProfileResponse{
		ID:           user.ID,
		PhotoProfile: utils.ConvertToNullString(user.PhotoProfile),
		FullName:     user.FullName,
		Laporan:      laporan,
		Pending:      pending,
		Proccess:     proccess,
		Resolved:     resolved,
	}
	return &resp, nil
}

func DeleteUserByID(userID uint) error {
	user, err := database.GetUserByID(userID)
	if err != nil {
		return err
	}

	complaints, err := database.GetComplaintsByUserID(user.ID, constant.StatusAll)
	if err != nil && err != errors.New("record not found") {
		return err
	}

	if len(complaints) > 0 {
		for _, complaint := range complaints {
			err = database.DeleteComplaint(complaint)
			if err != nil {
				return err
			}
		}
	}

	comments, err := database.GetCommentByUserID(user.ID)
	if err != nil && err != errors.New("record not found") {
		return err
	}

	if len(comments) > 0 {
		for _, comment := range comments {
			err = database.DeleteComment(comment)
			if err != nil {
				return err
			}
		}
	}

	pinned, err := database.GetPinnedComplaintsByUserId(user.ID)
	if err != nil && err != errors.New("record not found") {
		return err
	}

	if len(pinned) > 0 {
		for _, pin := range pinned {
			err = database.DeletePinnedComplaint(&pin)
			if err != nil {
				return err
			}
		}
	}

	err = database.DeleteUser(user)
	if err != nil {
		return err
	}
	return nil
}
