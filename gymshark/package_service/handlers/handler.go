package handlers

import (
	"github.com/emreisler/gymshark/package_service/usecases"
	"github.com/emreisler/gymshark/package_service/validators"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CalculatePacksHandler(c *gin.Context) {
	orderQuantityStr := c.Query("orderQuantity")
	orderQuantity, err := strconv.Atoi(orderQuantityStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid orderQuantity"})
		return
	}

	packsNeeded := usecases.CalculatePacks(usecases.GetPackSizes(), orderQuantity)
	c.JSON(http.StatusOK, packsNeeded)
}

func UpdatePackSizesHandler(c *gin.Context) {
	var request struct {
		PackSizes []int `json:"packSizes"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	if !validators.ValidatePackSizes(request.PackSizes) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format, pack size should be comma separated and will not include 0 size"})
		return
	}
	usecases.SetPackSizes(request.PackSizes)
	c.JSON(http.StatusOK, gin.H{"message": "Pack sizes updated"})
}
