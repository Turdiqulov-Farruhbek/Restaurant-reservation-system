package main

import (
	"log"
	"net"

	"github.com/Project_Restaurant/Restaurant-Service/genproto/menu"
	"github.com/Project_Restaurant/Restaurant-Service/genproto/reservation"
	"github.com/Project_Restaurant/Restaurant-Service/genproto/restaurant"
	"github.com/Project_Restaurant/Restaurant-Service/service"
	"github.com/Project_Restaurant/Restaurant-Service/storage/postgres"
	"google.golang.org/grpc"
)

func main() {

	db, err := postgres.DBConn()
	if err != nil {
		log.Fatal(err)
	}

	liss, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	menu.RegisterMenuServiceServer(s, service.NewMenuService(db))
	reservation.RegisterReservationServiceServer(s, service.NewReservationService(db))
	restaurant.RegisterRestaurantServiceServer(s, service.NewRestaurantService(db))
	log.Printf("server listening at %v", liss.Addr())
	if err := s.Serve(liss); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
