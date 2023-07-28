package service

type UserRepo interface {
	Create(usr *User) error
	Get(id string) (*User, error)
	Update(id string, user *User) error
	Delete(id string) error
}

type Service interface {
	CreateUser(usr *User) error
	GetUser(id string) (*User, error)
	UpdateUser(id string, user *User) error
	DeleteUser(id string) error
}
