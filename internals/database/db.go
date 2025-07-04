package database

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/aaryansinhaa/patient-management-system/internals/config"
	_ "github.com/lib/pq" // PostgreSQL driver
)

type DatabaseConnection struct {
	Connection *sql.DB
}

func LoadPSqlDb(config *config.DatabaseConfig) (*DatabaseConnection, error) {
	connStr := "host=" + config.Host + " port=" + strconv.Itoa(config.Port) +
		" dbname=" + config.DbName + " user=" + config.User + " password=" + config.Password + " sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping db: %w", err)
	}

	// Create gender enum if not exists
	_, err = db.Exec(`DO $$
	BEGIN
		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'gender_type') THEN
			CREATE TYPE gender_type AS ENUM ('male', 'female', 'other');
		END IF;
	END
	$$;`)
	if err != nil {
		return nil, fmt.Errorf("failed to create gender enum: %w", err)
	}

	// Create users table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id UUID PRIMARY KEY,
		name TEXT NOT NULL,
		role TEXT NOT NULL CHECK (role IN ('doctor', 'receptionist')),
		username TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL,
		phone_number TEXT UNIQUE NOT NULL,
		created_at TIMESTAMPTZ DEFAULT NOW(),
		updated_at TIMESTAMPTZ DEFAULT NOW() 
	);`)
	if err != nil {
		return nil, fmt.Errorf("failed to create users table: %w", err)
	}

	// Create patient table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS patient (
		id UUID PRIMARY KEY,
		name TEXT NOT NULL,
		age INT NOT NULL CHECK (age >= 0),
		gender gender_type DEFAULT 'other',
		phone_number TEXT UNIQUE NOT NULL,
		created_at TIMESTAMPTZ DEFAULT NOW(),
		updated_at TIMESTAMPTZ DEFAULT NOW()
	);`)
	if err != nil {
		return nil, fmt.Errorf("failed to create patient table: %w", err)
	}

	// Create diagnosis table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS diagnosis (
		id SERIAL PRIMARY KEY,
		patient_id UUID NOT NULL REFERENCES patient(id) ON DELETE CASCADE,
		doctor_id UUID NOT NULL REFERENCES users(id) ON DELETE SET NULL,
		description TEXT NOT NULL,
		created_at TIMESTAMPTZ DEFAULT NOW()
	);`)
	if err != nil {
		return nil, fmt.Errorf("failed to create diagnosis table: %w", err)
	}

	return &DatabaseConnection{Connection: db}, nil
}
