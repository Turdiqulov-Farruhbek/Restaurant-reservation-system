package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	pb "github.com/Project_Restaurant/Payment-Service/genproto/payment"
	"github.com/Project_Restaurant/Payment-Service/storage/postgres"
	"github.com/google/uuid"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

const (
	dbHost     = "localhost"
	dbPort     = 5432
	dbUser     = "postgres"
	dbPassword = "root"
	dbName     = "payment_db"
)

func newTestDB(t *testing.T) *sql.DB {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		t.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		t.Fatal(err)
	}

	return db
}

func TestPaymentRepo(t *testing.T) {
	db := newTestDB(t)
	defer db.Close()
	repo := postgres.NewPaymentRepo(db)

	t.Run("CreatePayment", func(t *testing.T) {
		req := &pb.CreatePaymentRequest{
			Payment: &pb.Payment{
				ReservationId: uuid.NewString(),
				Amount:        100.00,
				PaymentMethod: "card",
				PaymentStatus: "unpaid",
			},
		}
		payment, err := repo.CreatePayment(context.Background(), req)
		if err != nil {
			t.Fatal(err)
		}
		assert.NotEmpty(t, payment.Id)
		assert.Equal(t, req.Payment.ReservationId, payment.ReservationId)
		assert.Equal(t, req.Payment.Amount, payment.Amount)
		assert.Equal(t, req.Payment.PaymentMethod, payment.PaymentMethod)
	})

	t.Run("GetPaymentById", func(t *testing.T) {
		createdPayment, err := repo.CreatePayment(context.Background(), &pb.CreatePaymentRequest{
			Payment: &pb.Payment{
				ReservationId: uuid.NewString(),
				Amount:        150.50,
				PaymentMethod: "paypal",
				PaymentStatus: "paid",
			},
		})
		if err != nil {
			t.Fatal(err)
		}

		payment, err := repo.GetPaymentById(context.Background(), &pb.GetPaymentRequest{Id: createdPayment.Id})
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, createdPayment.Id, payment.Id)
		assert.Equal(t, createdPayment.ReservationId, payment.ReservationId)
		assert.Equal(t, createdPayment.Amount, payment.Amount)
		assert.Equal(t, createdPayment.PaymentMethod, payment.PaymentMethod)
		assert.Equal(t, createdPayment.PaymentStatus, payment.PaymentStatus)
		// Add assertions for CreatedAt and UpdatedAt
	})

	t.Run("UpdatePayment", func(t *testing.T) {
		createdPayment, err := repo.CreatePayment(context.Background(), &pb.CreatePaymentRequest{
			Payment: &pb.Payment{
				ReservationId: uuid.NewString(),
				Amount:        200.00,
				PaymentMethod: "cash",
				PaymentStatus: "unpaid",
			},
		})
		if err != nil {
			t.Fatal(err)
		}

		updateReq := &pb.UpdatePaymentRequest{
			Payment: &pb.Payment{
				Id:            createdPayment.Id,
				ReservationId: uuid.NewString(),
				Amount:        250.00,
				PaymentMethod: "credit card",
				PaymentStatus: "paid",
			},
		}

		updatedPayment, err := repo.UpdatePayment(context.Background(), updateReq)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, updateReq.Payment.Id, updatedPayment.Id)
		assert.Equal(t, updateReq.Payment.ReservationId, updatedPayment.ReservationId)
		assert.Equal(t, updateReq.Payment.Amount, updatedPayment.Amount)
		assert.Equal(t, updateReq.Payment.PaymentMethod, updatedPayment.PaymentMethod)
		assert.Equal(t, updateReq.Payment.PaymentStatus, updatedPayment.PaymentStatus)
		// Add assertions for UpdatedAt
	})

	t.Run("DeletePayment", func(t *testing.T) {
		createdPayment, err := repo.CreatePayment(context.Background(), &pb.CreatePaymentRequest{
			Payment: &pb.Payment{
				ReservationId: uuid.NewString(),
				Amount:        300.00,
				PaymentMethod: "bank transfer",
				PaymentStatus: "paid",
			},
		})
		if err != nil {
			t.Fatal(err)
		}

		_, err = repo.DeletePayment(context.Background(), &pb.DeletePaymentRequest{Id: createdPayment.Id})
		if err != nil {
			t.Fatal(err)
		}

		_, err = repo.GetPaymentById(context.Background(), &pb.GetPaymentRequest{Id: createdPayment.Id})
		assert.NotNil(t, err) // Expect an error since the payment should be deleted
	})
}

func TestPaymentRepo_PayForReservation(t *testing.T) {
	db := newTestDB(t)
	defer db.Close()
	repo := postgres.NewPaymentRepo(db)

	t.Run("FullyPayReservation", func(t *testing.T) {
		// 1. Create a payment (unpaid status by default)
		createdPayment, err := repo.CreatePayment(context.Background(), &pb.CreatePaymentRequest{
			Payment: &pb.Payment{
				ReservationId: uuid.NewString(),
				Amount:        100.00,
				PaymentMethod: "card",
				PaymentStatus: "unpaid", // You might not need this if it's the default
			},
		})
		if err != nil {
			t.Fatal(err)
		}

		// 2. Pay the full amount
		remainingBalance, err := repo.PayForReservation(context.Background(), &pb.PayForReservationReq{
			Id:   createdPayment.Id,
			Paid: 100.00,
		})
		if err != nil {
			t.Fatal(err)
		}

		// 3. Assertions
		assert.Equal(t, 0.0, remainingBalance)

		// 4. Retrieve the updated payment and check the status
		updatedPayment, err := repo.GetPaymentById(context.Background(), &pb.GetPaymentRequest{Id: createdPayment.Id})
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, "paid", updatedPayment.PaymentStatus)
	})

	t.Run("PartiallyPayReservation", func(t *testing.T) {
		createdPayment, err := repo.CreatePayment(context.Background(), &pb.CreatePaymentRequest{
			Payment: &pb.Payment{
				ReservationId: uuid.NewString(),
				Amount:        200.00,
				PaymentMethod: "paypal",
				PaymentStatus: "unpaid",
			},
		})
		if err != nil {
			t.Fatal(err)
		}

		remainingBalance, err := repo.PayForReservation(context.Background(), &pb.PayForReservationReq{
			Id:   createdPayment.Id,
			Paid: 50.00,
		})
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, 150.0, remainingBalance)

		updatedPayment, err := repo.GetPaymentById(context.Background(), &pb.GetPaymentRequest{Id: createdPayment.Id})
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, "unpaid", updatedPayment.PaymentStatus)
	})
}
