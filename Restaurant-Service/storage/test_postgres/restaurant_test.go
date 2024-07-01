package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/Project_Restaurant/Restaurant-Service/genproto/restaurant"
	"github.com/Project_Restaurant/Restaurant-Service/storage/postgres"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
)

func newTestRestaurant(t *testing.T) *postgres.RestaurantDb {
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
	return &postgres.RestaurantDb{Db: db}
}

func createTestRestaurant() *restaurant.CreateRestaurantRequest {
	return &restaurant.CreateRestaurantRequest{
		Restaurant: &restaurant.Restaurant{
			Name:        "Test Restaurant",
			Address:     "123 Test St",
			PhoneNumber: "+1234567890",
			Description: "A test restaurant",
		},
	}
}

func TestCreateRestaurant(t *testing.T) {
	rDb := newTestRestaurant(t)
	testRestaurant := createTestRestaurant()

	createdRestaurant, err := rDb.Create(context.Background(), testRestaurant)
	if err != nil {
		t.Fatalf("Error creating restaurant: %v", err)
	}

	assert.NotEmpty(t, createdRestaurant.Restaurant.Id)
	assert.Equal(t, testRestaurant.Restaurant.Name, createdRestaurant.Restaurant.Name)
	assert.Equal(t, testRestaurant.Restaurant.Address, createdRestaurant.Restaurant.Address)
	assert.Equal(t, testRestaurant.Restaurant.PhoneNumber, createdRestaurant.Restaurant.PhoneNumber)
	assert.Equal(t, testRestaurant.Restaurant.Description, createdRestaurant.Restaurant.Description)
}

func TestGetRestaurantById(t *testing.T) {
	rDb := newTestRestaurant(t)
	testRestaurant := createTestRestaurant()

	createdRestaurant, err := rDb.Create(context.Background(), testRestaurant)
	if err != nil {
		t.Fatalf("Error creating restaurant: %v", err)
	}

	getRestaurant, err := rDb.GetById(context.Background(), &restaurant.GetRestaurantRequest{Id: createdRestaurant.Restaurant.Id})
	if err != nil {
		t.Fatalf("Error getting restaurant by ID: %v", err)
	}

	assert.Equal(t, createdRestaurant.Restaurant.Id, getRestaurant.Restaurant.Id)
	assert.Equal(t, createdRestaurant.Restaurant.Name, getRestaurant.Restaurant.Name)
	// ... other assertions for address, phone number, description
}

func TestUpdateRestaurant(t *testing.T) {
	rDb := newTestRestaurant(t)
	testRestaurant := createTestRestaurant()

	createdRestaurant, err := rDb.Create(context.Background(), testRestaurant)
	if err != nil {
		t.Fatalf("Error creating restaurant: %v", err)
	}

	updatedRestaurantReq := &restaurant.UpdateRestaurantRequest{
		Restaurant: &restaurant.Restaurant{
			Id:          createdRestaurant.Restaurant.Id,
			Name:        "Updated Restaurant Name",
			Address:     "456 Updated St",
			PhoneNumber: "+9876543210",
			Description: "An updated restaurant",
		},
	}

	updatedRestaurant, err := rDb.Update(context.Background(), updatedRestaurantReq)
	if err != nil {
		t.Fatalf("Error updating restaurant: %v", err)
	}

	assert.Equal(t, updatedRestaurantReq.Restaurant.Id, updatedRestaurant.Restaurant.Id)
	assert.Equal(t, updatedRestaurantReq.Restaurant.Name, updatedRestaurant.Restaurant.Name)
	// ... other assertions for address, phone number, description
}

func TestDeleteRestaurant(t *testing.T) {
	rDb := newTestRestaurant(t)
	testRestaurant := createTestRestaurant()

	createdRestaurant, err := rDb.Create(context.Background(), testRestaurant)
	if err != nil {
		t.Fatalf("Error creating restaurant: %v", err)
	}

	_, err = rDb.Delete(context.Background(), &restaurant.DeleteRestaurantRequest{Id: createdRestaurant.Restaurant.Id})
	if err != nil {
		t.Fatalf("Error deleting restaurant: %v", err)
	}

	// Try to get the deleted restaurant
	deletedRestaurant, err := rDb.GetById(context.Background(), &restaurant.GetRestaurantRequest{Id: createdRestaurant.Restaurant.Id})
	assert.ErrorIs(t, err, postgres.ErrRestaurantNotFound)
	assert.Nil(t, deletedRestaurant)
}

func TestListRestaurants(t *testing.T) {
	rDb := newTestRestaurant(t)

	// Create some test restaurants
	restaurantID1 := uuid.NewString()
	restaurantID2 := uuid.NewString()

	testRestaurants := []*restaurant.CreateRestaurantRequest{
		{Restaurant: &restaurant.Restaurant{Id: restaurantID1, Name: "Restaurant A", Address: "123 Main St", PhoneNumber: "+1234567890", Description: "Delicious food"}},
		{Restaurant: &restaurant.Restaurant{Id: restaurantID2, Name: "Cafe B", Address: "456 Oak Ave", PhoneNumber: "+9876543210", Description: "Cozy cafe"}},
		{Restaurant: &restaurant.Restaurant{Name: "Bistro C", Address: "789 Pine Ln", PhoneNumber: "+1122334455", Description: "Fine dining"}},
	}

	for _, tr := range testRestaurants {
		_, err := rDb.Create(context.Background(), tr)
		if err != nil {
			t.Fatalf("Error creating test restaurant: %v", err)
		}
	}

	t.Run("ListRestaurants without filters", func(t *testing.T) {
		resp, err := rDb.ListRestaurants(context.Background(), &restaurant.ListRestaurantsRequest{})
		if err != nil {
			t.Fatalf("Error listing all restaurants: %v", err)
		}
		assert.GreaterOrEqual(t, len(resp.Restaurants), len(testRestaurants))
	})

	t.Run("Filter by name", func(t *testing.T) {
		resp, err := rDb.ListRestaurants(context.Background(), &restaurant.ListRestaurantsRequest{Name: "B"})
		if err != nil {
			t.Fatalf("Error listing restaurants by name: %v", err)
		}
		assert.LessOrEqual(t, 2, len(resp.Restaurants))
	})

	t.Run("Filter by address", func(t *testing.T) {
		resp, err := rDb.ListRestaurants(context.Background(), &restaurant.ListRestaurantsRequest{Address: "Main"})
		if err != nil {
			t.Fatalf("Error listing restaurants by address: %v", err)
		}
		assert.LessOrEqual(t, 1, len(resp.Restaurants))
	})

	// Add more test cases as needed for other filter and pagination scenarios
}
