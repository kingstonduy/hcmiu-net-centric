package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

const (
	HTTP_PORT = 8080
	HTTP_HOST = "localhost"

	DB_HOST     = "localhost"
	DB_PORT     = 5432
	DB_USER     = "postgres"
	DB_PASSWORD = "changeme"
	DB_NAME     = "postgres"
)

var (
	db *gorm.DB
)

// Define User struct for GORM
type User struct {
	UserID            int       `json:"user_id" gorm:"primary_key"`
	FirstName         string    `json:"first_name"`
	LastName          string    `json:"last_name"`
	Username          string    `json:"username"`
	Email             string    `json:"email"`
	Avatar            string    `json:"avatar,omitempty"`
	PhoneNumber       string    `json:"phone_number,omitempty"`
	DateOfBirth       time.Time `json:"date_of_birth"`
	AddressCountry    string    `json:"address_country"`
	AddressCity       string    `json:"address_city"`
	AddressStreetName string    `json:"address_street_name"`
	AddressStreetAddr string    `json:"address_street_address"`
}

type UserDto struct {
	ID                    int          `json:"id"`
	UID                   string       `json:"uid"`
	Password              string       `json:"password"`
	FirstName             string       `json:"first_name"`
	LastName              string       `json:"last_name"`
	Username              string       `json:"username"`
	Email                 string       `json:"email"`
	Avatar                string       `json:"avatar"`
	Gender                string       `json:"gender"`
	PhoneNumber           string       `json:"phone_number"`
	SocialInsuranceNumber string       `json:"social_insurance_number"`
	DateOfBirth           string       `json:"date_of_birth"` // You might want to use time.Time if you're parsing dates
	Employment            Employment   `json:"employment"`
	Address               Address      `json:"address"`
	CreditCard            CreditCard   `json:"credit_card"`
	Subscription          Subscription `json:"subscription"`
}

type Employment struct {
	Title    string `json:"title"`
	KeySkill string `json:"key_skill"`
}

type Address struct {
	City          string      `json:"city"`
	StreetName    string      `json:"street_name"`
	StreetAddress string      `json:"street_address"`
	ZipCode       string      `json:"zip_code"`
	State         string      `json:"state"`
	Country       string      `json:"country"`
	Coordinates   Coordinates `json:"coordinates"`
}

type Coordinates struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type CreditCard struct {
	CCNumber string `json:"cc_number"`
}

type Subscription struct {
	Plan          string `json:"plan"`
	Status        string `json:"status"`
	PaymentMethod string `json:"payment_method"`
	Term          string `json:"term"`
}

// Initialize GORM and connect to the database
func initDB() {
	// Set up PostgreSQL connection (adjust with your credentials)
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
	con, err := gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Error opening database:", err)
	}

	// Auto-migrate the User struct to create the table
	con.AutoMigrate(&User{})
	db = con
}

// Random User Data from random-data-api.com
func getRandomUserData() (*User, error) {
	resp, err := http.Get("https://random-data-api.com/api/users/random_user?size=1")
	if err != nil {
		return nil, fmt.Errorf("error fetching random user data: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch random data, status code: %d", resp.StatusCode)
	}

	var users []UserDto
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		return nil, fmt.Errorf("error decoding random user data: %v", err)
	}

	return &User{
		FirstName:         users[0].FirstName,
		LastName:          users[0].LastName,
		Username:          users[0].Username,
		Email:             users[0].Email,
		Avatar:            users[0].Avatar,
		PhoneNumber:       users[0].PhoneNumber,
		DateOfBirth:       time.Now(),
		AddressCountry:    users[0].Address.Country,
		AddressCity:       users[0].Address.City,
		AddressStreetName: users[0].Address.StreetName,
		AddressStreetAddr: users[0].Address.StreetAddress,
	}, nil
}

// Create a new user (with or without random data)
func createUser(c *fiber.Ctx) error {
	var user User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid input")
	}

	// If some fields are missing, fetch random data
	if user.FirstName == "" || user.LastName == "" || user.Username == "" || user.Email == "" {
		randomUser, err := getRandomUserData()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error generating random user data")
		}
		user = *randomUser
	}

	// Insert into the database
	if err := db.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error inserting user into database")
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

// Get a user by ID (user ID passed in body)
func getUser(c *fiber.Ctx) error {
	var req struct {
		UserID int `json:"user_id"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}

	var user User
	if err := db.First(&user, req.UserID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).SendString("User not found")
	}

	return c.JSON(user)
}

// Update an existing user (user ID and data passed in body)
func updateUser(c *fiber.Ctx) error {
	var user User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid input")
	}

	// Find the existing user
	var existingUser User
	if err := db.First(&existingUser, user.UserID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).SendString("User not found")
	}

	// Update the existing user
	if err := db.Model(&existingUser).Updates(user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error updating user")
	}

	return c.JSON(existingUser)
}

// Delete a user by ID (user ID passed in body)
func deleteUser(c *fiber.Ctx) error {
	var req struct {
		UserID int `json:"user_id"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}

	// Delete the user by ID
	if err := db.Delete(&User{}, req.UserID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error deleting user")
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// Get users with filtering and sorting
func getUsers(c *fiber.Ctx) error {
	// Extract query parameters for filtering and sorting
	username := c.Query("username", "")
	firstName := c.Query("first_name", "")
	lastName := c.Query("last_name", "")
	sortBy := c.Query("sort_by", "username")  // Default sorting by username
	sortOrder := c.Query("sort_order", "asc") // Default sorting order is ascending

	// Build the query based on filters
	var users []User
	query := db.Model(&User{})

	// Apply filters based on query parameters
	if username != "" {
		query = query.Where("username LIKE ?", "%"+username+"%")
	}
	if firstName != "" {
		query = query.Where("first_name LIKE ?", "%"+firstName+"%")
	}
	if lastName != "" {
		query = query.Where("last_name LIKE ?", "%"+lastName+"%")
	}

	// Apply sorting based on query parameters
	if sortOrder == "desc" {
		query = query.Order(sortBy + " desc")
	} else {
		query = query.Order(sortBy + " asc")
	}

	// Fetch the filtered and sorted users
	if err := query.Find(&users).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error retrieving users")
	}

	return c.JSON(users)
}

func main() {
	// Initialize database connection
	initDB()
	defer db.Close()

	// Create a new Fiber app
	app := fiber.New()

	// Routes
	app.Post("/create-user", createUser) // Create new user
	app.Get("/get-users", getUsers)      // Get users with filtering and sorting
	app.Post("/get-user", getUser)       // Get user by ID (POST request with user_id)
	app.Post("/update-user", updateUser) // Update user (POST request with user data including user_id)
	app.Post("/delete-user", deleteUser) // Delete user (POST request with user_id)

	// Start the server
	log.Fatal(app.Listen(":8080"))
}
