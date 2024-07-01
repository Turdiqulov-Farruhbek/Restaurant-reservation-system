package postgres

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Project_Restaurant/Payment-Service/config"
	"github.com/Project_Restaurant/Payment-Service/storage"
	_ "github.com/lib/pq"
)

type Storage struct {
	Db       *sql.DB
	Payments storage.PaymentI
}

func DbConnection() (*Storage, error) {
	cfg := config.Load()
	con := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%d sslmode=disable",
		cfg.PostgresHost, cfg.PostgresUser, cfg.PostgresDatabase, cfg.PostgresPassword, cfg.PostgresPort)
	db, err := sql.Open("postgres", con)
	if err != nil {
		log.Fatal("Error while db connection" , err)
		return nil, nil
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Error while db ping connection ", err)
		return nil, nil
	}
	payment := NewPaymentRepo(db)
	return &Storage{
		Db:       db,
		Payments: payment,
	}, nil
}

func (stg *Storage) Payment() *storage.PaymentI {
	if stg.Payments == nil {
		stg.Payments = NewPaymentRepo(stg.Db)
	}
	return &stg.Payments
}
