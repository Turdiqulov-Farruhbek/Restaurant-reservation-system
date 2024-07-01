package postgres

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Project_Restaurant/Restaurant-Service/genproto/menu"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
)

// ErrMenuNotFound represents an error when a menu item is not found.
var ErrMenuNotFound = errors.New("menu item not found")

// MenuDb provides database operations for menu items
type MenuDb struct {
	Db *pgx.Conn
}

// NewMenu creates a new instance of MenuDb
func NewMenu(db *pgx.Conn) *MenuDb {
	return &MenuDb{Db: db}
}

// Create creates a new menu item in the database
func (mDb *MenuDb) Create(ctx context.Context, req *menu.CreateMenuRequest) (*menu.CreateMenuResponse, error) {
	// Generate a new UUID for the menu item
	menuID := uuid.New().String()

	query := `
		INSERT INTO 
			menus (
				id,
				restaurant_id, 
				name,
				description,
				price
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
			restaurant_id,
			name,
			description,
			price,
			created_at,
			updated_at,
			deleted_at
	`
	var (
		dbMenu    menu.Menu
		createdAt time.Time
		updatedAt time.Time
	)

	err := mDb.Db.QueryRow(ctx, query,
		menuID,
		req.Menu.RestaurantId,
		req.Menu.Name,
		req.Menu.Description,
		req.Menu.Price,
	).Scan(
		&dbMenu.Id,
		&dbMenu.RestaurantId,
		&dbMenu.Name,
		&dbMenu.Description,
		&dbMenu.Price,
		&createdAt,
		&updatedAt,
		&dbMenu.DeletedAt,
	)
	if err != nil {
		log.Error().Err(err).Msg("Error creating menu item")
		return nil, err
	}

	// Convert time.Time to string
	dbMenu.CreatedAt = createdAt.Format(time.RFC3339)
	dbMenu.UpdatedAt = updatedAt.Format(time.RFC3339)

	return &menu.CreateMenuResponse{Menu: &dbMenu}, nil
}

// GetById gets a menu item by its ID
func (mDb *MenuDb) GetById(ctx context.Context, req *menu.GetMenuRequest) (*menu.GetMenuResponse, error) {
	var (
		dbMenu    menu.Menu
		createdAt time.Time
		updatedAt time.Time
	)

	query := `
		SELECT
			id,
			restaurant_id, 
			name,
			description,
			price,
			created_at,
			updated_at,
			deleted_at
		FROM 
			menus
		WHERE 
			id = $1 AND deleted_at = 0
	`
	err := mDb.Db.QueryRow(ctx, query, req.Id).Scan(
		&dbMenu.Id,
		&dbMenu.RestaurantId,
		&dbMenu.Name,
		&dbMenu.Description,
		&dbMenu.Price,
		&createdAt,
		&updatedAt,
		&dbMenu.DeletedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			log.Error().Err(err).Msg("Menu item not found")
			return nil, ErrMenuNotFound
		}
		log.Error().Err(err).Msg("Error getting menu item by ID")
		return nil, err
	}

	// Convert time.Time to string
	dbMenu.CreatedAt = createdAt.Format(time.RFC3339)
	dbMenu.UpdatedAt = updatedAt.Format(time.RFC3339)

	return &menu.GetMenuResponse{Menu: &dbMenu}, nil
}

// Update updates an existing menu item in the database
func (mDb *MenuDb) Update(ctx context.Context, req *menu.UpdateMenuRequest) (*menu.UpdateMenuResponse, error) {
	var args []interface{}
	count := 1
	query := `
		UPDATE 
			menus 
		SET `
	filter := ``

	if len(req.Menu.RestaurantId) > 0 {
		filter += fmt.Sprintf(" restaurant_id = $%d, ", count)
		args = append(args, req.Menu.RestaurantId)
		count++
	}

	if len(req.Menu.Name) > 0 {
		filter += fmt.Sprintf(" name = $%d, ", count)
		args = append(args, req.Menu.Name)
		count++
	}

	if len(req.Menu.Description) > 0 {
		filter += fmt.Sprintf(" description = $%d, ", count)
		args = append(args, req.Menu.Description)
		count++
	}

	if req.Menu.Price != 0 {
		filter += fmt.Sprintf(" price = $%d, ", count)
		args = append(args, req.Menu.Price)
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
		AND
			deleted_at = 0
			`, count)

	args = append(args, req.Menu.Id)
	query += filter

	_, err := mDb.Db.Exec(ctx, query, args...)
	if err != nil {
		if err == pgx.ErrNoRows {
			log.Error().Err(err).Msg("Menu item not found")
			return nil, ErrMenuNotFound
		}
		log.Error().Err(err).Msg("Error updating menu item")
		return nil, err
	}

	// After successful update, get the updated menu item from the database
	updatedMenu, err := mDb.GetById(ctx, &menu.GetMenuRequest{Id: req.Menu.Id})
	if err != nil {
		return nil, err
	}

	return &menu.UpdateMenuResponse{Menu: updatedMenu.Menu}, nil
}

// Delete deletes a menu item by using deleted_at table (Soft delete).
func (mDb *MenuDb) Delete(ctx context.Context, req *menu.DeleteMenuRequest) (*menu.DeleteMenuResponse, error) {
	query := `
		UPDATE 
			menus
		SET 
			deleted_at = $1 
		WHERE 
			id = $2
	`
	_, err := mDb.Db.Exec(ctx, query, time.Now().Unix(), req.Id)
	if err != nil {
		if err == pgx.ErrNoRows {
			log.Error().Err(err).Msg("Menu item not found")
			return nil, ErrMenuNotFound
		}
		log.Error().Err(err).Msg("Error deleting menu item")
		return nil, err
	}
	return &menu.DeleteMenuResponse{Message: "Menu item deleted successfully"}, nil
}

func (mDb *MenuDb) GetAll(ctx context.Context, req *menu.GetAllMenusRequest) (*menu.GetAllMenusResponse, error) {
	var args []interface{}
	count := 1
	query := `
		SELECT
			id,
			restaurant_id, 
			name,
			description,
			price,
			created_at,
			updated_at,
			deleted_at
		FROM 
			menus
		WHERE 
			deleted_at = 0
	`
	filter := ""

	if req.RestaurantId != "" {
		filter += fmt.Sprintf(" AND restaurant_id = $%d", count)
		args = append(args, req.RestaurantId)
		count++
	}

	if req.Name != "" {
		filter += fmt.Sprintf(" AND name ILIKE $%d", count) // ILIKE for case-insensitive search
		args = append(args, "%"+req.Name+"%")               // Add wildcards for partial match
		count++
	}

	if req.Description != "" {
		filter += fmt.Sprintf(" AND description ILIKE $%d", count)
		args = append(args, "%"+req.Description+"%")
		count++
	}

	if req.MinPrice != 0 {
		filter += fmt.Sprintf(" AND price >= $%d", count)
		args = append(args, req.MinPrice)
		count++
	}

	if req.MaxPrice != 0 {
		filter += fmt.Sprintf(" AND price <= $%d", count)
		args = append(args, req.MaxPrice)
		count++
	}

	query += filter

	rows, err := mDb.Db.Query(ctx, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("Error fetching menus from the database")
		return nil, err
	}
	defer rows.Close()

	var menus []*menu.Menu
	for rows.Next() {
		var (
			createdAt time.Time
			updatedAt time.Time
			m         menu.Menu
		)

		err := rows.Scan(
			&m.Id,
			&m.RestaurantId,
			&m.Name,
			&m.Description,
			&m.Price,
			&createdAt,
			&updatedAt,
			&m.DeletedAt,
		)
		if err != nil {
			log.Error().Err(err).Msg("Error scanning menu row")
			return nil, err
		}
		m.CreatedAt = createdAt.Format(time.RFC3339)
		m.UpdatedAt = updatedAt.Format(time.RFC3339)

		menus = append(menus, &m)
	}

	if err = rows.Err(); err != nil {
		log.Error().Err(err).Msg("Error iterating over menu rows")
		return nil, err
	}

	return &menu.GetAllMenusResponse{Menus: menus}, nil
}
