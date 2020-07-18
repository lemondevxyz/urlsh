package json

import (
	"encoding/json"
	"fmt"

	nanobox "github.com/nanobox-io/golang-scribble"
	"github.com/toms1441/urlsh/internal/repo"
	"github.com/toms1441/urlsh/internal/shortener"
)

type shortenerRepository struct {
	db *nanobox.Driver
}

var shortenersCollection = "shorteners"

func NewShortenerRepository(filepath string) (shortener.Repository, error) {
	db, err := nanobox.New(filepath, nil)
	if err != nil {
		return nil, fmt.Errorf("nanobox.New: %v", err)
	}

	sr := &shortenerRepository{
		db: db,
	}

	return sr, nil
}

func (sr *shortenerRepository) Create(urlsh shortener.Model) error {

	if sr.db == nil {
		return repo.ErrInvalidRepository
	}

	if len(urlsh.ID) == 0 {
		return repo.ErrInvalidID
	}

	if _, err := sr.Get(urlsh.ID); err == nil {
		return repo.ErrShortenerAlreadyExists
	}

	err := sr.db.Write(shortenersCollection, urlsh.ID, urlsh)
	if err != nil {
		return fmt.Errorf("sr.db.Write: %v", err)
	}

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

	err := sr.db.Delete(shortenersCollection, id)
	if err != nil {
		return fmt.Errorf("sr.db.Delete: %v", err)
	}

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

	if err := sr.db.Read(shortenersCollection, id, &model); err != nil {
		return model, fmt.Errorf("sr.db.Read: %w", err)
	}

	return model, nil

}

func (sr *shortenerRepository) GetAll() ([]shortener.Model, error) {

	models := []shortener.Model{}
	if sr.db == nil {
		return models, repo.ErrInvalidRepository
	}

	ids, err := sr.db.ReadAll(shortenersCollection)
	if err != nil {
		return models, fmt.Errorf("sr.db.ReadAll: %w", err)
	}

	for _, v := range ids {

		newmodel := shortener.Model{}
		err = json.Unmarshal([]byte(v), &newmodel)

		if err == nil {
			models = append(models, newmodel)
		}
	}

	return models, nil
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

	err := sr.db.Write(shortenersCollection, id, m)
	if err != nil {
		return fmt.Errorf("sr.db.Write: %w", err)
	}

	return nil
}
