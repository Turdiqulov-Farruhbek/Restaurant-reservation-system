package service

import (
	"context"
	"log"
	"log/slog"
	"strconv"

	pb "github.com/Project_Restaurant/Payment-Service/genproto/payment"
	"github.com/Project_Restaurant/Payment-Service/storage/postgres"
)

type PaymentService struct {
	stg postgres.Storage
	pb.UnimplementedPaymentServiceServer
}

func NewPAymentService(stg postgres.Storage) *PaymentService {
	return &PaymentService{stg: stg}
}

func (ps *PaymentService) CreatePayment(ctx context.Context, req *pb.CreatePaymentRequest) (*pb.CreatePaymentResponse, error) {
	payment, err := ps.stg.Payments.CreatePayment(ctx, req)
	if err != nil {
		slog.Error("Error while create service", err)
		return &pb.CreatePaymentResponse{}, err
	}
	return &pb.CreatePaymentResponse{Payment: payment}, nil
}

func (ps *PaymentService) UpdatePayment(ctx context.Context, req *pb.UpdatePaymentRequest) (*pb.UpdatePaymentResponse, error) {
	payment, err := ps.stg.Payments.UpdatePayment(ctx, req)
	if err != nil {
		slog.Error("Error while update service", err)
		return &pb.UpdatePaymentResponse{}, err
	}
	return &pb.UpdatePaymentResponse{Payment: payment}, nil
}

func (ps *PaymentService) GetPayment(ctx context.Context, req *pb.GetPaymentRequest) (*pb.GetPaymentResponse, error) {
	payment, err := ps.stg.Payments.GetPaymentById(ctx, req)
	if err != nil {
		slog.Error("Error while update service", err)
		return &pb.GetPaymentResponse{}, err
	}
	return &pb.GetPaymentResponse{Payment: payment}, nil
}

func (ps *PaymentService) DeletePayment(ctx context.Context, req *pb.DeletePaymentRequest) (*pb.DeletePaymentResponse, error) {
	payment, err := ps.stg.Payments.DeletePayment(ctx, req)
	if err != nil {
		slog.Error("Error while update service", err)
		return &pb.DeletePaymentResponse{Message: "Error"}, err
	}
	return payment, nil
}

func (ps *PaymentService) PayForReservation(ctx context.Context, req *pb.PayForReservationReq) (*pb.PayForReservationRes, error) {
	remainingBalance, err := ps.stg.Payments.PayForReservation(ctx, req)
	if err != nil {
		log.Printf("Error while paying for reservation: %v", err)
		return nil, err
	}
	message := "You remaining balance : " + strconv.Itoa(int(remainingBalance))
	if remainingBalance <= 0 {
		message = "Every thing is paid"
	}
	return &pb.PayForReservationRes{Message: message}, nil
}
