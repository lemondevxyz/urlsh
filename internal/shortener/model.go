package shortener

import (
	"fmt"
	"net/url"

	"github.com/go-playground/validator"
)

// Model is a struct defining fields for the shortener package.
// It needs an ID and URL String to be valid.
type Model struct {
	ID        string `validate:"required" json:"id"`
	URLString string `validate:"required" json:"url"`
}

var validate *validator.Validate

func (m Model) Validate() error {
	if validate == nil {
		validate = validator.New()
	}

	err := validate.Struct(m)
	if err != nil {
		return fmt.Errorf("validate.Struct: %v", err)
	}

	curl, err := url.Parse(m.URLString)
	if err != nil {
		return fmt.Errorf("url.Parse: %v", err)
	}

	if len(curl.Scheme) == 0 || len(curl.Host) == 0 {
		return fmt.Errorf("invalid url")
	}

	return nil
}
