package main

import (
	"ginvue/common"
	"github.com/gin-gonic/gin"
	_ "gorm.io/driver/mysql"
)




func main() {
	db := common.InitDB()
	defer db.Close()

	r := gin.Default()

	r = CollectRoute(r)
	panic(r.Run())
}
