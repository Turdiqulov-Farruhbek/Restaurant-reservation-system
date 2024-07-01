package api
// "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTgyODM3OTYsImlhdCI6MTcxODE5NzM5NiwiaWQiOiJqYXZveGlyIiwidXNlcm5hbWUiOiI4NzEwMWRjNy0wZTc2LTQ4YWMtYTJmYi0xNjllOWE5MGQ1N2IifQ.InfjefXlbVoAd-7S52ZJiNf-iMuPZcyxq75AinHF_o4"
import (
	"api-getway/api/handlers"
	"api-getway/api/middlerware"
	_ "api-getway/api/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/grpc"
)

// @title Restaurant reservation system API
// @version 1.0
// @description API for managing Restaurant reservation syste resources
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewGin(payme *grpc.ClientConn) *gin.Engine {
	handler := handlers.NewHandler(payme)

	router := gin.Default()
	router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Use(middlerware.Authorizations)
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Adjust for your specific origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	restaurant := router.Group("api/v1")
	{
		restaurant.POST("/restaurants", handler.CreateRestaurantHandler)
		restaurant.GET("/restaurants", handler.GetRestaurantAllHandler)
		restaurant.GET("/restaurants/:id", handler.GetRestaurantByIdHandler)
		restaurant.PUT("/restaurants/:id", handler.UpdateRestaurantHandler)
		restaurant.DELETE("/restaurants/:id", handler.DeleteRestaurantHandler)
	}

	reservation := router.Group("api/v1")
	{
		reservation.POST("/reservations", handler.CreateReservationHandler)
		reservation.GET("/reservations", handler.GetReservationAllHandler)
		reservation.GET("/reservations/:id", handler.GetReservationByIdHandler)
		reservation.PUT("/reservations/:id", handler.UpdateReservationHandler)
		reservation.DELETE("/reservations/:id", handler.DeleteReservationHandler)
		reservation.POST("/reservations/check", handler.ReservationCheckHandler)
		reservation.POST("/reservations/:id/order", handler.ReservationOrderIdHandler)
		reservation.POST("/reservations/:id/payment", handler.ReservationPaymentHandler)
		reservation.POST("/reservations/OrderFood:id", handler.ReservationOrderFood)
	}

	menu := router.Group("api/v1")
	{
		menu.POST("/menu", handler.CreateMenuHandler)
		menu.GET("/menu", handler.GetAllMenuHandler)
		menu.GET("/menu/:id", handler.GetMenuByIdHandler)
		menu.PUT("/menu/:id", handler.UpdateMenuHandler)
		menu.DELETE("/menu/:id", handler.DeleteMenuHandler)
	}

	payment := router.Group("api/v1")
	{
		payment.POST("/payments", handler.CreatePaymentHandler)
		payment.GET("/payments/:id", handler.GetByIdPaymentHandler)
		payment.PUT("/payments/:id", handler.UpdatePaymentHandler)
	}

	return router
}
