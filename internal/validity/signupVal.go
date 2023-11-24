package validity

import (
	"net/url"
	"regexp"
	// t "forum/internal/types"
)

var RegexEmail = regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)

type Form struct {
	url.Values
}

func GetForm(form url.Values) *Form {
	return &Form{
		form,
	}
}

func (f *Form) CheckEmail() bool {
	emailValue := f.Get("email") ////////получает данные по ключу из form method=post
	if emailValue == "" {
		return false
	}

	if !RegexEmail.MatchString(emailValue) {
		return false
	}
	if len(emailValue) > 50 {
		return false
	}

	return true
}

func (f *Form) CheckName() bool {
	username := f.Get("username")
	return username != ""
}

func (f *Form) CheckPassword() bool {
	password := f.Get("password")
	if password == "" {
		return false
	}

	if len(password) < 8 {
		return false
	}
	if len(password) > 50 {
		return false
	}
	if f.PasswordRepeat(password) {
		return true
	}
	return false
	// return true
}

func (f *Form) PasswordRepeat(firstPass string) bool {
	secondPass := f.Get("password2")
	if secondPass == "" {
		return false
	}
	if firstPass != secondPass {
		return false
	}
	return true
}
