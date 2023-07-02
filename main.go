package main

import (
	"fmt"

	"github.com/arigates/golang-jwt/pkg"
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	err := route.Run(fmt.Sprintf("%s:%s", pkg.GodotEnv("APP_HOST"), pkg.GodotEnv("APP_PORT")))
	if err != nil {
		panic(err)
	}
}
