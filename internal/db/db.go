package db

import (
	"fmt"
	"log"

	"github.com/AhmedMohamed800/go-backend-template/config"
	_ "github.com/go-sql-driver/mysql" // MySQL driver
	_ "github.com/jackc/pgx/v4"        // PostgreSQL driver (for pgx)
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

type Storage struct {
	DB *sqlx.DB
}

func NewStorage(DBConfig config.DBConfig) (*Storage, error) {
    dbType := DBConfig.DBType // Fetch DB_TYPE from the config

    var db *sqlx.DB
    var err error

    switch dbType {
    case "postgres":
        dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
            DBConfig.DBUser,
            DBConfig.DBPass,
            DBConfig.DBHost,
            DBConfig.DBPort,
            DBConfig.DBType) // PostgreSQL DB name or schema

        db, err = sqlx.Connect("postgres", dbURL)
        if err != nil {
            return nil, fmt.Errorf("unable to connect to PostgreSQL: %v", err)
        }
        fmt.Println("Connected to PostgreSQL")

    case "sqlite":
        dbFile := DBConfig.DBFile // Assuming DBFile is the path to the SQLite database file
        db, err = sqlx.Connect("sqlite3", dbFile)
        if err != nil {
            return nil, fmt.Errorf("unable to connect to SQLite: %v", err)
        }
        fmt.Println("Connected to SQLite")

    case "mysql":
        dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
            DBConfig.DBUser,
            DBConfig.DBPass,
            DBConfig.DBHost,
            DBConfig.DBPort,
            DBConfig.DBType) // MySQL DB name or schema

        db, err = sqlx.Connect("mysql", dbURL)
        if err != nil {
            return nil, fmt.Errorf("unable to connect to MySQL: %v", err)
        }
        fmt.Println("Connected to MySQL")

    default:
        return nil, fmt.Errorf("unsupported database type: %s", dbType)
    }

    return &Storage{DB: db}, nil
}

// Close closes the database connection
func (s *Storage) Close() {
    if err := s.DB.Close(); err != nil {
        log.Println("Error closing database connection:", err)
    }
}