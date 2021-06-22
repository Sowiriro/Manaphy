package usecase

import (
	"github.com/3-shake/pcidss-dashbord-api/internal/domain/entity"
	"github.com/3-shake/pcidss-dashbord-api/internal/domain/repository"
	"golang.org/x/net/context"
)

type UpdateUser struct {
	userRepo repository.UserRepository
}

//func NewUserUseCase(
//	userRepository repository.UserRepository,
//) UserUseCase {
//	return &CreateUser{
//		userRepository,
//	}
//}

func (s *UpdateUser) Execute(ctx context.Context, input InputPort) (OutputPort, error) {
	req, ok := input.(entity.UserPostRequest)
	if !ok {
		return nil, entity.ErrFailedCastingUseCaseInput
	}

	user := entity.InitUser(
		req.ID ,
		req.Name,
		req.Email,
		req.Password,
	)

	resp, err := s.userRepo.All(ctx, user)
	if err != nil {
		return nil, entity.ErrUserNotFound
	}

	return entity.UserPostResponse{
		Name: resp.Name,
		Email:	resp.Email,
	}, nil
}

