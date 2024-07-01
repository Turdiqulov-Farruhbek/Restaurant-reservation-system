package handlers

import (
	"api-getway/genproto/payment"
	"api-getway/genproto/reservation"
	"api-getway/genproto/user"

	"google.golang.org/grpc"
)

type HandlerStruct struct {
	User        user.UserServiceClient
	Payment     payment.PaymentServiceClient
	Reservation reservation.ReservationServiceClient
	Restaurant  reservation.RestaurantServiceClient
	Menu        reservation.MenuServiceClient
}

func NewHandler(conPay *grpc.ClientConn) *HandlerStruct {
	return &HandlerStruct{
		// User:        user.NewUserServiceClient(conMenu),
		Payment:     payment.NewPaymentServiceClient(conPay),
		// Reservation: reservation.NewReservationServiceClient(conReser),
	}
}
