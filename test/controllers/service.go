package controllers

import (
	"app/test/datasource"
	"app/test/models"
)

type UserService interface {
	Register(user models.User) (result models.Result)
	Login(m models.User) (result models.Result)
}
type userServices struct {}

func NewUserServices() UserService {
	return &userServices{}
}
var userRepo = datasource.NewUserRepository()

func (u userServices) Register(m models.User) (result models.Result){
	user := userRepo.GetUserByUsername(m.Username)
	if user.UserId != 0 {
		result.Msg = "The username repetition"
		return
	}
	code,p := userRepo.PostUser(m)
	if code == -1 {
		result.Code = -1
		result.Msg = "fail to save"
		return
	}
	result.Code = 0
	result.Data = p
	return
}

func (u userServices) Login(m models.User) (result models.Result) {
	if m.Username == "" {
		result.Code = -1
		result.Msg = "Please enter a user name:"
		return
	}
	if m.Password == "" {
		result.Code = -1
		result.Msg = "Please enter the password:"
		return
	}
	user := userRepo.GetUserByUsernameAndPwd(m.Username, m.Password)
	if user.UserId == 0 {
		result.Code = -1
		result.Msg = "The user name or password error!"
		return
	}
	result.Code = 0
	result.Data = user
	result.Msg = "succeed to login"
	return
}
