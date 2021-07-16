package models

type User struct {
	UserId   int    `bson:"user_id"`
	Username string `bson:"username"`
	Password string `bson:"password"`
	Name     string `bson:"name"`
	Email    string `bson:"email"`
	Mobile   string `bson:"mobile"`
	QQ       string `bson:"qq"`
	Gender   int    `bson:"gender"`
	Age      int    `bson:"age"`
	Remark   string `bson:"remark"`
}

func (u User)CollectionName() string  {
	return "myrl"
}