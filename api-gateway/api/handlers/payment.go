package handlers

import (
	pb "api-getway/genproto/payment" // Update this path to match your actual protobuf path
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreatePaymentHandler creates a new payment.
// @Summary Create new payment
// @Description Create new payment
// @Tags payment
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param payment body pb.Payment true "Payment"
// @Success 200 {object} string "Invalid request body"
// @Failure 500 {object} string "Internal server error"
// @Router /api/v1/payments [post]
func (h *HandlerStruct) CreatePaymentHandler(c *gin.Context) {
	var req pb.CreatePaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(">>>>>>>", req)
	resp, err := h.Payment.CreatePayment(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetByIdPaymentHandler gets a payment by ID.
// @Summary Get payment by ID
// @Description Get payment by ID
// @Tags payment
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id path string true "Payment ID"
// @Success 200 {object} string "Invalid request body"
// @Failure 500 {object} string "Internal server error"
// @Router /api/v1/payments/{id} [get]
func (h *HandlerStruct) GetByIdPaymentHandler(c *gin.Context) {
	id := c.Param("id")
	req := pb.GetPaymentRequest{Id: id}

	resp, err := h.Payment.GetPayment(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdatePaymentHandler updates an existing payment.
// @Summary Update payment
// @Description Update payment
// @Tags payment
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id path string true "Payment ID"
// @Param payment body pb.UpdatePaymentRequest true "Payment"
// @Success 200 {object} string "Invalid request body"
// @Failure 500 {object} string "Internal server error"
// @Router /api/v1/payments/{id} [put]
func (h *HandlerStruct) UpdatePaymentHandler(c *gin.Context) {
	id := c.Param("id")
	var req pb.UpdatePaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.Payment.Id = id
	resp, err := h.Payment.UpdatePayment(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
