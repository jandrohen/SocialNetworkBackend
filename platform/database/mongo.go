package database

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"SocialNetworkBackend/pkg/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDBConnection func for connection to MongoDB database.
func MongoDBConnection() (*mongo.Client, error) {
	// Define database connection settings.
	maxConn, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))
	maxIdleConn, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
	maxLifetimeConn, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME_CONNECTIONS"))

	// Build MongoDB connection URL.
	mongoDbConnURL, err := utils.ConnectionURLBuilder("mongo")
	if err != nil {
		return nil, err
	}

	// Set database connection settings:
	// 	- SetMaxPoolSize: the maximum number of connections in the connection pool
	// 	- SetMaxConnIdleTime: the maximum number of milliseconds that a connection can remain idle in the pool
	// 	- SetMaxConnIdleTime: the maximum number of milliseconds that a connection can remain idle in the pool
	// Set options for MongoDB connection.
	options := options.Client().ApplyURI(mongoDbConnURL)
	options.SetMaxPoolSize(uint64(maxConn))
	options.SetMaxConnIdleTime(time.Duration(maxIdleConn))
	options.SetMaxConnIdleTime(time.Duration(maxLifetimeConn))

	// Define database connection for MongoDB.
	db, err := mongo.Connect(context.Background(), options)
	if err != nil {
		return nil, fmt.Errorf("error, not connected to database, %w", err)
	}


	

	// Try to ping database.
	if err := db.Ping(context.Background(), nil); err != nil {
		defer db.Disconnect(context.Background()) // close database connection
		return nil, fmt.Errorf("error, not sent ping to database, %w", err)
	}

	return db, nil
}
