package main

import (
	"ccgin/configs"
	"ccgin/internal/router"
	"fmt"
	"github.com/gin-gonic/gin"
)

func init() {
	configs.Once()
}

func main() {
	r := gin.Default()

	var err error

	router.RegisterRouter(r)

	fmt.Printf("Start running...\n")
	if err = r.Run(":3200"); err != nil {
		panic(err)
	}
}
