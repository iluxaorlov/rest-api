package store_test

import (
	"os"
	"testing"
)

var url string

func TestMain(m *testing.M) {
	url = os.Getenv("DATABASE_URL")
	if url == "" {
		url = "host=localhost dbname=db user=user password=password sslmode=disable"
	}

	os.Exit(m.Run())
}