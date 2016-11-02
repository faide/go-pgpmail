package imap

import (
	"github.com/emersion/go-imap/backend"
)

type Backend struct {
	backend.Backend

	unlock UnlockFunction
}

func New(be backend.Backend, unlock UnlockFunction) *Backend {
	return &Backend{be, unlock}
}

func (be *Backend) Login(username, password string) (backend.User, error) {
	if u, err := be.Backend.Login(username, password); err != nil {
		return nil, err
	} else if kr, err := be.unlock(username, password); err != nil {
		return nil, err
	} else {
		return &user{u, kr}, nil
	}
}
