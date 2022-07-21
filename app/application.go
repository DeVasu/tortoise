package app

import (
	"fmt"

	"github.com/DeVasu/tortoise/datasources/mysql/tortoise_db"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {

	fmt.Println("You need to add information about MYSQL server to start")

	var username, password, host, db_name string
	host = "localhost:3306"
	username = "root"
	password = ""
	db_name = "tortoise"

	tortoise_db.Init(username, password, host, db_name)


	mapUrls()


	router.Run(":3030")
}
