package main

import (
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/tinrab/retry"
)

type Config struct {
	DatabaseURL string `envconfig:"DATABASE_URL"`
}

func main(){
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	var r.catalog.Repository
	retry.ForeverSleep(2*time.Second, func(_int)(err error))
}