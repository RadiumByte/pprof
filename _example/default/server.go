package main

import (
	"github.com/RadiumByte/gin"
	"github.com/RadiumByte/pprof"
)

func main() {
	router := gin.Default()
	pprof.Register(router)
	router.Run(":8080")
}
