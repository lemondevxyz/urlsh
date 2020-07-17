package shortenertest

import (
	"testing"

	"github.com/toms1441/urlsh/internal/repo/plain"
	"github.com/toms1441/urlsh/internal/shortener"
)

var ss shortener.Service
var sr shortener.Repository
var modelid, modelurl string

func TestNewService(t *testing.T) {
	var err error
	ss, err = shortener.NewService(nil, shortener.Config{})
	if err == nil {
		t.Fatalf("shortener.NewService == nil - should return 'repository is invalid'")
	}

	ss, err = shortener.NewService(sr, shortener.Config{})
	if err == nil {
		t.Fatalf("shortener.NewService == nil - should return 'config is invalid'")
	}

	sr, err = plain.NewShortenerRepository()
	if err != nil {
		t.Fatalf("plain.NewShortenerRepository: %v", err)
	}

	ss, err = shortener.NewService(sr, validconfig)
	if err != nil {
		t.Fatalf("shortener.NewService: %v", err)
	}
}

func TestServiceNewShortener(t *testing.T) {
	model, err := ss.NewShortener("")
	if err == nil {
		t.Fatalf("ss.NewShortener == nil - should return validate.Struct error")
	}

	model, err = ss.NewShortener("invalid")
	if err == nil {
		t.Fatalf("ss.NewShortener == nil - should return invalid url")
	}

	model, err = ss.NewShortener("https://www.google.com")
	if err != nil {
		t.Fatalf("ss.NewShortener: %v", err)
	}

	modelid = model.ID
	modelurl = model.URLString
}

func TestServiceGetShortener(t *testing.T) {
	tempurl, err := ss.GetShortener(modelid)
	if err != nil {
		t.Fatalf("ss.GetShortener: %v", err)
	}

	if tempurl != modelurl {
		t.Fatal("tempurl != modelurl")
	}
}

func TestServiceRemoveShortener(t *testing.T) {
	err := ss.RemoveShortener(modelid)
	if err != nil {
		t.Fatalf("ss.RemoveShortener: %v", err)
	}

	_, err = ss.GetShortener(modelid)
	if err == nil {
		t.Fatalf("ss.GetShortener == nil, should return err")
	}
}
