package usecase

type UserUseCase struct {
	userRepo *Repository.UserRepository
}

type UserUseCaseI interface {
	All()
	Get()
	Create()
	Update()
	Delete()
}

func NewUserUsecase(userRepo repository.UserRepository) UserUseCaseI {
	return &UserUsecase{
		userRepo,
	}
}

func (u *Usecase) All() {
	return
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
