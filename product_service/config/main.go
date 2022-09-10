package config

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Env Cannot Loaded ", err.Error())
	}
	log.Println("Env Loaded")
}
