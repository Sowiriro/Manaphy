package entity

import(
	"time"
	"errors"
)

var (
	ErrUserNotFound = errors.New("user is not found")
)

type (
	User struct {
		ID       string    `json:"id"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Time   []Time
	}
	Users []User
)

type Time struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt	time.Time `json:"deleted_at"`
}

type (
	UserPostRequest struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
)
	UserPostRequests []UserPostRequest
}

type (
	UserPostResponse struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
)
	UserPostResponses []UserPostResponse
}

func (u *User)NewUser() User {
	uid, _ := uuid.NewUUID()
	u.UID = uid.String()
	reuturn User {
		UID: uid,
		Name: u.Name,
		Email: u.Email,
		Password: u.Password,
	}
}

func (u *User)InitUser() User {
	return User{
		UID: u.UID,
		Name: u.Name,
		Email: u.Email,
		Password: u.Password,
	}
}