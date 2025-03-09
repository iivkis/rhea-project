package ports

import (
	"context"
	"fmt"
	"tethys-go/internal/core/domain"
)

type CreateUserDTO struct {
	VkID domain.UserVkID `json:"vk_id"`
}

func (d *CreateUserDTO) Validate() error {
	if d.VkID == 0 {
		return fmt.Errorf("%w: vk_id", ErrInvalidData)
	}
	return nil
}

type IUserService interface {
	CreateUser(ctx context.Context, dto *CreateUserDTO) (*domain.UserModel, error)
	GetUser(ctx context.Context, id domain.UserID) (*domain.UserModel, error)
	GetUserByVkID(ctx context.Context, vkID domain.UserVkID) (*domain.UserModel, error)
}

type IUserRepository interface {
	CreateUser(ctx context.Context, exec PgExecutor, dto *CreateUserDTO) (*domain.UserModel, error)
	GetUser(ctx context.Context, exec PgExecutor, id domain.UserID) (*domain.UserModel, error)
	GetUserByVkID(ctx context.Context, exec PgExecutor, vkID domain.UserVkID) (*domain.UserModel, error)
}
