package services

import (
	"goauth/models"
	"goauth/repository"
	"testing"
)

func TestLogin(t *testing.T) {
	repository.InitDB("test.db")
	defer repository.CloseDB()

	// Create a test user
	testUser := &models.User{
		Username: "testuser",
		Password: "testpassword",
	}
	err := registerUser(testUser.Username, testUser.Password)
	if err != nil {
		t.Fatalf("Failed to register user: %v", err)
	}

	// Test login with correct credentials
	status, _ := Login(testUser)
	if status != 200 {
		t.Errorf("Expected status 200, got %d", status)
	}

	// Test login with incorrect credentials
	testUser.Password = "wrongpassword"
	status, _ = Login(testUser)
	if status != 401 {
		t.Errorf("Expected status 401, got %d", status)
	}
}

func TestRegister(t *testing.T) {
	repository.InitDB("test.db")
	defer repository.CloseDB()

	// Create a test user
	testUser := &models.User{
		Username: "newuser",
		Password: "newpassword",
	}

	// Test registration
	status, _ := Register(testUser)
	if status != 200 {
		t.Errorf("Expected status 200, got %d", status)
	}

	// Test registration with existing username
	status, _ = Register(testUser)
	if status != 500 {
		t.Errorf("Expected status 500, got %d", status)
	}
}

func TestRefreshToken(t *testing.T) {
	repository.InitDB("test.db")
	defer repository.CloseDB()

	// Create a test user
	testUser := &models.User{
		Username: "testuser",
		Password: "testpassword",
	}
	err := registerUser(testUser.Username, testUser.Password)
	if err != nil {
		t.Fatalf("Failed to register user: %v", err)
	}

	// Test refresh token
	status, _ := RefreshToken(testUser)
	if status != 200 {
		t.Errorf("Expected status 200, got %d", status)
	}
}

func TestLogout(t *testing.T) {
	// Test logout (currently a stub)
	err := Logout(nil)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
