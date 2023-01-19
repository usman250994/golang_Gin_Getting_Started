package services

type UserService interface {
	CreateUser(*models.User) error
	GetUser(*string) (*models.User,error)
	GetAll()[*models.User,error]
	Delete(*string) error
}