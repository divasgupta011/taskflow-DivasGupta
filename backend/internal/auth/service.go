package auth

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo      *Repository
	jwtSecret string
}

func NewService(repo *Repository, jwtSecret string) *Service {
	return &Service{repo: repo, jwtSecret: jwtSecret}
}

func (s *Service) Register(ctx context.Context, name, email, password string) (*User, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return nil, err
	}

	user := &User{
		Name:     name,
		Email:    email,
		Password: string(hashed),
	}

	err = s.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) Login(ctx context.Context, email, password string) (string, *User, error) {
	user, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return "", nil, errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", nil, errors.New("invalid credentials")
	}

	token, err := GenerateJWT(user.ID, user.Email, s.jwtSecret)
	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}
