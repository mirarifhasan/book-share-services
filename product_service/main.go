package main

import (
	"boi.com/product_service/config"
	"fmt"
	"os"
)

func main() {

	config.LoadEnv()

	DB_DRIVER := os.Getenv("DB_DRIVER")
	fmt.Println(DB_DRIVER)

	fmt.Println("Product service is running...")
}
