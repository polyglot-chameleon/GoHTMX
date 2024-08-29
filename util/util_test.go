package util

import (
	"os"
	"testing"
)

func TestLoadDotEnv(t *testing.T) {
	LoadDotEnv("../.env.exmpl")

	if os.Getenv("DB_DRIVER") != "sqlite3" {
		t.Fatalf("DB_DRIVER != %s", os.Getenv("DB_DRIVER"))
	}

	if os.Getenv("DB_URL") != "db/gohtmx.db" {
		t.Fatalf("DB_URL != %s", os.Getenv("DB_URL"))
	}

	if os.Getenv("PORT") != ":8080" {
		t.Fatalf("PORT != %s", os.Getenv(":8080"))
	}
}
