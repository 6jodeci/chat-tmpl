package user

import (
	"context"
	"time"
)

type service struct {
	Repository
	timeout time.Duration
}

func NewService(repository Repository) Service {
	return &service{
		Repository: repository,
		timeout:    time.Duration(2) * time.Second,
	}
}

func (s *service) CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	// TODO Hash password

	user := &User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}
	rspCreateUser, err := s.Repository.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	res := &CreateUserResponse{
		ID:       rspCreateUser.ID,
		Username: responseCreateUser.Username,
		Email:    responseCreateUser.Email,
	}
	return res, nil
}
