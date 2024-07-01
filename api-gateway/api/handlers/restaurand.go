package handlers

import (
	pb "api-getway/genproto/reservation" // Update this path to match your actual protobuf path
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetRestaurantByIdHandler gets a restaurant by ID.
// @Summary Get restaurant by ID
// @Description Get restaurant by ID
// @Tags restaurant
// @Accept  json
// @Produce  json
// @Param id path string true "Restaurant ID"
// @Success 200 {object} string "Invalid request body"
// @Failure 500 {object} string "Internal server error"
// @Router /api/v1/restaurants/{id} [get]
func (h *HandlerStruct) GetRestaurantByIdHandler(c *gin.Context) {
	id := c.Param("id")
	req := pb.GetRestaurantRequest{Id: id}

	resp, err := h.Restaurant.GetRestaurant(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// CreateRestaurantHandler creates a new restaurant.
// @Summary Create new restaurant
// @Description Create new restaurant
// @Tags restaurant
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param restaurant body pb.CreateRestaurantRequest true "Restaurant"
// @Success 200 {object} string "Invalid request body"
// @Failure 500 {object} string "Internal server error"
// @Router /api/v1/restaurants [post]
func (h *HandlerStruct) CreateRestaurantHandler(c *gin.Context) {
	// fmt.Println( c.GetHeader("Authorization"), c.Request.Header)
	var req pb.CreateRestaurantRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("11134243243243423",h.Restaurant)
	resp, err := h.Restaurant.CreateRestaurant(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateRestaurantHandler updates an existing restaurant.
// @Summary Update restaurant
// @Description Update restaurant
// @Tags restaurant
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id path string true "Restaurant ID"
// @Param restaurant body pb.UpdateRestaurantRequest true "Restaurant"
// @Success 200 {object} string "Invalid request body"
// @Failure 500 {object} string "Internal server error"
// @Router /api/v1/restaurants/{id} [put]
func (h *HandlerStruct) UpdateRestaurantHandler(c *gin.Context) {
	id := c.Param("id")
	var req pb.UpdateRestaurantRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.Restaurant.Id = id
	resp, err := h.Restaurant.UpdateRestaurant(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// DeleteRestaurantHandler deletes a restaurant.
// @Summary Delete restaurant
// @Description Delete restaurant
// @Tags restaurant
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id path string true "Restaurant ID"
// @Success 200 {object} string "Invalid request body"
// @Failure 500 {object} string "Internal server error"
// @Router /api/v1/restaurants/{id} [delete]
func (h *HandlerStruct) DeleteRestaurantHandler(c *gin.Context) {
	id := c.Param("id")
	req := pb.DeleteRestaurantRequest{Id: id}

	resp, err := h.Restaurant.DeleteRestaurant(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetRestaurantAllHandler returns all restaurants.
// @Summary Get all restaurants
// @Description Get all restaurants
// @Tags restaurant
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Success 200 {object} string "Invalid request body"
// @Failure 500 {object} string "Internal server error"
// @Router /api/v1/restaurants [get]
func (h *HandlerStruct) GetRestaurantAllHandler(c *gin.Context) {
	req := pb.ListRestaurantsRequest{}

	resp, err := h.Restaurant.ListRestaurants(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
