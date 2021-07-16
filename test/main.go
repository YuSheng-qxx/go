package main
import (
	"app/test/conf"
	"app/test/route"
	"flag"
	"github.com/kataras/iris/v12"
)

func main() {
	flag.Parse()
	app := iris.New()
	route.InitRouter(app)
	err := app.Run(iris.Addr(":"+conf.Sys.Port), iris.WithoutServerError(iris.ErrServerClosed))
	if err != nil {
		panic(err)
	}
}