package repository

import (
	"context"
	"tethys-go/internal/core/domain"
	"tethys-go/internal/core/ports"
)

type PgUserRepository struct{}

func NewPgUserRepository() ports.IUserRepository {
	return &PgUserRepository{}
}

func (p *PgUserRepository) CreateUser(ctx context.Context, executor ports.PgExecutor, dto *ports.CreateUserDTO) (*domain.UserModel, error) {
	const query = `
		INSERT INTO users (
			vk_id
		) VALUES ($1)
		RETURNING id, vk_id
	`
	var user domain.UserModel
	err := executor.QueryRow(ctx, query, dto.VkID).
		Scan(&user.ID, &user.VkID)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (p *PgUserRepository) GetUserByVkID(ctx context.Context, executor ports.PgExecutor, vkID domain.UserVkID) (*domain.UserModel, error) {
	const query = `
		SELECT id, vk_id
		FROM users
		WHERE vk_id = $1
	`
	var user domain.UserModel
	err := executor.QueryRow(ctx, query, vkID).
		Scan(&user.ID, &user.VkID)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (p *PgUserRepository) GetUser(ctx context.Context, exec ports.PgExecutor, id domain.UserID) (*domain.UserModel, error) {
	const query = `
		SELECT id, vk_id
		FROM users
		WHERE id = $1
	`
	var user domain.UserModel

	err := exec.QueryRow(ctx, query, id).
		Scan(&user.ID, &user.VkID)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
