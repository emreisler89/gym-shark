package main

import (
	"fmt"
	"github.com/emreisler/gymshark/package_service/handlers"
	"github.com/emreisler/gymshark/package_service/usecases"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	//grab package size from .env
	packSizesStr := os.Getenv("PACKAGE_SIZES")
	packSizesStrList := strings.Split(packSizesStr, ",")
	for _, sizeStr := range packSizesStrList {
		size, err := strconv.Atoi(sizeStr)
		if err != nil {
			log.Fatalf("Invalid package size:%v", err)
		}
		usecases.PackSizes = append(usecases.PackSizes, size)
	}

	usecases.SetPackSizes(usecases.PackSizes)

	r := gin.Default()

	// Configure CORS to allow all origins
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	r.GET("/calculate-packs", handlers.CalculatePacksHandler)
	r.POST("/update-pack-sizes", handlers.UpdatePackSizesHandler)

	port := 9092 // You can change this to the desired port number
	address := fmt.Sprintf(":%d", port)
	r.Run(address)
}
