package storage

import (
	"github.com/ramonberrutti/jumpshot-ext-authz/internal/storage/engine"
)

type Storage interface {
	Sessions() Sessions
}

type storage struct {
	e engine.Engine
}

// New storage with engine.
func NewStorage(e engine.Engine) Storage {
	return &storage{e}
}

func (s *storage) Sessions() Sessions {
	return &sessions{s.e}
}
