package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/Project_Restaurant/Restaurant-Service/genproto/menu"
	"github.com/Project_Restaurant/Restaurant-Service/storage/postgres"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
)

func newTestMenu(t *testing.T) *postgres.MenuDb {
	connString := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
		"sayyidmuhammad",
		"root",
		"localhost",
		5432,
		"reservation",
	)

	db, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	return &postgres.MenuDb{Db: db}
}

func createTestMenu() *menu.CreateMenuRequest {
	return &menu.CreateMenuRequest{
		Menu: &menu.Menu{
			RestaurantId: "345e1148-678e-4ec7-9fab-d55df1e9cb54",
			Name:         "Test Menu",
			Description:  "This is a test menu item",
			Price:        9.99,
		},
	}
}

func TestCreateMenu(t *testing.T) {
	mDb := newTestMenu(t)
	testMenu := createTestMenu()

	createdMenu, err := mDb.Create(context.Background(), testMenu)
	if err != nil {
		t.Fatalf("Error creating menu: %v", err)
	}

	assert.NotEmpty(t, createdMenu.Menu.Id)
	assert.Equal(t, testMenu.Menu.RestaurantId, createdMenu.Menu.RestaurantId)
	assert.Equal(t, testMenu.Menu.Name, createdMenu.Menu.Name)
	assert.Equal(t, testMenu.Menu.Description, createdMenu.Menu.Description)
	assert.Equal(t, testMenu.Menu.Price, createdMenu.Menu.Price)
}

func TestGetMenuById(t *testing.T) {
	mDb := newTestMenu(t)
	testMenu := createTestMenu()

	createdMenu, err := mDb.Create(context.Background(), testMenu)
	if err != nil {
		t.Fatalf("Error creating menu: %v", err)
	}

	getMenu, err := mDb.GetById(context.Background(), &menu.GetMenuRequest{Id: createdMenu.Menu.Id})
	if err != nil {
		t.Fatalf("Error getting menu by ID: %v", err)
	}

	assert.Equal(t, createdMenu.Menu.Id, getMenu.Menu.Id)
	assert.Equal(t, createdMenu.Menu.RestaurantId, getMenu.Menu.RestaurantId)
	assert.Equal(t, createdMenu.Menu.Name, getMenu.Menu.Name)
	assert.Equal(t, createdMenu.Menu.Description, getMenu.Menu.Description)
	assert.Equal(t, createdMenu.Menu.Price, getMenu.Menu.Price)
}

func TestUpdateMenu(t *testing.T) {
	mDb := newTestMenu(t)
	testMenu := createTestMenu()

	createdMenu, err := mDb.Create(context.Background(), testMenu)
	if err != nil {
		t.Fatalf("Error creating menu: %v", err)
	}

	updatedMenuReq := &menu.UpdateMenuRequest{
		Menu: &menu.Menu{
			Id:           createdMenu.Menu.Id,
			RestaurantId: "345e1148-678e-4ec7-9fab-d55df1e9cb54",
			Name:         "Updated Menu Name",
			Description:  "Updated menu description",
			Price:        12.99,
		},
	}

	updatedMenu, err := mDb.Update(context.Background(), updatedMenuReq)
	if err != nil {
		t.Fatalf("Error updating menu: %v", err)
	}

	assert.Equal(t, updatedMenuReq.Menu.Id, updatedMenu.Menu.Id)
	assert.Equal(t, updatedMenuReq.Menu.RestaurantId, updatedMenu.Menu.RestaurantId)
	assert.Equal(t, updatedMenuReq.Menu.Name, updatedMenu.Menu.Name)
	assert.Equal(t, updatedMenuReq.Menu.Description, updatedMenu.Menu.Description)
	assert.Equal(t, updatedMenuReq.Menu.Price, updatedMenu.Menu.Price)
}

func TestDeleteMenu(t *testing.T) {
	mDb := newTestMenu(t)
	testMenu := createTestMenu()

	createdMenu, err := mDb.Create(context.Background(), testMenu)
	if err != nil {
		t.Fatalf("Error creating menu: %v", err)
	}

	_, err = mDb.Delete(context.Background(), &menu.DeleteMenuRequest{Id: createdMenu.Menu.Id})
	if err != nil {
		t.Fatalf("Error deleting menu: %v", err)
	}
	// Try to get the deleted menu
	deletedMenu, err := mDb.GetById(context.Background(), &menu.GetMenuRequest{Id: createdMenu.Menu.Id})
	assert.ErrorIs(t, err, postgres.ErrMenuNotFound)
	assert.Nil(t, deletedMenu)

}

func TestGetAllMenus(t *testing.T) {
	mDb := newTestMenu(t)

	// Create some test menus
	restaurantID1 := "345e1148-678e-4ec7-9fab-d55df1e9cb54"
	restaurantID2 := "60e83378-4e69-4e96-b24d-87c753363725"

	testMenus := []*menu.CreateMenuRequest{
		{Menu: &menu.Menu{RestaurantId: restaurantID1, Name: "Menu 1", Description: "Desc 1", Price: 10.99}},
		{Menu: &menu.Menu{RestaurantId: restaurantID1, Name: "Menu 2", Description: "Desc 2", Price: 5.99}},
		{Menu: &menu.Menu{RestaurantId: restaurantID2, Name: "Menu 3", Description: "Delicious Dish", Price: 15.99}},
	}

	for _, tm := range testMenus {
		_, err := mDb.Create(context.Background(), tm)
		if err != nil {
			t.Fatalf("Error creating test menu: %v", err)
		}
	}

	t.Run("GetAllMenus without filters", func(t *testing.T) {
		resp, err := mDb.GetAll(context.Background(), &menu.GetAllMenusRequest{})
		if err != nil {
			t.Fatalf("Error getting all menus: %v", err)
		}
		assert.LessOrEqual(t, len(testMenus), len(resp.Menus))
	})

	t.Run("Filter by restaurant ID", func(t *testing.T) {
		resp, err := mDb.GetAll(context.Background(), &menu.GetAllMenusRequest{RestaurantId: restaurantID1})
		if err != nil {
			t.Fatalf("Error getting menus by restaurant ID: %v", err)
		}
		assert.LessOrEqual(t, 2, len(resp.Menus)) // Should be 2 menus for restaurantID1
	})

	t.Run("Filter by name (case-insensitive)", func(t *testing.T) {
		resp, err := mDb.GetAll(context.Background(), &menu.GetAllMenusRequest{Name: "menu"})
		if err != nil {
			t.Fatalf("Error getting menus by name: %v", err)
		}
		assert.LessOrEqual(t, 3, len(resp.Menus)) // Should match all 3 menus
	})

	t.Run("Filter by description (partial match)", func(t *testing.T) {
		resp, err := mDb.GetAll(context.Background(), &menu.GetAllMenusRequest{Description: "cious"})
		if err != nil {
			t.Fatalf("Error getting menus by description: %v", err)
		}
		assert.LessOrEqual(t, 1, len(resp.Menus)) // Should match "Menu 3"
	})

	t.Run("Filter by price range", func(t *testing.T) {
		resp, err := mDb.GetAll(context.Background(), &menu.GetAllMenusRequest{MinPrice: 6.00, MaxPrice: 12.00})
		if err != nil {
			t.Fatalf("Error getting menus by price range: %v", err)
		}
		assert.LessOrEqual(t, 1, len(resp.Menus)) // Should match "Menu 1" only
	})

	// Add more test cases as needed to cover different filter combinations
}
