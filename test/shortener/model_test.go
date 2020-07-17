package shortenertest

import (
	"testing"

	"github.com/toms1441/urlsh/internal/shortener"
)

var invalidmodel = [4]shortener.Model{
	// first one empty
	{},
	// should be invalid cause id is empty
	{
		URLString: "awd",
	},
	// should be invalid cause url is empty
	{
		ID: "12341234",
	},
	// should be invalid cause url is invalid
	{
		ID:        "12341234",
		URLString: "www.example.com",
	},
}

var validmodel = shortener.Model{
	ID:        "12341234",
	URLString: "http://www.example.com",
}

func TestModelValidation(t *testing.T) {
	for k, v := range invalidmodel {
		err := v.Validate()
		if err == nil {
			t.Fatalf("v.Validate == nil - %d", k)
		}
	}

	err := validmodel.Validate()
	if err != nil {
		t.Fatalf("validmodel.Validate: %v", err)
	}
}
