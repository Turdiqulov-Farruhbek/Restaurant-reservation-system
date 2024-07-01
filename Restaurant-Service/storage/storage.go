package storage

import (
	"context"

	"github.com/Project_Restaurant/Restaurant-Service/genproto/menu"
	"github.com/Project_Restaurant/Restaurant-Service/genproto/reservation"
	"github.com/Project_Restaurant/Restaurant-Service/genproto/restaurant"
)

// StorageI defines the interface for interacting with the restaurant service storage.
type StorageI interface {
	Menu() MenuI
	Reservation() ReservationI
	Restaurant() RestaurantI
}

// MenuI defines methods for managing menu items.
type MenuI interface {
	Create(ctx context.Context, req *menu.CreateMenuRequest) (*menu.CreateMenuResponse, error)
	GetById(ctx context.Context, req *menu.GetMenuRequest) (*menu.GetMenuResponse, error)
	Update(ctx context.Context, req *menu.UpdateMenuRequest) (*menu.UpdateMenuResponse, error)
	Delete(ctx context.Context, req *menu.DeleteMenuRequest) (*menu.DeleteMenuResponse, error)
	GetAll(ctx context.Context, req *menu.GetAllMenusRequest) (*menu.GetAllMenusResponse, error)
}

// ReservationI defines methods for managing reservations.
type ReservationI interface {
	Create(ctx context.Context, req *reservation.CreateReservationRequest) (*reservation.CreateReservationResponse, error)
	GetById(ctx context.Context, req *reservation.GetReservationRequest) (*reservation.GetReservationResponse, error)
	Update(ctx context.Context, req *reservation.UpdateReservationRequest) (*reservation.UpdateReservationResponse, error)
	Delete(ctx context.Context, req *reservation.DeleteReservationRequest) (*reservation.DeleteReservationResponse, error)
	ListReservations(ctx context.Context, req *reservation.ListReservationsRequest) (*reservation.ListReservationsResponse, error)
	CheckAvailability(ctx context.Context, req *reservation.CheckAvailabilityRequest) (*reservation.CheckAvailabilityResponse, error)
	FoodList(ctx context.Context, in *reservation.OrderFoodListReq) (*reservation.OrderFoodListRes, error)
	OrderFood(ctx context.Context, in *reservation.OrderFoodReq) (*reservation.OrderFoodRes, error)
	IsValidReservation(ctx context.Context, in *reservation.IsValidReq) (*reservation.IsValidRes, error)
}

// RestaurantI defines methods for managing restaurant data.
type RestaurantI interface {
	Create(ctx context.Context, req *restaurant.CreateRestaurantRequest) (*restaurant.CreateRestaurantResponse, error)
	GetById(ctx context.Context, req *restaurant.GetRestaurantRequest) (*restaurant.GetRestaurantResponse, error)
	Update(ctx context.Context, req *restaurant.UpdateRestaurantRequest) (*restaurant.UpdateRestaurantResponse, error)
	Delete(ctx context.Context, req *restaurant.DeleteRestaurantRequest) (*restaurant.DeleteRestaurantResponse, error)
	ListRestaurants(ctx context.Context, req *restaurant.ListRestaurantsRequest) (*restaurant.ListRestaurantsResponse, error)
}
