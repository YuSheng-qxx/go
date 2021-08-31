package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

type Person struct {
	name string
	age int
}

func register(ctx context.Context){

}

func main() {
	app := iris.New()
	app.Post("/post", func(context context.Context) {
		path := context.Path()
		app.Logger().Info(path)
		var person Person
		if err := context.ReadJSON(&person); err != nil{
			panic(err.Error())
		}
		context.Writef("received:%#+v\n", person)
	})

	app.Get("/userpath", func(ctx context.Context){
		path := ctx.Path()
		app.Logger().Info(path)
		ctx.WriteString("请求路径：" + path)
	})

	app.Handle("GET", "/register", register)

	app.Get("/hello", func(ctx context.Context){
		path := ctx.Path()
		app.Logger().Info(path)
		username := ctx.URLParam("username")
		app.Logger().Info(username)
		password := ctx.URLParam("password")
		app.Logger().Info(password)
		ctx.HTML("<h1>" + username + password + "</h1>")
	})

	app.Run(iris.Addr(":8000"))
}