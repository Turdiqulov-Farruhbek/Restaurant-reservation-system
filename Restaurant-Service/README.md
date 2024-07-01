## Restaurant Service

This repository contains the source code for a gRPC-based Restaurant Service, providing functionalities for managing restaurants, menus, and reservations.

### Features

- **Restaurant Management:**

  - Create, retrieve, update, and delete restaurant information (name, address, phone number, description).
  - List restaurants with optional filtering and pagination.

- **Menu Management:**

  - Create, retrieve, update, and delete menu items (name, description, price) associated with restaurants.
  - Retrieve all menu items, optionally filtered by restaurant ID, name, description, and price range.

- **Reservation Management:**
  - Create, retrieve, update, and delete reservations for restaurants.
  - List reservations with filtering options (by user, restaurant, status, time range) and pagination.
  - Check the availability of a reservation time slot for a specific restaurant.

### Technologies Used

- **Go:** The primary programming language used for the service implementation.
- **gRPC:** A high-performance, open-source universal RPC framework used for communication between services.
- **Protocol Buffers:** Language-neutral, platform-neutral mechanism for serializing structured data used for defining service interfaces and message formats.
- **PostgreSQL:** The database used for persisting restaurant, menu, and reservation data.

### Project Structure

- `genproto/`: Contains the Protocol Buffer definitions for the service interfaces and messages.
- `storage/`: Defines the storage interface and concrete implementations (e.g., `postgres` for PostgreSQL).
- `service/`: Contains the gRPC service implementations for restaurant, menu, and reservation operations.
- `test/`: Contains unit tests for the service and storage implementations.

### Getting Started

1. **Prerequisites:**

   - Go (1.18 or later) installed on your system.
   - PostgreSQL database server running.
   - Protocol Buffer compiler (`protoc`) installed.

2. **Generate gRPC Code:**
   ```bash
   protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       ./genproto/*.proto
   ```
