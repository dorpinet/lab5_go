package models
import "time"

type Incident struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"size:100;not null" json:"title"`
	Description string    `gorm:"type:text;not null" json:"description"`
	Priority    string    `gorm:"size:20;not null" json:"priority"`
	Status      string    `gorm:"size:20;default:'open'" json:"status"`
	UserID      *uint     `json:"user_id,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}
type User struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Email        string    `gorm:"uniqueIndex;not null" json:"email"`
	PasswordHash string    `gorm:"not null" json:"-"`
	CreatedAt    time.Time `json:"created_at"`
}
type CreateIncidentDTO struct {
	Title       string `json:"title" binding:"required,min=3,max=100"`
	Description string `json:"description" binding:"required"`
	Priority    string `json:"priority" binding:"required,oneof=low medium high critical"`
}
type UpdateStatusDTO struct {
	Status string `json:"status" binding:"required,oneof=open in_progress resolved closed"`
}
type RegisterDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}
type LoginDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}