package database

import (
	"os"

	queries "SocialNetworkBackend/app/queries/mongo"

	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
)

// Queries struct for collect all app queries.
type Queries struct {
	*queries.UserQueries // load queries from User model
	// *queries.BookQueries // load queries from Book model
}

// OpenDBConnection func for opening database connection.
func OpenDBConnection() (*Queries, error) {
	// Define Database connection variables.
	var _ *sqlx.DB
	var dbNoSql *mongo.Client
	var err error


	// Get DB_TYPE value from .env file.
	dbType := os.Getenv("DB_TYPE")

	// Define a new Database connection with right DB type.
	switch dbType {
	case "pgx":
		_, err = PostgreSQLConnection()
	case "mongodb":
		dbNoSql, err = MongoDBConnection()
	}

	if err != nil {
		return nil, err
	}

	// Define Queries struct for collect all app queries(sql or NoSql).
	return &Queries{
		UserQueries: &queries.UserQueries{Collection: dbNoSql.Database(os.Getenv("DB_NAME")).Collection("users")},
		// BookQueries: &queries.BookQueries{DB: dbsql},
	}, nil
}