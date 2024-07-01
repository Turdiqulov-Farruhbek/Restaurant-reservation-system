package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/Project_Restaurant/Payment-Service/genproto/payment"
	"github.com/Project_Restaurant/Payment-Service/service"
	"github.com/Project_Restaurant/Payment-Service/storage/postgres"

	"google.golang.org/grpc"
)

func main() {
	db, err := postgres.DbConnection()
	if err != nil {
		log.Fatal("Error while db connection", err)
		return
	}

	newServer := grpc.NewServer()

	pb.RegisterPaymentServiceServer(newServer, service.NewPAymentService(*db))

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("Error while listen tcp", err)
		return
	}
	fmt.Println("Server is liste PaymentServer", lis.Addr())
	err = newServer.Serve(lis)
	if err != nil {
		log.Fatal("Error while newServe", err)
	}
}
