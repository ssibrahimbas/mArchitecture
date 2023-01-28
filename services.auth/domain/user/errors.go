package user

import "github.ssibrahimbas/mArchitecture/shared/i18n"

type Errors interface {
	NotFound(email string) *i18n.I18nError
	AlreadyExists(email string) *i18n.I18nError
	Failed(operation string) *i18n.I18nError
}

type userErrors struct{}

func newUserErrors() Errors {
	return &userErrors{}
}

func (e *userErrors) NotFound(email string) *i18n.I18nError {
	return i18n.NewError("error_user_not_found", i18n.P{"Email": email})
}

func (e *userErrors) AlreadyExists(email string) *i18n.I18nError {
	return i18n.NewError("error_user_already_exist", i18n.P{"Email": email})
}

func (e *userErrors) Failed(operation string) *i18n.I18nError {
	return i18n.NewError("error_user_failed", i18n.P{"Operation": operation})
}
