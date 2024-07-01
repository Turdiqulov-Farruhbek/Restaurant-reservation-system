package test

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/Project_Restaurant/Restaurant-Service/genproto/reservation"
	"github.com/Project_Restaurant/Restaurant-Service/storage/postgres"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
)

func newTestReservation(t *testing.T) *postgres.ReservationDb {
	connString := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
		"sayyidmuhammad",
		"root",
		"localhost",
		5432,
		"reservation",
	)

	db, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	return &postgres.ReservationDb{Db: db}
}

func createTestReservation() *reservation.CreateReservationRequest {
	reservationTime := time.Now().Format(time.RFC3339) // Example: Reservation for tomorrow
	return &reservation.CreateReservationRequest{
		Reservation: &reservation.Reservation{
			UserId:          uuid.NewString(),
			RestaurantId:    "345e1148-678e-4ec7-9fab-d55df1e9cb54",
			ReservationTime: reservationTime,
			Status:          "PENDING",
		},
	}
}

func TestCreateReservation(t *testing.T) {
	rDb := newTestReservation(t)
	testReservation := createTestReservation()

	createdReservation, err := rDb.Create(context.Background(), testReservation)
	if err != nil {
		t.Fatalf("Error creating reservation: %v", err)
	}

	assert.NotEmpty(t, createdReservation.Reservation.Id)
	assert.Equal(t, testReservation.Reservation.UserId, createdReservation.Reservation.UserId)
	assert.Equal(t, testReservation.Reservation.RestaurantId, createdReservation.Reservation.RestaurantId)
	assert.Equal(t, testReservation.Reservation.ReservationTime, createdReservation.Reservation.ReservationTime)
	assert.Equal(t, testReservation.Reservation.Status, createdReservation.Reservation.Status)
}

func TestGetReservationById(t *testing.T) {
	rDb := newTestReservation(t)
	testReservation := createTestReservation()

	createdReservation, err := rDb.Create(context.Background(), testReservation)
	if err != nil {
		t.Fatalf("Error creating reservation: %v", err)
	}

	getReservation, err := rDb.GetById(context.Background(), &reservation.GetReservationRequest{Id: createdReservation.Reservation.Id})
	if err != nil {
		t.Fatalf("Error getting reservation by ID: %v", err)
	}

	assert.Equal(t, createdReservation.Reservation.Id, getReservation.Reservation.Id)
	assert.Equal(t, createdReservation.Reservation.UserId, getReservation.Reservation.UserId)
	assert.Equal(t, createdReservation.Reservation.RestaurantId, getReservation.Reservation.RestaurantId)
	assert.Equal(t, createdReservation.Reservation.ReservationTime, getReservation.Reservation.ReservationTime)
	assert.Equal(t, createdReservation.Reservation.Status, getReservation.Reservation.Status)
}

func TestUpdateReservation(t *testing.T) {
	rDb := newTestReservation(t)
	testReservation := createTestReservation()

	createdReservation, err := rDb.Create(context.Background(), testReservation)
	if err != nil {
		t.Fatalf("Error creating reservation: %v", err)
	}

	updatedReservationReq := &reservation.UpdateReservationRequest{
		Reservation: &reservation.Reservation{
			Id:              createdReservation.Reservation.Id,
			UserId:          uuid.NewString(), // Change the user ID
			RestaurantId:    createdReservation.Reservation.RestaurantId,
			ReservationTime: createdReservation.Reservation.ReservationTime,
			Status:          "CONFIRMED", // Change the status
		},
	}

	updatedReservation, err := rDb.Update(context.Background(), updatedReservationReq)
	if err != nil {
		t.Fatalf("Error updating reservation: %v", err)
	}

	assert.Equal(t, updatedReservationReq.Reservation.Id, updatedReservation.Reservation.Id)
	assert.Equal(t, updatedReservationReq.Reservation.UserId, updatedReservation.Reservation.UserId)
	assert.Equal(t, updatedReservationReq.Reservation.Status, updatedReservation.Reservation.Status)
	// ... other assertions
}

func TestDeleteReservation(t *testing.T) {
	rDb := newTestReservation(t)
	testReservation := createTestReservation()

	createdReservation, err := rDb.Create(context.Background(), testReservation)
	if err != nil {
		t.Fatalf("Error creating reservation: %v", err)
	}

	_, err = rDb.Delete(context.Background(), &reservation.DeleteReservationRequest{Id: createdReservation.Reservation.Id})
	if err != nil {
		t.Fatalf("Error deleting reservation: %v", err)
	}

	_, err = rDb.GetById(context.Background(), &reservation.GetReservationRequest{Id: createdReservation.Reservation.Id})
	assert.ErrorIs(t, err, postgres.ErrReservationNotFound)
}

func TestListReservations(t *testing.T) {
	rDb := newTestReservation(t)

	// Create test reservations
	restaurantID1 := "345e1148-678e-4ec7-9fab-d55df1e9cb54"
	userID1 := uuid.NewString()
	reservationTime1 := time.Now().Add(24 * time.Hour).Format(time.RFC3339)
	reservationTime2 := time.Now().Add(48 * time.Hour).Format(time.RFC3339)

	testReservations := []*reservation.CreateReservationRequest{
		{Reservation: &reservation.Reservation{UserId: userID1, RestaurantId: restaurantID1, ReservationTime: reservationTime1, Status: "PENDING"}},
		{Reservation: &reservation.Reservation{UserId: userID1, RestaurantId: restaurantID1, ReservationTime: reservationTime2, Status: "CONFIRMED"}},
		{Reservation: &reservation.Reservation{UserId: uuid.NewString(), RestaurantId: "60e83378-4e69-4e96-b24d-87c753363725", ReservationTime: reservationTime1, Status: "PENDING"}},
	}

	for _, tr := range testReservations {
		_, err := rDb.Create(context.Background(), tr)
		if err != nil {
			t.Fatalf("Error creating test reservation: %v", err)
		}
	}

	t.Run("ListReservations without filters", func(t *testing.T) {
		resp, err := rDb.ListReservations(context.Background(), &reservation.ListReservationsRequest{})
		if err != nil {
			t.Fatalf("Error listing reservations: %v", err)
		}
		assert.GreaterOrEqual(t, len(resp.Reservations), len(testReservations))
	})

	t.Run("Filter by user ID", func(t *testing.T) {
		resp, err := rDb.ListReservations(context.Background(), &reservation.ListReservationsRequest{UserId: userID1})
		if err != nil {
			t.Fatalf("Error listing reservations by user ID: %v", err)
		}
		assert.Equal(t, 2, len(resp.Reservations))
	})

	t.Run("Filter by restaurant ID", func(t *testing.T) {
		resp, err := rDb.ListReservations(context.Background(), &reservation.ListReservationsRequest{RestaurantId: restaurantID1})
		if err != nil {
			t.Fatalf("Error listing reservations by restaurant ID: %v", err)
		}
		assert.LessOrEqual(t, 2, len(resp.Reservations))
	})
}

func TestCheckAvailability_AvailableBeforeExistingReservation(t *testing.T) {
	rDb := newTestReservation(t)

	// Create a test reservation
	restaurantID := "e547d0d3-a804-4dba-921a-106b666c304b"
	userID := uuid.NewString()
	reservationTime := time.Now().Add(24 * time.Hour).Format(time.RFC3339)

	_, err := rDb.Create(context.Background(), &reservation.CreateReservationRequest{
		Reservation: &reservation.Reservation{
			UserId:          userID,
			RestaurantId:    restaurantID,
			ReservationTime: reservationTime,
			Status:          "PENDING",
		},
	})
	if err != nil {
		t.Fatalf("Error creating test reservation: %v", err)
	}

	// Test an available time slot one hour before the existing reservation
	reqTime := time.Now().Add(22 * time.Hour).Format(time.RFC3339)
	resp, err := rDb.CheckAvailability(context.Background(), &reservation.CheckAvailabilityRequest{
		RestaurantId:    restaurantID,
		ReservationTime: reqTime,
	})

	assert.NoError(t, err)
	log.Println(resp.Available)
	assert.True(t, resp.Available, "Time slot should be available")
}

func TestOrderFoodList(t *testing.T) {
	rDb := newTestReservation(t)
	reservationID := "f92552af-8e4f-4101-95ca-fedc3c994b7d"

	// Create a test reservation for which we'll list foods
	_, err := rDb.Create(context.Background(), &reservation.CreateReservationRequest{
		Reservation: &reservation.Reservation{
			UserId:          uuid.NewString(),
			RestaurantId:    "345e1148-678e-4ec7-9fab-d55df1e9cb54",
			ReservationTime: time.Now().Add(24 * time.Hour).Format(time.RFC3339),
			Status:          "PENDING",
		},
	})
	if err != nil {
		t.Fatalf("Error creating test reservation: %v", err)
	}

	// Create some test menus
	menus := []*reservation.Menus{
		{
			Id:           uuid.NewString(),
			RestaurantId: "345e1148-678e-4ec7-9fab-d55df1e9cb54",
			Name:         "Pizza Margherita",
			Description:  "Classic Margherita pizza with tomato sauce, mozzarella, and basil",
			Price:        12.99,
		},
		{
			Id:           uuid.NewString(),
			RestaurantId: "345e1148-678e-4ec7-9fab-d55df1e9cb54",
			Name:         "Pasta Carbonara",
			Description:  "Creamy carbonara pasta with pancetta, eggs, and parmesan cheese",
			Price:        15.99,
		},
	}

	for _, m := range menus {
		_, err := rDb.Db.Exec(context.Background(), `
			INSERT INTO menus (
				id,
				restaurant_id,
				name,
				description,
				price
			) VALUES ($1, $2, $3, $4, $5)
		`, m.Id, m.RestaurantId, m.Name, m.Description, m.Price)
		if err != nil {
			t.Fatalf("Error creating test menu: %v", err)
		}
	}

	// Get the list of foods for the reservation
	foodListResp, err := rDb.FoodList(context.Background(), &reservation.OrderFoodListReq{ReservationId: reservationID})
	if err != nil {
		t.Fatalf("Error getting food list: %v", err)
	}

	// Assert that the food list is not empty and contains the correct menus
	assert.NotEmpty(t, foodListResp.Menus)
	assert.GreaterOrEqual(t, len(foodListResp.Menus), len(menus))
	log.Println(foodListResp.Menus)
	found := false
	for _, m := range foodListResp.Menus {
		// Assert that the returned menu ID exists in the created menus list
		for _, tm := range menus {
			if tm.Id == m.Id {
				found = true
				break
			}
		}
	}
	assert.True(t, found, "Returned menu ID should be in the created menus list")
}

func TestOrderFood(t *testing.T) {
	rDb := newTestReservation(t)

	// Create a test reservation
	reservationID := "f92552af-8e4f-4101-95ca-fedc3c994b7d"
	_, err := rDb.Create(context.Background(), &reservation.CreateReservationRequest{
		Reservation: &reservation.Reservation{
			UserId:          uuid.NewString(),
			RestaurantId:    "345e1148-678e-4ec7-9fab-d55df1e9cb54",
			ReservationTime: time.Now().Add(24 * time.Hour).Format(time.RFC3339),
			Status:          "PENDING",
		},
	})
	if err != nil {
		t.Fatalf("Error creating test reservation: %v", err)
	}

	// Create a test menu
	menuID := uuid.NewString()
	_, err = rDb.Db.Exec(context.Background(), `
		INSERT INTO menus (
			id,
			restaurant_id,
			name,
			description,
			price
		) VALUES ($1, $2, $3, $4, $5)
	`, menuID, "345e1148-678e-4ec7-9fab-d55df1e9cb54", "Pizza Margherita", "Classic Margherita pizza", 12.99)
	if err != nil {
		t.Fatalf("Error creating test menu: %v", err)
	}

	// Order the menu item
	orderResp, err := rDb.OrderFood(context.Background(), &reservation.OrderFoodReq{
		ReservationId: reservationID,
		MenuId:        menuID,
		Count:         2,
	})
	if err != nil {
		t.Fatalf("Error ordering food: %v", err)
	}

	assert.Equal(t, "Order created succesfully", orderResp.Message)

	// Verify the order is created in the database
	var orderCount int
	err = rDb.Db.QueryRow(context.Background(), `
		SELECT COUNT(*)
		FROM orders
		WHERE reservation_id = $1 AND menu_id = $2
	`, reservationID, menuID).Scan(&orderCount)
	if err != nil {
		t.Fatalf("Error checking order count: %v", err)
	}

	assert.Equal(t, 1, orderCount, "Order should be created in the database")
}

func TestIsValid(t *testing.T) {
	rDb := newTestReservation(t)

	t.Run("CreateReservation and check validity", func(t *testing.T) {
		testReservation := createTestReservation()

		createdReservation, err := rDb.Create(context.Background(), testReservation)
		if err != nil {
			t.Fatalf("Error creating reservation: %v", err)
		}

		isValidResp, err := rDb.IsValidReservation(context.Background(), &reservation.IsValidReq{Id: createdReservation.Reservation.Id})
		if err != nil {
			t.Fatalf("Error checking reservation validity: %v", err)
		}

		assert.True(t, isValidResp.Valid, "Newly created reservation should be valid")
	})

	t.Run("DeleteReservation and check validity", func(t *testing.T) {
		// Create a temporary reservation to delete
		tempReservation := createTestReservation()
		tempReservation.Reservation.UserId = uuid.NewString() // Use a unique user ID for the temp reservation

		createdReservation, err := rDb.Create(context.Background(), tempReservation)
		if err != nil {
			t.Fatalf("Error creating temporary reservation: %v", err)
		}

		// Delete the reservation
		_, err = rDb.Delete(context.Background(), &reservation.DeleteReservationRequest{Id: createdReservation.Reservation.Id})
		if err != nil {
			t.Fatalf("Error deleting reservation: %v", err)
		}

		// Check if the deleted reservation is no longer valid
		isValidResp, err := rDb.IsValidReservation(context.Background(), &reservation.IsValidReq{Id: createdReservation.Reservation.Id})
		if err != nil {
			t.Fatalf("Error checking reservation validity: %v", err)
		}

		assert.False(t, isValidResp.Valid, "Deleted reservation should be invalid")
	})
}
