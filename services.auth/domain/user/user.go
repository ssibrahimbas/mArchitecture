package user

type User struct {
	UUID      string
	Email     string
	Password  []byte
	IsActive  bool
	CreatedAt string
	UpdatedAt string
}

func (u *User) SetPassword(password []byte) {
	u.Password = password
}

func (u *User) CleanPassword() {
	u.Password = nil
}
