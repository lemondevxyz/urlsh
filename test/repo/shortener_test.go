package repotest

import (
	"testing"

	"github.com/toms1441/urlsh/internal/repo"
	"github.com/toms1441/urlsh/internal/repo/json"
	"github.com/toms1441/urlsh/internal/repo/plain"
	"github.com/toms1441/urlsh/internal/shortener"
)

const method = "json"

var sr shortener.Repository
var samples = [4]shortener.Model{
	{
		ID:        "ok",
		URLString: "example.com",
	},
	{
		ID:        "megababy",
		URLString: "github.com/toms1441/urlsh",
	},
	{
		ID:        "12341234",
		URLString: "",
	},
}

func TestNewRepository(t *testing.T) {
	var err error

	switch method {
	case "plain":
		sr, err = plain.NewShortenerRepository()
		if err != nil {
			t.Fatalf("plain.NewShortenerRepository: %v", err)
		}
	case "json":
		sr, err = json.NewShortenerRepository("./db")
		if err != nil {
			t.Fatalf("json.NewShortenerRepository: %v", err)
		}
	}

	t.Logf("Testing with method: %v", method)

}

func TestRepositoryCreate(t *testing.T) {
	for k, v := range samples {
		err := sr.Create(v)
		if k != 3 {
			if err != nil {
				t.Fatalf("sr.Create: %d %v", k, err)
			}
		} else {
			if err != repo.ErrInvalidID {
				t.Fatalf("sr.Create != repo.ErrInvalidID")
			}
		}
	}
}

func TestRepositoryGet(t *testing.T) {
	for k, v := range samples {
		model, err := sr.Get(v.ID)

		if k != 3 {
			if err != nil {
				t.Fatalf("sr.Get: %d %v", k, err)
			}

			if model != v {
				t.Fatalf("sr.Get != sample - %d", k)
			}
		} else {
			if err != repo.ErrInvalidID {
				t.Fatalf("sr.Get != repo.ErrInvalidID")
			}
		}
	}
}

func TestRepositoryUpdate(t *testing.T) {
	for k, v := range samples {
		newv := v
		newv.URLString = "update"

		err := sr.Update(v.ID, newv)
		if k != 3 {
			if err != nil {
				t.Fatalf("sr.Update: %d %v", k, err)
			}

			model, err := sr.Get(v.ID)
			if err != nil {
				t.Fatalf("sr.Get: %d %v", k, err)
			}

			if model != newv {
				t.Fatalf("sr.Update != newsample - %d", k)
			}
		} else {
			if err != repo.ErrInvalidID {
				t.Fatalf("sr.Get != repo.ErrInvalidID")
			}
		}
	}
}

func TestRepositoryRemove(t *testing.T) {
	for k, v := range samples {
		err := sr.Remove(v.ID)
		if k != 3 {
			if err != nil {
				t.Fatalf("sr.Remove: %d %v", k, err)
			}

			_, err = sr.Get(v.ID)
			if err == nil {
				t.Fatalf("sr.Get == nil - %d", k)
			}
		} else {
			if err != repo.ErrInvalidID {
				t.Fatalf("sr.Get != repo.ErrInvalidID")
			}

		}
	}
}
