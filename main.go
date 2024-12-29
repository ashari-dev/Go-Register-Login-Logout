package main

import (
	"trial/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routers.RouterCombain(r)
	r.Run(":8080")
}
