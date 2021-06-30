package main

import (
	"github.com/booklibrary/logging"
	"github.com/booklibrary/util"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

var config *util.Config

func init() {
	log.SetFlags(1)
	logging.SetUpLogging()
	myconfig, err := util.LoadConfig(".")
	if err != nil {
		logging.AppLog.WriteLogsWError("cannot load config:" ,
			map[string]interface{} {"source": config, "err": err})
	}
	config = &myconfig

	util.LoadCollection("./static/")
	util.InitializeRedis()
}

func main() {
	logging.AppLog.WriteLogsInfo("Application running in environment:",
		map[string]interface{} {
		"runtime_setup": viper.GetString("RUNTIME_SETUP"),
		"app_port": viper.GetInt("APP_PORT")})

	var router *gin.Engine
	router = gin.Default()

	router.Static("/static/", "./static")
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.GET("/collection", getCollection)

	router.Run(viper.GetString("SERVER_ADDRESS") + ":" +
		viper.GetString("APP_PORT"))
}

func getCollection(c *gin.Context) {
	val := util.GetBookList()
	c.HTML(http.StatusOK, "library.html", gin.H{
		"books": val.BookList,
	})
}
