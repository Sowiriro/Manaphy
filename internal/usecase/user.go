package usecase

type UserUseCase struct {
	userRepo *Repository.UserRepository
}

type UserUseCaseI interface {
	All(ctx context.Context, input InputPort)(OutputPort, error)
	Get(ctx context.Context, input InputPort)(OutputPort, error)
	Create(ctx context.Context, input InputPort)(OutputPort, error)
	Update()
	Delete()
}

func NewUserUsecase(userRepo repository.UserRepository) UserUseCaseI {
	return &UserUsecase{
		userRepo,
	}
}

func (u *Usecase) All(ctx context.Context, input InputPort)(OutputPort, error) {
	req, ok := input.(entity.UserPostRequest)
	if !ok {
		return nil, entity.ErrFailedCastingInput
	}

	user, err := u.userRepo.All(ctx, req.Name)
	if err != nil {
		return nil, err
	}

	return entity.UserPostResponses{
		ID: user.ID,
		Name: user.Name,
	},
}

func (u *Usecase) Get() {
	return
}

func (u *Usecase) Create() {
	return
}

func (u *Usecase) Update() {
	return
}

func (u *Usecase) Delete() {
	return
}
