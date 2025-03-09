package services

import (
	"context"
	"tethys-go/internal/core/domain"
	"tethys-go/internal/core/ports"
)

type UserService struct {
	state *ports.ApiState
	repo  ports.IUserRepository
}

func NewUserService(state *ports.ApiState, repo ports.IUserRepository) ports.IUserService {
	return &UserService{state: state, repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, dto *ports.CreateUserDTO) (*domain.UserModel, error) {
	if err := dto.Validate(); err != nil {
		return nil, err
	}
	return s.repo.CreateUser(ctx, s.state.PgExec, dto)
}

func (s *UserService) GetUser(ctx context.Context, id domain.UserID) (*domain.UserModel, error) {
	return s.repo.GetUser(ctx, s.state.PgExec, id)
}

func (s *UserService) GetUserByVkID(ctx context.Context, vkID domain.UserVkID) (*domain.UserModel, error) {
	return s.repo.GetUserByVkID(ctx, s.state.PgExec, vkID)
}
