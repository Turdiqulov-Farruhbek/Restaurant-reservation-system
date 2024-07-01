package service

import (
	"context"

	"github.com/Project_Restaurant/Restaurant-Service/genproto/menu"
	"github.com/Project_Restaurant/Restaurant-Service/storage"
	"github.com/rs/zerolog/log"
)

// MenuService implements the menu.MenuServiceServer interface.
type MenuService struct {
	stg storage.StorageI
	menu.UnimplementedMenuServiceServer
}

// NewMenuService creates a new MenuService.
func NewMenuService(stg storage.StorageI) *MenuService {
	return &MenuService{stg: stg}
}

// CreateMenu creates a new menu item.
func (s *MenuService) CreateMenu(ctx context.Context, req *menu.CreateMenuRequest) (*menu.CreateMenuResponse, error) {
	log.Info().Msg("MenuService: CreateMenu called")

	resp, err := s.stg.Menu().Create(ctx, req)
	if err != nil {
		log.Error().Err(err).Msg("MenuService: Error creating menu")
		return nil, err
	}
	return resp, nil
}

// GetMenu gets a menu item by its ID.
func (s *MenuService) GetMenu(ctx context.Context, req *menu.GetMenuRequest) (*menu.GetMenuResponse, error) {
	log.Info().Msg("MenuService: GetMenu called")

	resp, err := s.stg.Menu().GetById(ctx, req)
	if err != nil {
		log.Error().Err(err).Msg("MenuService: Error getting menu by ID")
		return nil, err
	}
	return resp, nil
}

// UpdateMenu updates a menu item.
func (s *MenuService) UpdateMenu(ctx context.Context, req *menu.UpdateMenuRequest) (*menu.UpdateMenuResponse, error) {
	log.Info().Msg("MenuService: UpdateMenu called")

	resp, err := s.stg.Menu().Update(ctx, req)
	if err != nil {
		log.Error().Err(err).Msg("MenuService: Error updating menu")
		return nil, err
	}
	return resp, nil
}

// DeleteMenu deletes a menu item.
func (s *MenuService) DeleteMenu(ctx context.Context, req *menu.DeleteMenuRequest) (*menu.DeleteMenuResponse, error) {
	log.Info().Msg("MenuService: DeleteMenu called")

	resp, err := s.stg.Menu().Delete(ctx, req)
	if err != nil {
		log.Error().Err(err).Msg("MenuService: Error deleting menu")
		return nil, err
	}
	return resp, nil
}

// GetAllMenus gets all menu items.
func (s *MenuService) GetAllMenus(ctx context.Context, req *menu.GetAllMenusRequest) (*menu.GetAllMenusResponse, error) {
	log.Info().Msg("MenuService: GetAllMenus called")

	resp, err := s.stg.Menu().GetAll(ctx, req)
	if err != nil {
		log.Error().Err(err).Msg("MenuService: Error getting all menus")
		return nil, err
	}
	return resp, nil
}
