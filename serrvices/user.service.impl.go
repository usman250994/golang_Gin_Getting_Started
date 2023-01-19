package service

type UserServiceImpl struct {
	usercollection *mogngo.Collection
	ctx context.Context
}

func (u *UserServiceImpl) CreateUser(user *models.User) error {
	return nil
}
func (u *UserServiceImpl) GetUser(name *string ) (*models.User,error) error {
	return nil,nil
}

func (u *UserServiceImpl) GetAll() []*models.User error {
	return nil
}

func (u *UserServiceImpl) UpdateUser(user *models.User)  error {
	return nil
}

func (u *UserServiceImpl) DeleteUser(name *string)  error {
	return nil
}