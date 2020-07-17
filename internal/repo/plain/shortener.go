package plain

import (
	"sync"

	"github.com/toms1441/urlsh/internal/repo"
	"github.com/toms1441/urlsh/internal/shortener"
)

// This package is a mock that implements all repositories.
// It's meant to be used in tests mainly.

type shortenerRepository struct {
	mtx sync.Mutex
	db  map[string]shortener.Model
}

func NewShortenerRepository() (shortener.Repository, error) {
	sr := &shortenerRepository{}
	sr.db = make(map[string]shortener.Model)

	return sr, nil
}

func (sr *shortenerRepository) Create(urlsh shortener.Model) error {
	if sr.db == nil {
		return repo.ErrInvalidRepository
	}

	sr.mtx.Lock()
	defer sr.mtx.Unlock()

	if len(urlsh.ID) == 0 {
		return repo.ErrInvalidID
	}

	_, ok := sr.db[urlsh.ID]
	if ok {
		return repo.ErrShortenerAlreadyExists
	}

	sr.db[urlsh.ID] = urlsh

	return nil
}

func (sr *shortenerRepository) Get(id string) (shortener.Model, error) {

	model := shortener.Model{}
	if sr.db == nil {
		return model, repo.ErrInvalidRepository
	}

	if len(id) == 0 {
		return model, repo.ErrInvalidID
	}

	sr.mtx.Lock()
	defer sr.mtx.Unlock()
	model, ok := sr.db[id]
	if !ok {
		return model, repo.ErrShortener404
	}

	return model, nil
}

func (sr *shortenerRepository) Update(id string, m shortener.Model) error {

	if sr.db == nil {
		return repo.ErrInvalidRepository
	}

	if len(id) == 0 {
		return repo.ErrInvalidID
	}

	if _, err := sr.Get(id); err != nil {
		return repo.ErrShortener404
	}

	sr.mtx.Lock()
	defer sr.mtx.Unlock()

	sr.db[id] = m

	return nil
}

func (sr *shortenerRepository) Remove(id string) error {

	if sr.db == nil {
		return repo.ErrInvalidRepository
	}

	if len(id) == 0 {
		return repo.ErrInvalidID
	}

	if _, err := sr.Get(id); err != nil {
		return repo.ErrShortener404
	}

	delete(sr.db, id)

	return nil
}
