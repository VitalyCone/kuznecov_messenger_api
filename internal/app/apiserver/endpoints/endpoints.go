package endpoints

import (
	"github.com/VitalyCone/kuznecov_messenger_api/internal/app/store"
)

type Endpoints struct {
	store *store.Store
}

func NewEndpoints(s *store.Store) *Endpoints{
	return &Endpoints{
		store: s,
	}
}