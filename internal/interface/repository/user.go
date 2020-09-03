package repository


type userRepository sturuct {
	sqlDB		*sql.DB 
}

func NewUserRepository(db *sql.DB) respository.UserRepository{
	return &userRepository{
		db: db,
	}
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

func (repo *userRepository) Get(id int)(entity.User, error){
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