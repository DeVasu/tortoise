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
	var username, password, host, db_name string
	fmt.Println("You need to add information about MYSQL server to start")


//******** UPDATE DATABASE CONNECTION INFO HERE *************
//******** UPDATE DATABASE CONNECTION INFO HERE *************
//******** UPDATE DATABASE CONNECTION INFO HERE *************
//******** UPDATE DATABASE CONNECTION INFO HERE *************
	host = "localhost:3306"
	username = "root"
	password = ""
	db_name = "tortoise"



	tortoise_db.Init(username, password, host, db_name)
	mapUrls()

//***********UPDATE PORT INFO HERE************
	router.Run(":3030")
}
