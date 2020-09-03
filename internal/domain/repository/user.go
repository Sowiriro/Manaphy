package repository

import(
	"github.com/Sowiriro/Manaphy/domain"
)


type UserRepositoryI interface {
	All(entity.Users))(entity.Users, error)
	Get(id int)(entity.User, error)
	GetByID(id int)(entity.User, error)
	Create(entity.User)error
	Update(entity.User)error
	Delete(entity.User)error
}

