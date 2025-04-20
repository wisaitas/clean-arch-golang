package usecases

import (
	"github.com/wisaitas/clean-arch-golang/internal/standard-service/domain/repositories"
)

type UserUsecase interface {
}

type userUsecase struct {
	userRepository repositories.UserRepository
}

func NewUserUsecase(
	userRepository repositories.UserRepository,
) UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}
