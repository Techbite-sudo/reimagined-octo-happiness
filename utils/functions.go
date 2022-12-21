package utils

import (
	"context"
	"errors"
	"lms_backend/graph/model"
	"crypto/rand"
	"encoding/hex"
)

// findUserByEmail finds a user with the given email in the database.
func findUserByEmail(email string) (*model.User, error) {
	// Replace this placeholder implementation with your own database query
	if email == "john@example.com" {
		return &model.User{
			ID:       "123",
			Email:    "john@example.com",
			Password: "password",
			Name:     "John Smith",
		}, nil
	}
	return nil, nil
}

// addUser adds a user to the database.
func addUser(user *model.User) error {
	// Replace this placeholder implementation with your own database query
	if user.Email == "john@example.com" {
		return errors.New("user with email john@example.com already exists")
	}
	return nil
}

// generateID generates a unique identifier.
func generateID() string {
	// Generate a random 16-byte slice
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		// This should never happen, but if it does, return an empty ID
		return ""
	}

	// Encode the random bytes as a hexadecimal string
	return hex.EncodeToString(b)
}