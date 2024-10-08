// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/menu": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get all menus",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "menu"
                ],
                "summary": "Get all menus",
                "responses": {
                    "200": {
                        "description": "Invalid request body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Create new menu",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "menu"
                ],
                "summary": "Create new menu",
                "parameters": [
                    {
                        "description": "Menu",
                        "name": "menu",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/reservation.CreateMenuRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Invalid request body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/menu/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get menu by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "menu"
                ],
                "summary": "Get menu by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Menu ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/reservation.GetMenuResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update menu",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "menu"
                ],
                "summary": "Update menu",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Menu ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Menu",
                        "name": "menu",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/reservation.UpdateMenuRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Invalid request body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Delete menu",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "menu"
                ],
                "summary": "Delete menu",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Menu ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Invalid request body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/payments": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Create new payment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "payment"
                ],
                "summary": "Create new payment",
                "parameters": [
                    {
                        "description": "Payment",
                        "name": "payment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/payment.CreatePaymentRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Invalid request body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/payments/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get payment by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "payment"
                ],
                "summary": "Get payment by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Payment ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Invalid request body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update payment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "payment"
                ],
                "summary": "Update payment",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Payment ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Payment",
                        "name": "payment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/payment.UpdatePaymentRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Invalid request body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/reservations": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get all reservations",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "reservation"
                ],
                "summary": "Get all reservations",
                "responses": {
                    "200": {
                        "description": "Invalid request body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Create new reservation",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "reservation"
                ],
                "summary": "Create new reservation",
                "parameters": [
                    {
                        "description": "Reservation",
                        "name": "reservation",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/reservation.CreateReservationRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Invalid request body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/reservations/check": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Check reservation",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "reservation"
                ],
                "summary": "Check reservation",
                "parameters": [
                    {
                        "description": "Check",
                        "name": "check",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Invalid request body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/reservations/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get reservation by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "reservation"
                ],
                "summary": "Get reservation by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Reservation ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Invalid request body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update reservation",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "reservation"
                ],
                "summary": "Update reservation",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Reservation ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Reservation",
                        "name": "reservation",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/reservation.UpdateReservationRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Invalid request body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Delete reservation",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "reservation"
                ],
                "summary": "Delete reservation",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Reservation ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Invalid request body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/reservations/{id}/order": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Order reservation by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "reservation"
                ],
                "summary": "Order reservation by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Reservation ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Order",
                        "name": "order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/reservation.OrderFoodReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Invalid request body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/reservations/{id}/payment": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Process payment for reservation",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "reservation"
                ],
                "summary": "Process payment for reservation",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Reservation ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Payment",
                        "name": "payment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/reservation.PaymentReservationRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Invalid request body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/restaurants": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get all restaurants",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "restaurant"
                ],
                "summary": "Get all restaurants",
                "responses": {
                    "200": {
                        "description": "Invalid request body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Create new restaurant",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "restaurant"
                ],
                "summary": "Create new restaurant",
                "parameters": [
                    {
                        "description": "Restaurant",
                        "name": "restaurant",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/reservation.CreateRestaurantRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Invalid request body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/restaurants/{id}": {
            "get": {
                "description": "Get restaurant by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "restaurant"
                ],
                "summary": "Get restaurant by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Restaurant ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Invalid request body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update restaurant",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "restaurant"
                ],
                "summary": "Update restaurant",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Restaurant ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Restaurant",
                        "name": "restaurant",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/reservation.UpdateRestaurantRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Invalid request body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Delete restaurant",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "restaurant"
                ],
                "summary": "Delete restaurant",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Restaurant ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Invalid request body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "payment.CreatePaymentRequest": {
            "type": "object",
            "properties": {
                "payment": {
                    "$ref": "#/definitions/payment.Payment"
                }
            }
        },
        "payment.Payment": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "payment_method": {
                    "type": "string"
                },
                "payment_status": {
                    "type": "string"
                },
                "reservation_id": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "payment.UpdatePaymentRequest": {
            "type": "object",
            "properties": {
                "payment": {
                    "$ref": "#/definitions/payment.Payment"
                }
            }
        },
        "reservation.CreateMenuRequest": {
            "type": "object",
            "properties": {
                "menu": {
                    "$ref": "#/definitions/reservation.Menu"
                }
            }
        },
        "reservation.CreateReservationRequest": {
            "type": "object",
            "properties": {
                "reservation": {
                    "$ref": "#/definitions/reservation.Reservation"
                }
            }
        },
        "reservation.CreateRestaurantRequest": {
            "type": "object",
            "properties": {
                "restaurant": {
                    "$ref": "#/definitions/reservation.Restaurant"
                }
            }
        },
        "reservation.GetMenuResponse": {
            "type": "object",
            "properties": {
                "menu": {
                    "$ref": "#/definitions/reservation.Menu"
                }
            }
        },
        "reservation.Menu": {
            "type": "object",
            "properties": {
                "created_at": {
                    "description": "Timestamp with time zone",
                    "type": "string"
                },
                "deleted_at": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "restaurant_id": {
                    "type": "string"
                },
                "updated_at": {
                    "description": "Timestamp with time zone",
                    "type": "string"
                }
            }
        },
        "reservation.OrderFoodReq": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "menu_id": {
                    "type": "string"
                },
                "reservation_id": {
                    "type": "string"
                }
            }
        },
        "reservation.PaymentReservationRequest": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "payment_method": {
                    "type": "string"
                },
                "payment_status": {
                    "type": "string"
                },
                "reservation_id": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "reservation.Reservation": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "reservation_time": {
                    "type": "string"
                },
                "restaurant_id": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "reservation.Restaurant": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "reservation.UpdateMenuRequest": {
            "type": "object",
            "properties": {
                "menu": {
                    "$ref": "#/definitions/reservation.Menu"
                }
            }
        },
        "reservation.UpdateReservationRequest": {
            "type": "object",
            "properties": {
                "reservation": {
                    "$ref": "#/definitions/reservation.Reservation"
                }
            }
        },
        "reservation.UpdateRestaurantRequest": {
            "type": "object",
            "properties": {
                "restaurant": {
                    "$ref": "#/definitions/reservation.Restaurant"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Restaurant reservation system API",
	Description:      "API for managing Restaurant reservation syste resources",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	// LeftDelim:        "{{",
	// RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
