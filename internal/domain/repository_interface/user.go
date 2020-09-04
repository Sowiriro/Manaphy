package repository_interface

type UserRepositoryI interface {
	All(entity.Users) (entity.Users, error)
}
