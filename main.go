package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/levnzzz/ginEssential/common"
	"github.com/levnzzz/ginEssential/routers"
)

func main() {
	db := common.InitDB()
	defer db.Close()
	r := gin.Default()
	r = routers.CollectRouter(r)
	panic(r.Run())// 监听并在 0.0.0.0:8080 上启动服务
}
