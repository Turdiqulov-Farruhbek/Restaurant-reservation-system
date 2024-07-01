package postgres

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Project_Restaurant/Restaurant-Service/genproto/reservation"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
)

// ErrReservationNotFound is returned when a reservation is not found.
var ErrReservationNotFound = errors.New("reservation not found")

// ReservationDb provides database operations for reservations.
type ReservationDb struct {
	Db *pgx.Conn
}

// NewReservation creates a new instance of ReservationDb.
func NewReservation(db *pgx.Conn) *ReservationDb {
	return &ReservationDb{Db: db}
}

// Create creates a new reservation in the database.
func (rDb *ReservationDb) Create(ctx context.Context, req *reservation.CreateReservationRequest) (*reservation.CreateReservationResponse, error) {

	reservationID := uuid.New().String()
	query := `
		INSERT INTO 
			reservations (
				id,
				user_id,
				restaurant_id,
				reservation_time,
				status
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
			user_id,
			restaurant_id,
			reservation_time,
			status,
			created_at,
			updated_at
	`
	var (
		dbReservation reservation.Reservation
		createdAt     time.Time
		updatedAt     time.Time
	)

	reservationTime, err := time.Parse(time.RFC3339, req.Reservation.ReservationTime)
	if err != nil {
		return nil, fmt.Errorf("invalid reservation time format: %w", err)
	}
	log.Logger.Println(reservationTime)
	err = rDb.Db.QueryRow(ctx, query,
		reservationID,
		req.Reservation.UserId,
		req.Reservation.RestaurantId,
		reservationTime,
		req.Reservation.Status,
	).Scan(
		&dbReservation.Id,
		&dbReservation.UserId,
		&dbReservation.RestaurantId,
		&reservationTime,
		&dbReservation.Status,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		log.Error().Err(err).Msg("Error creating reservation")
		return nil, err
	}

	dbReservation.CreatedAt = createdAt.Format(time.RFC3339)
	dbReservation.UpdatedAt = updatedAt.Format(time.RFC3339)
	dbReservation.ReservationTime = reservationTime.Format(time.RFC3339)

	return &reservation.CreateReservationResponse{Reservation: &dbReservation}, nil
}

// GetById gets a reservation by its ID.
func (rDb *ReservationDb) GetById(ctx context.Context, req *reservation.GetReservationRequest) (*reservation.GetReservationResponse, error) {
	var (
		dbReservation   reservation.Reservation
		createdAt       time.Time
		updatedAt       time.Time
		reservationTime time.Time
	)

	query := `
		SELECT
			id,
			user_id,
			restaurant_id,
			reservation_time,
			status,
			created_at,
			updated_at
		FROM 
			reservations 
		WHERE 
			id = $1
	`
	err := rDb.Db.QueryRow(ctx, query, req.Id).Scan(
		&dbReservation.Id,
		&dbReservation.UserId,
		&dbReservation.RestaurantId,
		&reservationTime,
		&dbReservation.Status,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			log.Error().Err(err).Msg("Reservation not found")
			return nil, ErrReservationNotFound
		}
		log.Error().Err(err).Msg("Error getting reservation by ID")
		return nil, err
	}

	dbReservation.CreatedAt = createdAt.Format(time.RFC3339)
	dbReservation.UpdatedAt = updatedAt.Format(time.RFC3339)
	dbReservation.ReservationTime = reservationTime.Format(time.RFC3339)

	return &reservation.GetReservationResponse{Reservation: &dbReservation}, nil
}

// Update updates an existing reservation in the database.
func (rDb *ReservationDb) Update(ctx context.Context, req *reservation.UpdateReservationRequest) (*reservation.UpdateReservationResponse, error) {
	var args []interface{}
	count := 1
	query := `
		UPDATE 
			reservations 
		SET `
	filter := ``

	if len(req.Reservation.UserId) > 0 {
		filter += fmt.Sprintf(" user_id = $%d, ", count)
		args = append(args, req.Reservation.UserId)
		count++
	}

	if len(req.Reservation.RestaurantId) > 0 {
		filter += fmt.Sprintf(" restaurant_id = $%d, ", count)
		args = append(args, req.Reservation.RestaurantId)
		count++
	}

	if len(req.Reservation.ReservationTime) > 0 {
		reservationTime, err := time.Parse(time.RFC3339, req.Reservation.ReservationTime)
		if err != nil {
			return nil, fmt.Errorf("invalid reservation time format: %w", err)
		}
		filter += fmt.Sprintf(" reservation_time = $%d, ", count)
		args = append(args, reservationTime)
		count++
	}

	if len(req.Reservation.Status) > 0 {
		filter += fmt.Sprintf(" status = $%d, ", count)
		args = append(args, req.Reservation.Status)
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

	args = append(args, req.Reservation.Id)
	query += filter

	_, err := rDb.Db.Exec(ctx, query, args...)
	if err != nil {
		if err == pgx.ErrNoRows {
			log.Error().Err(err).Msg("Reservation not found")
			return nil, ErrReservationNotFound
		}
		log.Error().Err(err).Msg("Error updating reservation")
		return nil, err
	}

	// After successful update, get the updated reservation from the database
	updatedReservation, err := rDb.GetById(ctx, &reservation.GetReservationRequest{Id: req.Reservation.Id})
	if err != nil {
		return nil, err
	}

	return &reservation.UpdateReservationResponse{Reservation: updatedReservation.Reservation}, nil
}

// Delete deletes a reservation by its ID.
func (rDb *ReservationDb) Delete(ctx context.Context, req *reservation.DeleteReservationRequest) (*reservation.DeleteReservationResponse, error) {
	query := `
		DELETE FROM 
			reservations 
		WHERE 
			id = $1
	`
	_, err := rDb.Db.Exec(ctx, query, req.Id)
	if err != nil {
		if err == pgx.ErrNoRows {
			log.Error().Err(err).Msg("Reservation not found")
			return nil, ErrReservationNotFound
		}
		log.Error().Err(err).Msg("Error deleting reservation")
		return nil, err
	}
	return &reservation.DeleteReservationResponse{Message: "Restaurant deleted successfully"}, nil
}

// ListReservations retrieves a list of reservations with filtering and pagination.
func (rDb *ReservationDb) ListReservations(ctx context.Context, req *reservation.ListReservationsRequest) (*reservation.ListReservationsResponse, error) {
	var args []interface{}
	count := 1
	query := `
		SELECT
			id,
			user_id,
			restaurant_id,
			reservation_time,
			status,
			created_at,
			updated_at
		FROM 
			reservations
		WHERE 1=1 
	`

	filter := ""

	if req.UserId != "" {
		filter += fmt.Sprintf(" AND user_id = $%d", count)
		args = append(args, req.UserId)
		count++
	}

	if req.RestaurantId != "" {
		filter += fmt.Sprintf(" AND restaurant_id = $%d", count)
		args = append(args, req.RestaurantId)
		count++
	}

	if req.Status != "" {
		filter += fmt.Sprintf(" AND status = $%d", count)
		args = append(args, req.Status)
		count++
	}

	if req.StartTime != "" {
		startTime, err := time.Parse(time.RFC3339, req.StartTime)
		if err != nil {
			return nil, fmt.Errorf("invalid start time format: %w", err)
		}
		filter += fmt.Sprintf(" AND reservation_time >= $%d", count)
		args = append(args, startTime)
		count++
	}

	if req.EndTime != "" {
		endTime, err := time.Parse(time.RFC3339, req.EndTime)
		if err != nil {
			return nil, fmt.Errorf("invalid end time format: %w", err)
		}
		filter += fmt.Sprintf(" AND reservation_time <= $%d", count)
		args = append(args, endTime)
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
		log.Error().Err(err).Msg("Error listing reservations")
		return nil, err
	}
	defer rows.Close()

	var reservations []*reservation.Reservation
	for rows.Next() {
		var (
			dbReservation   reservation.Reservation
			createdAt       time.Time
			updatedAt       time.Time
			reservationTime time.Time
		)
		err := rows.Scan(
			&dbReservation.Id,
			&dbReservation.UserId,
			&dbReservation.RestaurantId,
			&reservationTime,
			&dbReservation.Status,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			log.Error().Err(err).Msg("Error scanning reservation row")
			return nil, err
		}
		dbReservation.CreatedAt = createdAt.Format(time.RFC3339)
		dbReservation.UpdatedAt = updatedAt.Format(time.RFC3339)
		dbReservation.ReservationTime = reservationTime.Format(time.RFC3339)
		reservations = append(reservations, &dbReservation)
	}

	if err = rows.Err(); err != nil {
		log.Error().Err(err).Msg("Error iterating over reservation rows")
		return nil, err
	}

	return &reservation.ListReservationsResponse{Reservations: reservations}, nil
}

// CheckAvailability checks if a reservation is available for a given restaurant and time.
func (rDb *ReservationDb) CheckAvailability(ctx context.Context, req *reservation.CheckAvailabilityRequest) (*reservation.CheckAvailabilityResponse, error) {
	reservationTime, err := time.Parse(time.RFC3339, req.ReservationTime)
	if err != nil {
		return nil, fmt.Errorf("invalid reservation time format: %w", err)
	}

	// Calculate the end time of the reservation (reservation time + 1 hour)
	endTime := reservationTime.Add(time.Hour)

	var exists bool
	query := `
		SELECT EXISTS (
			SELECT 1
			FROM reservations
			WHERE restaurant_id = $1
			AND (
				$2 BETWEEN reservation_time AND reservation_time + INTERVAL '1 hour'
				OR 
				$3 BETWEEN reservation_time AND reservation_time + INTERVAL '1 hour'
			)
		)
	`
	err = rDb.Db.QueryRow(ctx, query, req.RestaurantId, endTime, reservationTime).Scan(&exists)
	if err != nil {
		log.Error().Err(err).Msg("Error checking reservation availability")
		return nil, err
	}

	return &reservation.CheckAvailabilityResponse{Available: !exists}, nil
}

func (rDB *ReservationDb) FoodList(ctx context.Context, req *reservation.OrderFoodListReq) (*reservation.OrderFoodListRes, error) {
	query := `
		SELECT
			m.id, 
			m.restaurant_id,
			m.name,
			m.description,
			m.price
		FROM 
			menus m
		JOIN
			reservations r
		ON
			r.restaurant_id = m.restaurant_id
		WHERE
			r.id = $1
		AND
			r.deleted_at = 0
		AND
			m.deleted_at = 0
	`
	menusRes := []*reservation.Menus{}
	rows, err := rDB.Db.Query(ctx, query, req.ReservationId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		menu := reservation.Menus{}
		err := rows.Scan(
			&menu.Id,
			&menu.RestaurantId,
			&menu.Name,
			&menu.Description,
			&menu.Price,
		)
		if err != nil {
			log.Error().Err(err).Msg("Error scanning reservation row")
			return nil, err
		}
		menusRes = append(menusRes, &menu)
	}
	return &reservation.OrderFoodListRes{Menus: menusRes}, nil
}

func (rDB *ReservationDb) OrderFood(ctx context.Context, req *reservation.OrderFoodReq) (*reservation.OrderFoodRes, error) {
	if req.Id == "" {
		req.Id = uuid.NewString()
	}
	query := `
		INSERT INTO 
			orders (
				id,
				reservation_id,
				menu_id,
				count
			) 
		VALUES (
				$1, 
				$2, 
				$3,
				$4
			)
	`
	_, err := rDB.Db.Exec(ctx, query, req.Id, req.ReservationId, req.MenuId, req.Count)
	if err != nil {
		log.Error().Err(err).Msg("Error inserting orders information")
		return nil, err
	}

	return &reservation.OrderFoodRes{Message: "Order created succesfully"}, nil
}

func (rDb *ReservationDb) IsValidReservation(ctx context.Context, req *reservation.IsValidReq) (*reservation.IsValidRes, error) {
	var exists bool
	query := `
		SELECT EXISTS (
			SELECT 1
			FROM reservations
			WHERE id = $1 AND deleted_at = 0
		)
	`
	err := rDb.Db.QueryRow(ctx, query, req.Id).Scan(&exists)
	if err != nil {
		log.Error().Err(err).Msg("Error checking reservation validity")
		return nil, err
	}

	return &reservation.IsValidRes{Valid: exists}, nil
}
