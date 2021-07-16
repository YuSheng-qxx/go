package datasource

import "C"
import (
	"app/test/models"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type UserRepository interface {
	PostUser(user models.User) (int, models.User)
	GetUserByUsernameAndPwd(username string, password string) (user models.User)
	GetUserByUsername(username string) (user models.User)
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

type userRepository struct{}

func(n userRepository) PostUser(user models.User) (int, models.User){
	code := 0
	col := GetCollection()
	err := col.Insert(&models.User{user.UserId, user.Username, user.Password, user.Name,
		user.Email, user.Mobile, user.QQ, user.Gender, user.Age, user.Remark})

	if err != nil {
		log.Println(err)
		code = -1
	}
	return code, user
}

func(n userRepository) GetUserByUsernameAndPwd(username string, password string) (user models.User){
	col := GetCollection()

	err :=col.Find(bson.M{"username":username,"password":password}).One(&user)
	if err != nil {
		log.Println(err)
	}
	return
}

func(n userRepository) GetUserByUsername(username string) (user models.User) {
	col := GetCollection()
	err :=col.Find(bson.M{"username":username}).One(&user)
	if err != nil {
		log.Println(err)
	}
	return
}