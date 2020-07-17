package shortener

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator"
	"github.com/thanhpk/randstr"
)

type Service interface {
	// NewShortener creates a new shortener by a url string
	NewShortener(url string) (Model, error)
	// GetShortener returns the url string by ID. or an error.
	GetShortener(id string) (string, error)
	// RemoveShortener removes the shortenered object containing the ID and the URL.
	RemoveShortener(id string) error
	// UpdateShortener discarded for security reasons
	//UpdateShortener(id, url string) (error)
}

type service struct {
	repo   Repository
	config Config
}

var (
	ErrInvalidID = errors.New("Invalid ID Length")
)

// NewService returns a shortener service based on repository and idlength.
func NewService(srepo Repository, config Config) (Service, error) {
	s := &service{}

	if srepo == nil {
		return nil, errors.New("respository is invalid")
	}

	if validate == nil {
		validate = validator.New()
	}

	if err := validate.Struct(config); err != nil {
		return nil, fmt.Errorf("config.Validate: %w", err)
	}

	s.repo = srepo
	s.config = config

	return s, nil
}

func (s *service) NewShortener(url string) (m Model, err error) {
	id := randstr.String(s.config.Length, s.config.Characters)

	model := Model{
		ID:        id,
		URLString: url,
	}

	if err = model.Validate(); err != nil {
		return
	}

	if err = s.repo.Create(model); err != nil {
		return
	}

	m = model
	return
}

func (s *service) GetShortener(id string) (string, error) {

	if len(id) != s.config.Length {
		return "", ErrInvalidID
	}

	m, err := s.repo.Get(id)
	if err != nil {
		return "", fmt.Errorf("repo.Get: %w", err)
	}

	return m.URLString, nil

}

func (s *service) RemoveShortener(id string) error {

	if len(id) != s.config.Length {
		return ErrInvalidID
	}

	err := s.repo.Remove(id)
	if err != nil {
		return fmt.Errorf("repo.Remove: %w", err)
	}

	return nil
}
