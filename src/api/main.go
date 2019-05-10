package main

import (
"github.com/gin-gonic/gin"
"github.com/mercadolibre/ejercicio1/src/api/controllers/result"
)

const(
	port = ":8080"
)

var (
	router = gin.Default()
)



func main(){

	router.GET("/response/:id/:url", result.GetResponse)
	router.GET("/responseWg/:id/:urlType", result.GetResponseWg)
	router.GET("/responseCh/:id/:urlType", result.GetResponseCh)
	//router.GET("/responseCh2/:id/:urlType", result.GetResponseCh2)
	router.GET("/ping", result.Ping)
	router.Run(port)
}


