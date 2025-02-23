package main

import (
	"fmt"
	"time"

	"github.com/radenadri/go-boilerplate/config"
	"github.com/radenadri/go-boilerplate/internal/delivery/http/routes"
	"github.com/radenadri/go-boilerplate/pkg"
)

func main() {
	time.LoadLocation(config.AppTimezone)

	// Initialize logger
	config.InitLogger()

	// Init database
	config.InitDB()
	config.InitRedis()

	// Init validator
	pkg.InitValidator()

	r := routes.InitRouter()
	if err := r.Run(fmt.Sprintf(":%s", config.AppPort)); err != nil {
		panic(err)
	}
}
