package postgres

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Project_Restaurant/Restaurant-Service/genproto/restaurant"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
)

// ErrRestaurantNotFound is returned when a restaurant is not found.
var ErrRestaurantNotFound = errors.New("restaurant not found")

// RestaurantDb provides database operations for restaurants.
type RestaurantDb struct {
	Db *pgx.Conn
}

// NewRestaurant creates a new instance of RestaurantDb.
func NewRestaurant(db *pgx.Conn) *RestaurantDb {
	return &RestaurantDb{Db: db}
}

// Create creates a new restaurant in the database.
func (rDb *RestaurantDb) Create(ctx context.Context, req *restaurant.CreateRestaurantRequest) (*restaurant.CreateRestaurantResponse, error) {
	restaurantID := uuid.New().String()
	query := `
		INSERT INTO 
			restaurants (
				id,
				name,
				address,
				phone_number,
				description
			) 
		VALUES (
				$1, 
				$2, 
				$3,
				$4,
				$5
			)
		RETURNING 
			id, 
			name,
			address,
			phone_number,
			description,
			created_at,
			updated_at
	`
	var (
		dbRestaurant restaurant.Restaurant
		createdAt    time.Time
		updatedAt    time.Time
	)

	err := rDb.Db.QueryRow(ctx, query,
		restaurantID,
		req.Restaurant.Name,
		req.Restaurant.Address,
		req.Restaurant.PhoneNumber,
		req.Restaurant.Description,
	).Scan(
		&dbRestaurant.Id,
		&dbRestaurant.Name,
		&dbRestaurant.Address,
		&dbRestaurant.PhoneNumber,
		&dbRestaurant.Description,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		log.Error().Err(err).Msg("Error creating restaurant")
		return nil, err
	}

	dbRestaurant.CreatedAt = createdAt.Format(time.RFC3339)
	dbRestaurant.UpdatedAt = updatedAt.Format(time.RFC3339)

	return &restaurant.CreateRestaurantResponse{Restaurant: &dbRestaurant}, nil
}

// GetById gets a restaurant by its ID.
func (rDb *RestaurantDb) GetById(ctx context.Context, req *restaurant.GetRestaurantRequest) (*restaurant.GetRestaurantResponse, error) {
	var (
		dbRestaurant restaurant.Restaurant
		createdAt    time.Time
		updatedAt    time.Time
	)

	query := `
		SELECT
			id,
			name,
			address,
			phone_number,
			description,
			created_at,
			updated_at
		FROM 
			restaurants 
		WHERE 
			id = $1
		AND 
			deleted_at = 0
	`
	err := rDb.Db.QueryRow(ctx, query, req.Id).Scan(
		&dbRestaurant.Id,
		&dbRestaurant.Name,
		&dbRestaurant.Address,
		&dbRestaurant.PhoneNumber,
		&dbRestaurant.Description,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			log.Error().Err(err).Msg("Restaurant not found")
			return nil, ErrRestaurantNotFound
		}
		log.Error().Err(err).Msg("Error getting restaurant by ID")
		return nil, err
	}

	dbRestaurant.CreatedAt = createdAt.Format(time.RFC3339)
	dbRestaurant.UpdatedAt = updatedAt.Format(time.RFC3339)

	return &restaurant.GetRestaurantResponse{Restaurant: &dbRestaurant}, nil
}

// Update updates an existing restaurant in the database.
func (rDb *RestaurantDb) Update(ctx context.Context, req *restaurant.UpdateRestaurantRequest) (*restaurant.UpdateRestaurantResponse, error) {
	var args []interface{}
	count := 1
	query := `
		UPDATE 
			restaurants 
		SET `
	filter := ``

	if len(req.Restaurant.Name) > 0 {
		filter += fmt.Sprintf(" name = $%d, ", count)
		args = append(args, req.Restaurant.Name)
		count++
	}

	if len(req.Restaurant.Address) > 0 {
		filter += fmt.Sprintf(" address = $%d, ", count)
		args = append(args, req.Restaurant.Address)
		count++
	}

	if len(req.Restaurant.PhoneNumber) > 0 {
		filter += fmt.Sprintf(" phone_number = $%d, ", count)
		args = append(args, req.Restaurant.PhoneNumber)
		count++
	}

	if len(req.Restaurant.Description) > 0 {
		filter += fmt.Sprintf(" description = $%d, ", count)
		args = append(args, req.Restaurant.Description)
		count++
	}

	if filter == "" {
		log.Error().Msg("No fields provided for update.")
		return nil, errors.New("no fields provided for update")
	}

	filter += fmt.Sprintf(`
			updated_at = NOW()
		WHERE
			id = $%d
			`, count)

	args = append(args, req.Restaurant.Id)
	query += filter

	_, err := rDb.Db.Exec(ctx, query, args...)
	if err != nil {
		if err == pgx.ErrNoRows {
			log.Error().Err(err).Msg("Restaurant not found")
			return nil, ErrRestaurantNotFound
		}
		log.Error().Err(err).Msg("Error updating restaurant")
		return nil, err
	}

	// After successful update, get the updated restaurant from the database
	updatedRestaurant, err := rDb.GetById(ctx, &restaurant.GetRestaurantRequest{Id: req.Restaurant.Id})
	if err != nil {
		return nil, err
	}

	return &restaurant.UpdateRestaurantResponse{Restaurant: updatedRestaurant.Restaurant}, nil
}

func (rDb *RestaurantDb) Delete(ctx context.Context, req *restaurant.DeleteRestaurantRequest) (*restaurant.DeleteRestaurantResponse, error) {
	query := `
		UPDATE 
			restaurants 
		SET 
			deleted_at = $1 
		WHERE 
			id = $2
	`
	_, err := rDb.Db.Exec(ctx, query, time.Now().Unix(), req.Id)
	if err != nil {
		if err == pgx.ErrNoRows {
			log.Error().Err(err).Msg("Restaurant not found")
			return nil, ErrRestaurantNotFound
		}
		log.Error().Err(err).Msg("Error deleting (soft) restaurant")
		return nil, err
	}
	return &restaurant.DeleteRestaurantResponse{Message: "Restaurant deleted (soft) successfully"}, nil
}

// ListRestaurants retrieves a list of restaurants with optional filtering and pagination.
func (rDb *RestaurantDb) ListRestaurants(ctx context.Context, req *restaurant.ListRestaurantsRequest) (*restaurant.ListRestaurantsResponse, error) {
	var args []interface{}
	count := 1
	query := `
		SELECT
			id,
			name,
			address,
			phone_number,
			description,
			created_at,
			updated_at
		FROM 
			restaurants
		WHERE 
			deleted_at = 0
	`
	filter := ""

	if req.Name != "" {
		filter += fmt.Sprintf(" AND name ILIKE $%d", count)
		args = append(args, "%"+req.Name+"%")
		count++
	}

	if req.Address != "" {
		filter += fmt.Sprintf(" AND address ILIKE $%d", count)
		args = append(args, "%"+req.Address+"%")
		count++
	}

	query += filter

	// Apply pagination
	if req.Limit <= 0 {
		req.Limit = 10 // Default limit
	}
	if req.Page <= 0 {
		req.Page = 1 // Default page
	}
	offset := (req.Page - 1) * req.Limit
	query += fmt.Sprintf(" OFFSET %d LIMIT %d", offset, req.Limit)

	rows, err := rDb.Db.Query(ctx, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("Error listing restaurants")
		return nil, err
	}
	defer rows.Close()

	var restaurants []*restaurant.Restaurant
	for rows.Next() {
		var (
			dbRestaurant restaurant.Restaurant
			createdAt    time.Time
			updatedAt    time.Time
		)
		err := rows.Scan(
			&dbRestaurant.Id,
			&dbRestaurant.Name,
			&dbRestaurant.Address,
			&dbRestaurant.PhoneNumber,
			&dbRestaurant.Description,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			log.Error().Err(err).Msg("Error scanning restaurant row")
			return nil, err
		}
		dbRestaurant.CreatedAt = createdAt.Format(time.RFC3339)
		dbRestaurant.UpdatedAt = updatedAt.Format(time.RFC3339)

		restaurants = append(restaurants, &dbRestaurant)
	}

	if err = rows.Err(); err != nil {
		log.Error().Err(err).Msg("Error iterating over restaurant rows")
		return nil, err
	}

	return &restaurant.ListRestaurantsResponse{Restaurants: restaurants}, nil
}
