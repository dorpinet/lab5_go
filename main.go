package main

import (
	"log"

	"incident-tracker/config"
	"incident-tracker/handlers"
	"incident-tracker/models"

	"github.com/gin-gonic/gin"
)
func main() {
	db := config.InitDB()
	db.AutoMigrate(&models.User{}, &models.Incident{})
	log.Println("[GORM] Auto-migration completed: users, incidents")

	secret := "super-secret-key-change-me-in-prod"
	h := handlers.NewHandler(db, secret)
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/register", h.Register)
			auth.POST("/login", h.Login)
		}

		incidents := v1.Group("/incidents")
		{
			incidents.GET("", h.ListIncidents)
			incidents.POST("", h.CreateIncident)
			incidents.GET("/:id", h.GetIncident)
			incidents.PATCH("/:id/status", h.UpdateStatus)
			incidents.DELETE("/:id", h.DeleteIncident)
		}
	}
	log.Println("[SERVER] running on http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}