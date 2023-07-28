package service

import (
	"fmt"
)

func (s *service) CreateUser(user *User) error {
	err := s.userRepo.Create(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetUser(id string) (*User, error) {
	user, err := s.userRepo.Get(id)
	if err != nil {
		return nil, fmt.Errorf("cannot get user: %v", err)
	}
	return user, nil
}

func (s *service) UpdateUser(id string, user *User) error {
	err := s.userRepo.Update(id, user)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) DeleteUser(id string) error {
	err := s.userRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
