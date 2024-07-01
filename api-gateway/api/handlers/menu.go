package handlers

import (
	pb "api-getway/genproto/reservation" // Update this path to match your actual protobuf path
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetMenuByIdHandler gets a menu item by ID.
// @Summary Get menu by ID
// @Description Get menu by ID
// @Tags menu
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id path string true "Menu ID"
// @Success 200 {object} pb.GetMenuResponse
// @Failure 500 {object}  string "Internal server error"
// @Router /api/v1/menu/{id} [get]
func (h *HandlerStruct) GetMenuByIdHandler(c *gin.Context) {
	id := c.Param("id")
	req := pb.GetMenuRequest{Id: id}

	resp, err := h.Menu.GetMenu(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// CreateMenuHandler creates a new menu item.
// @Summary Create new menu
// @Description Create new menu
// @Tags menu
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param menu body pb.CreateMenuRequest true "Menu"
// @Success 200 {object} string "Invalid request body"
// @Failure 500 {object} string "Internal server error"
// @Router /api/v1/menu [post]
func (h *HandlerStruct) CreateMenuHandler(c *gin.Context) {
	var req pb.CreateMenuRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.Menu.CreateMenu(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateMenuHandler updates an existing menu item.
// @Summary Update menu
// @Description Update menu
// @Tags menu
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id path string true "Menu ID"
// @Param menu body pb.UpdateMenuRequest true "Menu"
// @Success 200 {object} string "Invalid request body"
// @Failure 500 {object} string "Internal server error"
// @Router /api/v1/menu/{id} [put]
func (h *HandlerStruct) UpdateMenuHandler(c *gin.Context) {
	id := c.Param("id")
	var req pb.UpdateMenuRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.Menu.Id = id
	resp, err := h.Menu.UpdateMenu(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// DeleteMenuHandler deletes a menu item.
// @Summary Delete menu
// @Description Delete menu
// @Tags menu
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id path string true "Menu ID"
// @Success 200 {object} string "Invalid request body"
// @Failure 500 {object} string "Internal server error"
// @Router /api/v1/menu/{id} [delete]
func (h *HandlerStruct) DeleteMenuHandler(c *gin.Context) {
	id := c.Param("id")
	req := pb.DeleteMenuRequest{Id: id}

	resp, err := h.Menu.DeleteMenu(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetAllMenuHandler returns all menu items.
// @Summary Get all menus
// @Description Get all menus
// @Tags menu
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Success 200 {object} string "Invalid request body"
// @Failure 500 {object} string "Internal server error"
// @Router /api/v1/menu [get]
func (h *HandlerStruct) GetAllMenuHandler(c *gin.Context) {
	req := pb.GetAllMenusRequest{}

	resp, err := h.Menu.GetAllMenus(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
