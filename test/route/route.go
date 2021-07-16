package route

import (
	"app/test/controllers"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func InitRouter(app *iris.Application) {
	bathUrl := "/api"
	mvc.New(app.Party(bathUrl + "/user")).Handle(controllers.NewUserController())
}