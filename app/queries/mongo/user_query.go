package queries

import (
	"SocialNetworkBackend/app/models"
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserQueries struct for queries from User model.
type UserQueries struct {
	*mongo.Collection
}

// GetUserByID query for getting one User by given ID.
func (q *UserQueries) GetUserByID(id uuid.UUID) (models.User, error) {
	// Define User variable.
	var user models.User

	// Define filter
	filter := bson.M{"id": id}

	// Send query to database.
	err := q.Collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		// Return empty object and error.
		return user, err
	}

	// Return query result.
	return user, nil
}

// GetUserByEmail query for getting one User by given Email.
func (q *UserQueries) GetUserByEmail(email string) (models.User, error) {
	// Define User variable.
	var user models.User

	// Define filter
	filter := bson.M{"email": email}

	// Send query to database.
	err := q.Collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		// Return empty object and error.
		return user, err
	}

	// Return query result.
	return user, nil
}

// CreateUser query for creating a new user by given email and password hash.
func (q *UserQueries) CreateUser(u *models.User) error {
	// Define user document
	user := bson.M{
		"id":           u.ID,
		"created_at":   u.CreatedAt,
		"updated_at":   u.UpdatedAt,
		"email":        u.Email,
		"passwordHash": u.PasswordHash,
		"userStatus":   u.UserStatus,
		"userRole":     u.UserRole,
	}

	// Send query to database.
	_, err := q.Collection.InsertOne(context.TODO(), user)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}