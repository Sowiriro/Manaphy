package repository_interface

type UserRepositoryI interface {
	All(entity.Users) (entity.Users, error)
	GetByID(int)(entity.User, error)
	Create(entity.User)error
	Update(entity.User)error
	Delete(entity.User)error
}
