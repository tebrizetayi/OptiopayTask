package api

import (
	"net/http"
	"testing"
	"github.com/tebrizetayi/optiopay/internal/storage"
)

type enviroment struct {
	controller Controller
	api        http.Handler
	storage    *storage.InMemoryStorage
}

var env enviroment

func TestMain(m *testing.M) {
}
