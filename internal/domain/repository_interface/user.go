package repository_interface

type UserRepositoryI interface {
	All(context.Context, entity.Users) (entity.Users, error)
	GetByID(context.Context, int)(entity.User, error)
	Create(context.Context, entity.User)error
	Update(context.Context, entity.User)error
	Delete(context.Context, entity.User)error
}
