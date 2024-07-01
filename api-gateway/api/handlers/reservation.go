package handlers

import (
	pb "api-getway/genproto/reservation" // Update this path to match your actual protobuf path
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateReservationHandler creates a new reservation.
// @Summary Create new reservation
// @Description Create new reservation
// @Tags reservation
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param reservation body pb.CreateReservationRequest true "Reservation"
// @Success 200 {object} string "Invalid request body"
// @Failure 500 {object} string "Internal server error"
// @Router /api/v1/reservations [post]
func (h *HandlerStruct) CreateReservationHandler(c *gin.Context) {
	var req pb.CreateReservationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.Reservation.CreateReservation(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetReservationAllHandler returns all reservations.
// @Summary Get all reservations
// @Description Get all reservations
// @Tags reservation
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Success 200 {object} string "Invalid request body"
// @Failure 500 {object} string "Internal server error"
// @Router /api/v1/reservations [get]
func (h *HandlerStruct) GetReservationAllHandler(c *gin.Context) {
	req := pb.ListReservationsRequest{}

	resp, err := h.Reservation.ListReservations(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetReservationByIdHandler gets a reservation by ID.
// @Summary Get reservation by ID
// @Description Get reservation by ID
// @Tags reservation
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id path string true "Reservation ID"
// @Success 200 {object} string "Invalid request body"
// @Failure 500 {object} string "Internal server error"
// @Router /api/v1/reservations/{id} [get]
func (h *HandlerStruct) GetReservationByIdHandler(c *gin.Context) {
	id := c.Param("id")
	req := pb.GetReservationRequest{Id: id}

	resp, err := h.Reservation.GetReservation(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateReservationHandler updates an existing reservation.
// @Summary Update reservation
// @Description Update reservation
// @Tags reservation
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id path string true "Reservation ID"
// @Param reservation body pb.UpdateReservationRequest true "Reservation"
// @Success 200 {object} string "Invalid request body"
// @Failure 500 {object} string "Internal server error"
// @Router /api/v1/reservations/{id} [put]
func (h *HandlerStruct) UpdateReservationHandler(c *gin.Context) {
	id := c.Param("id")
	var req pb.UpdateReservationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.Reservation.Id = id
	resp, err := h.Reservation.UpdateReservation(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// DeleteReservationHandler deletes a reservation.
// @Summary Delete reservation
// @Description Delete reservation
// @Tags reservation
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id path string true "Reservation ID"
// @Success 200 {object} string "Invalid request body"
// @Failure 500 {object} string "Internal server error"
// @Router /api/v1/reservations/{id} [delete]
func (h *HandlerStruct) DeleteReservationHandler(c *gin.Context) {
	id := c.Param("id")
	req := pb.DeleteReservationRequest{Id: id}

	resp, err := h.Reservation.DeleteReservation(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// ReservationCheckHandler checks a reservation.
// @Summary Check reservation
// @Description Check reservation
// @Tags reservation
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param check body  string true "Check"
// @Success 200 {object} string "Invalid request body"
// @Failure 500 {object} string "Internal server error"
// @Router /api/v1/reservations/check [post]
func (h *HandlerStruct) ReservationCheckHandler(c *gin.Context) {
	var req pb.CheckAvailabilityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.Reservation.CheckAvailability(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// ReservationOrderIdHandler orders by reservation ID.
// @Summary Order reservation by ID
// @Description Order reservation by ID
// @Tags reservation
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id path string true "Reservation ID"
// @Param order body pb.OrderFoodReq true "Order"
// @Success 200 {object} string "Invalid request body"
// @Failure 500 {object} string "Internal server error"
// @Router /api/v1/reservations/{id}/order [post]
func (h *HandlerStruct) ReservationOrderIdHandler(c *gin.Context) {
	id := c.Param("id")
	var req pb.OrderFoodReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.Id = id
	resp, err := h.Reservation.OrderFood(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// ReservationPaymentHandler processes payment for a reservation.
// @Summary Process payment for reservation
// @Description Process payment for reservation
// @Tags reservation
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id path string true "Reservation ID"
// @Param payment body pb.PaymentReservationRequest true "Payment"
// @Success 200 {object} string "Invalid request body"
// @Failure 500 {object} string "Internal server error"
// @Router /api/v1/reservations/{id}/payment [post]
func (h *HandlerStruct) ReservationPaymentHandler(c *gin.Context) {
	id := c.Param("id")
	var req pb.PaymentReservationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.Id = id
	resp, err := h.Reservation.PaymentReservation(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// ReservationOrderFood processes reservation for a reservation.
// @Summary Order reservation by ID
// @Description Order reservation by ID
// @Tags reservation
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id path string true "Reservation ID"
// @Param reservation body pb.OrderFoodReq true "Reservation Order"
// @Success 200 {object} string "Invalid request body"
// @Failure 500 {object} string "Internal server error"
// @Router /api/v1/reservations/{id}/reservation [post]

func (h *HandlerStruct) ReservationOrderFood(c *gin.Context) {
	id := c.Param("id")
    var req pb.OrderFoodReq
    if err := c.ShouldBindJSON(&req); err!= nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    req.Id = id
    resp, err := h.Reservation.OrderFood(context.Background(), &req)
    if err!= nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, resp)
}
