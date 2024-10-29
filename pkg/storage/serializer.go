package storage

import (
	er "github.com/MiLara8888/caching_web_server/pkg/errors"
)

type RegisterSerializer struct {
	TokenAdmin string `json:"token"`
	Login      string `json:"login"`
	Password   string `json:"pswd"`
}

type ErrSerializer struct {
	Error er.Error `json:"error"`
}


// login
// Минимальная длина 8, латиница и цифры
// pswd
// •минимальная длина 8,
// •минимум 2 буквы в разных регистрах
// •минимум 1 цифра
// •минимум 1 символ (не буква и не цифра)
func (r *RegisterSerializer) Valid() bool {
	if r.Login == "" || r.TokenAdmin == "" || r.Password == "" {
		return false
	}
	return true
}
