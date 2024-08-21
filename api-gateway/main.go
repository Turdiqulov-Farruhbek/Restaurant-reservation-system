package main

import (
	"api-getway/api" // Import your api package
	"fmt"
	"log"
	// _ "api-getway/docs"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


func main() {
	// Set up gRPC connections
	paymeConn, err := grpc.NewClient("Restourand-payment:50051", grpc.WithTransportCredentials(insecure.NewCredentials())) // Update the address
	if err != nil {
		log.Fatalf("Failed to connect to payment service: %v", err)
	}
	defer paymeConn.Close()

	reserConn, err := grpc.NewClient("restourand-service:8082", grpc.WithTransportCredentials(insecure.NewCredentials())) // Update the address
	if err != nil {
		log.Fatalf("Failed to connect to reservation service: %v", err)
	}
	defer reserConn.Close()

	// menuConn, err := grpc.NewClient("localhost:50053", grpc.WithTransportCredentials(insecure.NewCredentials())) // Update the address
	// if err != nil {
	// 	log.Fatalf("Failed to connect to menu service: %v", err)
	// }
	// defer menuConn.Close()



	// Create the Gin engine with the gRPC clients
	router := api.NewGin(paymeConn, reserConn)
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to connect to gin engine: %v", err)
	}
	fmt.Println("API Gateway running on http://localhost:8080")

}
