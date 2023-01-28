package user

import "github.ssibrahimbas/mArchitecture/shared/i18n"

type Factory struct {
	Errors Errors
}

func NewFactory() Factory {
	return Factory{
		Errors: newUserErrors(),
	}
}

func (f Factory) IsZero() bool {
	return f.Errors == nil || f.Errors.NotFound("") == nil
}

func (f Factory) NewUser(email string) *User {
	return &User{
		Email:    email,
		IsActive: true,
	}
}

func (f Factory) Unmarshal(uuid string, email string, isActive bool) *User {
	return &User{
		UUID:     uuid,
		Email:    email,
		IsActive: isActive,
	}
}

func (f Factory) Validate(u *User) *i18n.I18nError {
	if err := f.validateEmail(u.Email); err != nil {
		return err
	}
	return nil
}

func (f Factory) validateEmail(email string) *i18n.I18nError {
	if email == "" {
		return i18n.NewError("error_user_email_empty")
	}
	return nil
}
