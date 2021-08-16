package facede

type IUser interface {
	Login(int, int) (*User, error)
	Register(int, int) (*User, error)
}

type IUserFacede interface {
	LoginOrRegister(int, int) (*User, error)
}

type User struct {
	Name string
}

type UserService struct{}

func (u UserService) Login(phone int, code int) (*User, error) {
	return &User{Name: "login test"}, nil
}

func (u UserService) Register(phone int, code int) (*User, error) {
	return &User{Name: "register test"}, nil
}

func (u UserService) LoginOrRegister(phone int, code int) (*User, error) {
	user, err := u.Login(phone, code)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return user, nil
	}

	return u.Register(phone, code)
}
