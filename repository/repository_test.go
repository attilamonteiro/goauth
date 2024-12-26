package repository

import (
    "goauth/models"
    "testing"
)

func TestInitDB(t *testing.T) {
    InitDB("test.db")
    defer CloseDB()

    if db == nil {
        t.Fatal("Expected db to be initialized, got nil")
    }
}

func TestFindByUsername(t *testing.T) {
    InitDB("test.db")
    defer CloseDB()

    // Create a test user
    testUser := &models.User{
        UUID:     "test-uuid",
        Username: "testuser",
        Password: "testpassword",
    }
    err := CreateUser(testUser)
    if err != nil {
        t.Fatalf("Failed to create user: %v", err)
    }

    // Test finding the user by username
    user, err := FindByUsername(testUser.Username)
    if err != nil {
        t.Fatalf("Failed to find user: %v", err)
    }
    if user.Username != testUser.Username {
        t.Errorf("Expected username %s, got %s", testUser.Username, user.Username)
    }
}

func TestCreateUser(t *testing.T) {
    InitDB("test.db")
    defer CloseDB()

    // Create a test user
    testUser := &models.User{
        UUID:     "test-uuid",
        Username: "newuser",
        Password: "newpassword",
    }

    // Test creating the user
    err := CreateUser(testUser)
    if err != nil {
        t.Fatalf("Failed to create user: %v", err)
    }

    // Test creating the user with the same username
    err = CreateUser(testUser)
    if err == nil {
        t.Fatal("Expected error, got nil")
    }
}
