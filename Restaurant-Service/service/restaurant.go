package service

import (
	"context"

	"github.com/Project_Restaurant/Restaurant-Service/genproto/restaurant"
	"github.com/Project_Restaurant/Restaurant-Service/storage"
	"github.com/rs/zerolog/log"
)

// RestaurantService implements the restaurant.RestaurantServiceServer interface.
type RestaurantService struct {
	stg storage.StorageI
	restaurant.UnimplementedRestaurantServiceServer
}

// NewRestaurantService creates a new RestaurantService.
func NewRestaurantService(stg storage.StorageI) *RestaurantService {
	return &RestaurantService{stg: stg}
}

// CreateRestaurant creates a new restaurant.
func (s *RestaurantService) CreateRestaurant(ctx context.Context, req *restaurant.CreateRestaurantRequest) (*restaurant.CreateRestaurantResponse, error) {
	log.Info().Msg("RestaurantService: CreateRestaurant called")

	resp, err := s.stg.Restaurant().Create(ctx, req)
	if err != nil {
		log.Error().Err(err).Msg("RestaurantService: Error creating restaurant")
		return nil, err
	}
	return resp, nil
}

// GetRestaurant gets a restaurant by its ID.
func (s *RestaurantService) GetRestaurant(ctx context.Context, req *restaurant.GetRestaurantRequest) (*restaurant.GetRestaurantResponse, error) {
	log.Info().Msg("RestaurantService: GetRestaurant called")

	resp, err := s.stg.Restaurant().GetById(ctx, req)
	if err != nil {
		log.Error().Err(err).Msg("RestaurantService: Error getting restaurant by ID")
		return nil, err
	}
	return resp, nil
}

// UpdateRestaurant updates a restaurant.
func (s *RestaurantService) UpdateRestaurant(ctx context.Context, req *restaurant.UpdateRestaurantRequest) (*restaurant.UpdateRestaurantResponse, error) {
	log.Info().Msg("RestaurantService: UpdateRestaurant called")

	resp, err := s.stg.Restaurant().Update(ctx, req)
	if err != nil {
		log.Error().Err(err).Msg("RestaurantService: Error updating restaurant")
		return nil, err
	}
	return resp, nil
}

// DeleteRestaurant deletes a restaurant.
func (s *RestaurantService) DeleteRestaurant(ctx context.Context, req *restaurant.DeleteRestaurantRequest) (*restaurant.DeleteRestaurantResponse, error) {
	log.Info().Msg("RestaurantService: DeleteRestaurant called")

	resp, err := s.stg.Restaurant().Delete(ctx, req)
	if err != nil {
		log.Error().Err(err).Msg("RestaurantService: Error deleting restaurant")
		return nil, err
	}
	return resp, nil
}

// ListRestaurants lists restaurants with filtering and pagination.
func (s *RestaurantService) ListRestaurants(ctx context.Context, req *restaurant.ListRestaurantsRequest) (*restaurant.ListRestaurantsResponse, error) {
	log.Info().Msg("RestaurantService: ListRestaurants called")

	resp, err := s.stg.Restaurant().ListRestaurants(ctx, req)
	if err != nil {
		log.Error().Err(err).Msg("RestaurantService: Error listing restaurants")
		return nil, err
	}
	return resp, nil
}
