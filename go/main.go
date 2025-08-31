package main

import (
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jarvisjia/pmnt/go/task1"
	"github.com/jarvisjia/pmnt/go/task2"
	"github.com/jarvisjia/pmnt/go/task3"
	"github.com/jarvisjia/pmnt/go/task4/config"
	"github.com/jarvisjia/pmnt/go/task4/routers"
)

func main() {
	fmt.Println("==========Task1==================")
	task1.Task1()

	fmt.Println("==========Task2==================")
	task2.Task2()

	fmt.Println("==========Task3==================")
	task3.Task3()

	fmt.Println("==========Task4==================")
	db := config.DbConnet()
	config.InitDb(db)

	file, err := os.Create("task4/logs/task4.log")
	if err != nil {
		panic(err)
	}
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
	r := gin.Default()

	r.LoadHTMLGlob("task4/templates/**/*")

	routers.UserRouter(r, db)
	routers.PostRouter(r, db)
	routers.CommentRouter(r, db)
	r.Run()
}
