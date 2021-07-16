package controllers

import (
	"app/test/models"
	"github.com/kataras/iris/v12"
	"log"
)

type UserController struct {
	Ctx     iris.Context
	Service UserService
}

func NewUserController() *UserController {
	return &UserController{Service: NewUserServices()}
}

func (g *UserController) PostRegister() (res models.Result){
	var user models.User
	if err := g.Ctx.ReadJSON(&user); err != nil {
		log.Println(err)
		res.Msg = "Data error"
		return
	}
	return g.Service.Register(user)
}

func (g *UserController) PostLogin() models.Result {
	var m models.User
	err := g.Ctx.ReadJSON(&m)
	if err != nil {
		log.Println("ReadJSON Error:", err)
	}
	result := g.Service.Login(m)
	return result
}