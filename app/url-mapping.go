package app

import (
	"github.com/d5avard/bookstore-users-api/controllers/ping"
	"github.com/d5avard/bookstore-users-api/controllers/users"
)

func urlMap() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:id", users.GetUser)
	router.POST("/users", users.CreateUser)
}
