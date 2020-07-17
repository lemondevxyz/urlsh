package repo

import "errors"

var (
	ErrShortenerAlreadyExists = errors.New("shortener already exists")
	ErrShortener404           = errors.New("shortener does not exist")
	ErrInvalidRepository      = errors.New("invalid repository")
	ErrInvalidID              = errors.New("invalid id length")
)
