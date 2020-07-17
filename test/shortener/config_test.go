package shortenertest

import (
	"testing"

	"github.com/go-playground/validator"
	"github.com/toms1441/urlsh/internal/shortener"
)

var invalidconfig = [3]shortener.Config{
	{},
	{Length: 1},
	{Characters: "1"},
}

var validconfig = shortener.Config{
	Length:     4,
	Characters: "abcdef",
}

var validate *validator.Validate

func TestConfigValidation(t *testing.T) {
	if validate == nil {
		validate = validator.New()
	}

	for k, v := range invalidconfig {
		err := validate.Struct(v)
		if err == nil {
			t.Fatalf("validate.Struct == nil - %d", k)
		}
	}

	err := validate.Struct(validconfig)
	if err != nil {
		t.Fatalf("validate.Struct: %v", err)
	}
}
