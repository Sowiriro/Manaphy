package repository

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &userRepository{db: db}
}

func (repo *userRepository) All(entity.Users) (entity.Users, error) {
	return
}
