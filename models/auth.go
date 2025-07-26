package models

// LoginRequest represents the payload for login requests
// Name and Password are required

type LoginRequest struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}
