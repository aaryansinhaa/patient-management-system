package user_repo

// Package user_repo provides the implementation of the UserRepository interface
import (
	"database/sql"

	"github.com/aaryansinhaa/patient-management-system/internals/model"
)

type UserStorage struct {
	connection *sql.DB
}

func NewUserStorage(db *sql.DB) *UserStorage {
	return &UserStorage{
		connection: db,
	}
}

func (s *UserStorage) CreateUser(user model.User) error {
	query := `INSERT INTO users (id, name, role, username, password, phone_number) 
	          VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := s.connection.Exec(query, user.ID, user.Name, user.Role, user.Username, user.Password, user.PhoneNumber)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserStorage) DeleteUser(id string) (*model.User, error) {
	query := `DELETE FROM users WHERE id = $1 
	RETURNING id, name, role, username, password, phone_number`
	row := s.connection.QueryRow(query, id)

	var user model.User
	err := row.Scan(&user.ID, &user.Name, &user.Role, &user.Username, &user.Password, &user.PhoneNumber)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // User not found
		}
		return nil, err
	}
	return &user, nil
}

func (s *UserStorage) UpdateUser(user model.User) (*model.User, error) {
	query := `UPDATE users SET name = $1, role = $2, username = $3, password = $4, phone_number = $5, updated_at = NOW() 
	          WHERE id = $6
			  RETURNING id, name, role, username, password, phone_number`
	row := s.connection.QueryRow(query, user.Name, user.Role, user.Username, user.Password, user.PhoneNumber, user.ID)

	var updatedUser model.User
	err := row.Scan(&updatedUser.ID, &updatedUser.Name, &updatedUser.Role, &updatedUser.Username, &updatedUser.Password, &updatedUser.PhoneNumber)
	if err != nil {
		return nil, err
	}
	return &updatedUser, nil
}

func (s *UserStorage) GetUserByID(id string) (*model.User, error) {
	query := `SELECT id, name, role, username, password, phone_number FROM users WHERE id = $1`
	row := s.connection.QueryRow(query, id)

	var user model.User
	err := row.Scan(&user.ID, &user.Name, &user.Role, &user.Username, &user.Password, &user.PhoneNumber)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // User not found
		}
		return nil, err
	}
	return &user, nil
}

func (s *UserStorage) GetUserByUsername(username string) (*model.User, error) {
	query := `SELECT id, name, role, username, password, phone_number FROM users WHERE username = $1`
	row := s.connection.QueryRow(query, username)

	var user model.User
	err := row.Scan(&user.ID, &user.Name, &user.Role, &user.Username, &user.Password, &user.PhoneNumber)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // User not found
		}
		return nil, err
	}
	return &user, nil
}

func (s *UserStorage) GetUserIdByUsername(username string) (string, error) {
	query := `SELECT id FROM users WHERE username = $1`
	row := s.connection.QueryRow(query, username)

	var userID string
	err := row.Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil // User not found
		}
		return "", err
	}
	return userID, nil
}

func (s *UserStorage) GetAllUsers() ([]model.User, error) {
	query := `SELECT id, name, role, username, password, phone_number FROM users`
	rows, err := s.connection.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.ID, &user.Name, &user.Role, &user.Username, &user.Password, &user.PhoneNumber)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserStorage) GetAllUsersByRole(role string) ([]model.User, error) {
	query := `SELECT id, name, role, username, password, phone_number FROM users WHERE role = $1`
	rows, err := s.connection.Query(query, role)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.ID, &user.Name, &user.Role, &user.Username, &user.Password, &user.PhoneNumber)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserStorage) GetUsernameAndPasswordById(id string) (string, string, error) {
	query := `SELECT username, password FROM users WHERE id = $1`
	row := s.connection.QueryRow(query, id)

	var username, password string
	err := row.Scan(&username, &password)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", "", nil // User not found
		}
		return "", "", err
	}
	return username, password, nil
}

func (s *UserStorage) GetUserByPhoneNumber(phoneNumber string) (*model.User, error) {
	query := `SELECT id, name, role, username, password FROM users WHERE phone_number = $1`
	row := s.connection.QueryRow(query, phoneNumber)

	var user model.User
	err := row.Scan(&user.ID, &user.Name, &user.Role, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // User not found
		}
		return nil, err
	}
	return &user, nil
}
