package main

import (
	"mongoexample/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	router.Router(r)

	r.Run()
}
