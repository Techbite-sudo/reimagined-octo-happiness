package users

import(
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"lms_backend/graph/model"
	"regexp"
	"unicode"

	_ "github.com/lib/pq"
)

func FindUserByEmail(email string) (*model.User, error) {
	// Open a connection to the database
	db, err := sql.Open("postgres", "postgres://user:password@localhost/database?sslmode=disable")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Query the database for a user with the given email
	var user model.User
	err = db.QueryRow("SELECT id, email, password, name FROM users WHERE email = $1", email).Scan(&user.ID, &user.Email, &user.Password, &user.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
func AddUser(user *model.User) error {
	// Open a connection to the database
	db, err := sql.Open("postgres", "postgres://user:password@localhost/database?sslmode=disable")
	if err != nil {
		return err
	}
	defer db.Close()

	// Insert the user into the database
	_, err = db.Exec("INSERT INTO users (id, email, password, name) VALUES ($1, $2, $3, $4)", user.ID, user.Email, user.Password, user.Name)
	if err != nil {
		return err
	}

	return nil
}
func GenerateID() string {
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
func ValidateEmail(email string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return re.MatchString(email)
}
func ValidateName(name string) bool {
	if len(name) < 2 || len(name) > 100 {
		return false
	}
	for _, c := range name {
		if !unicode.IsLetter(c) && !unicode.IsSpace(c) {
			return false
		}
	}
	return true
}
func HashPassword(password string) string {
	h := sha256.New()
	h.Write([]byte(password))
	return hex.EncodeToString(h.Sum(nil))
}
func FindUserByID(id string) (*model.User, error) {
	// Open a connection to the database
	db, err := sql.Open("postgres", "postgres://user:password@localhost/database?sslmode=disable")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Query the database for a user with the given ID
	var user model.User
	err = db.QueryRow("SELECT id, email, password, name FROM users WHERE id = $1", id).Scan(&user.ID, &user.Email, &user.Password, &user.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
func UpdateUser(user *model.User) error {
	// Open a connection to the database
	db, err := sql.Open("postgres", "postgres://user:password@localhost/database?sslmode=disable")
	if err != nil {
		return err
	}
	defer db.Close()

	// Update the user in the database
	_, err = db.Exec("UPDATE users SET email = $1, password = $2, name = $3 WHERE id = $4", user.Email, user.Password, user.Name, user.ID)
	if err != nil {
		return err
	}

	return nil
}
func DeleteUser(user *model.User) error {
	// Open a connection to the database
	db, err := sql.Open("postgres", "postgres://user:password@localhost/database?sslmode=disable")
	if err != nil {
		return err
	}
	defer db.Close()

	// Delete the user from the database
	_, err = db.Exec("DELETE FROM users WHERE id = $1", user.ID)
	if err != nil {
		return err
	}

	return nil
}
func ListUsers() ([]*model.User, error) {
	// Open a connection to the database
	db, err := sql.Open("postgres", "postgres://user:password@localhost/database?sslmode=disable")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Query the database for a list of users
	rows, err := db.Query("SELECT id, email, password, name FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate through the rows and create a list of users
	var users []*model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Email, &user.Password, &user.Name); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}