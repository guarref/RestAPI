package userService

type UserService struct {
	urepo UserRepository
}

func NewUserService(urepo UserRepository) *UserService {
	return &UserService{urepo: urepo}
}

func (s *UserService) CreateUser(user User) (User, error) {
	return s.urepo.CreateUser(user)
}

func (s *UserService) GetAllUsers() ([]User, error) {
	return s.urepo.GetAllUsers()
}

func (s *UserService) UpdateUserByID(id uint, user User) (User, error) {
	return s.urepo.UpdateUserByID(id, user)
}

func (s *UserService) DeleteUserByID(id uint) error {
	return s.urepo.DeleteUserByID(id)
}
