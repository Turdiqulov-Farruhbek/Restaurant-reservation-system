package postgres

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/Project_Restaurant/Restaurant-Service/config"
	"github.com/Project_Restaurant/Restaurant-Service/storage"
	"github.com/jackc/pgx/v5"
)

type Storage struct {
	DB           pgx.Conn
	MenuS        storage.MenuI
	ReservationS storage.ReservationI
	RestaurantS  storage.RestaurantI
}

func DBConn() (*Storage, error) {
	var (
		db  *pgx.Conn
		err error
	)
	// Get postgres connection data from .env file
	cfg := config.Load()
	dbCon := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase)

	// Connecting to postgres
	db, err = pgx.Connect(context.Background(), dbCon)
	if err != nil {
		slog.Warn("Unable to connect to database:", err)
	}
	err = db.Ping(context.Background())
	if err != nil {
		return nil, err
	}
	return &Storage{}, err
}

func (s *Storage) Menu() storage.MenuI {
	if s.MenuS == nil {
		s.MenuS = NewMenu(&s.DB)
	}
	return s.MenuS
}
func (s *Storage) Reservation() storage.ReservationI {
	if s.ReservationS == nil {
		s.ReservationS = NewReservation(&s.DB)
	}
	return s.ReservationS
}

func (s *Storage) Restaurant() storage.RestaurantI {
	if s.RestaurantS == nil {
		s.RestaurantS = NewRestaurant(&s.DB)
	}
	return s.RestaurantS
}
