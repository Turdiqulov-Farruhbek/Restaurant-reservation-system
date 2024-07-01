package postgres

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"time"

	pb "github.com/Project_Restaurant/Payment-Service/genproto/payment"

	"github.com/google/uuid"
)

type PaymentRepo struct {
	Db *sql.DB
}

func NewPaymentRepo(db *sql.DB) *PaymentRepo {
	return &PaymentRepo{Db: db}
}

// ErrPaymentNotFound is returned when a payment is not found.
var ErrPaymentNotFound = errors.New("payment not found")

// CreatePayment creates a new payment in the database.
func (db *PaymentRepo) CreatePayment(ctx context.Context, req *pb.CreatePaymentRequest) (*pb.Payment, error) {
	id := uuid.New().String()
	query := `
		INSERT INTO 
			payments (id, reservation_id, amount, payment_method) 
        VALUES 
			($1, $2, $3, $4)
	`
	pay := pb.Payment{}
	var createdAt, updatedAt time.Time
	_, err := db.Db.Exec(query, id, req.Payment.ReservationId, req.Payment.Amount, req.Payment.PaymentMethod)
	if err != nil {
		slog.Error("Error while creating payment in postgres")
		return nil, err // Return the error
	}

	pay.CreatedAt = createdAt.Format(time.RFC3339)
	pay.UpdatedAt = updatedAt.Format(time.RFC3339)
	return &pay, nil
}

// UpdatePayment updates an existing payment.
func (db *PaymentRepo) UpdatePayment(ctx context.Context, req *pb.UpdatePaymentRequest) (*pb.Payment, error) {
	query := `
		UPDATE payments
		SET reservation_id = $1,
			amount = $2,
			payment_method = $3,
			payment_status = $4,
			updated_at = NOW()
		WHERE id = $5
	`
	_, err := db.Db.ExecContext(ctx, query, req.Payment.ReservationId, req.Payment.Amount, req.Payment.PaymentMethod, req.Payment.PaymentStatus, req.Payment.Id)
	if err != nil {
		slog.Error("Error while updating payment in postgres")
		return nil, err // Return the error
	}
	pay, err := db.GetPaymentById(ctx, &pb.GetPaymentRequest{Id: req.Payment.Id})
	if err != nil {
		slog.Error("Error getting payment after update")
		return nil, err
	}
	return pay, nil
}

// GetPaymentById retrieves a payment by its ID.
func (db *PaymentRepo) GetPaymentById(ctx context.Context, req *pb.GetPaymentRequest) (*pb.Payment, error) {
	query := `
		SELECT 
			id, 
			reservation_id,
			amount,
			payment_method,
			payment_status,
			created_at,
			updated_at
		FROM payments
		WHERE id = $1 AND deleted_at = 0 
	`
	pay := pb.Payment{}
	var createdAt, updatedAt time.Time
	err := db.Db.QueryRowContext(ctx, query, req.Id).Scan(
		&pay.Id,
		&pay.ReservationId, // Scan the reservation_id
		&pay.Amount,
		&pay.PaymentMethod,
		&pay.PaymentStatus,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrPaymentNotFound
		}
		slog.Error("Error while getting payment in postgres")
		return nil, err // Return the error
	}

	pay.CreatedAt = createdAt.Format(time.RFC3339)
	pay.UpdatedAt = updatedAt.Format(time.RFC3339)
	return &pay, nil
}

// DeletePayment soft deletes a payment by ID.
func (db *PaymentRepo) DeletePayment(ctx context.Context, req *pb.DeletePaymentRequest) (*pb.DeletePaymentResponse, error) {
	query := `
		UPDATE payments
		SET deleted_at = extract(epoch from NOW()) 
		WHERE id = $1 AND deleted_at = 0
	`
	_, err := db.Db.ExecContext(ctx, query, req.Id)
	if err != nil {
		slog.Error("Error while soft deleting payment")
		return nil, err
	}
	return &pb.DeletePaymentResponse{Message: "Payment soft deleted"}, nil
}

// PayForReservation updates the paid amount for a reservation,
// updates the payment status if fully paid, and returns the
// remaining balance or 0 if paid in full.
func (db *PaymentRepo) PayForReservation(ctx context.Context, req *pb.PayForReservationReq) (float64, error) {
	tx, err := db.Db.BeginTx(ctx, nil)
	if err != nil {
		slog.Error("Error starting transaction", "error", err)
		return 0, err
	}
	defer tx.Rollback() // Rollback if any error occurs

	// 1. Get the current payment details
	var currentPaidAmount, amount float64
	err = tx.QueryRowContext(ctx, "SELECT paid, amount FROM payments WHERE id = $1 FOR UPDATE", req.Id).Scan(&currentPaidAmount, &amount)
	if err != nil {
		slog.Error("Error getting current paid amount", "error", err)
		return 0, err
	}

	// 2. Calculate the new paid amount and remaining balance
	newPaidAmount := currentPaidAmount + req.Paid
	remainingBalance := amount - newPaidAmount

	// 3. Update the payment record
	var paymentStatus string
	if remainingBalance <= 0 {
		paymentStatus = "paid" // Set status to PAID
		remainingBalance = 0
	} else {
		paymentStatus = "unpaid" // Keep status UNPAID
	}

	_, err = tx.ExecContext(ctx,
		"UPDATE payments SET paid = $1, payment_status = $2, payment_method = $3 WHERE id = $4",
		newPaidAmount, paymentStatus, req.PaymentMethod, req.Id,
	)
	if err != nil {
		slog.Error("Error updating payment", "error", err)
		return 0, err
	}

	// 4. Commit the transaction
	if err := tx.Commit(); err != nil {
		slog.Error("Error committing transaction", "error", err)
		return 0, err
	}

	return remainingBalance, nil
}
