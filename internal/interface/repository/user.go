package repository

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &userRepository{db: db}
}

func (repo *userRepository) All(entity.Users)(entity.Users, error){
	entities := entity.Users{}

	// users, err := dbmodels.ForumQuestions().All(context.Background(), u.db)
	// if err != nil {
	// 	return nil, err
	// }

	for i := 0; i < len(users); i++ {
		entity := entity.User {
			ID:  	users[i].ID,
			Name: users[i].Name,
		}
		entities = append(entities, entity)
	}

	return entities, nil
}

func (repo *userRepository) GetByID(id int)(entity.User, error){
	return 
}

func (repo *userRepository) Create(entity.User)error{
	return 
}

func (repo *userRepository) Update(entity.User)error {
	return 
}

func (repo *userRepository) Delete(entity.User)error {
	return 
}

