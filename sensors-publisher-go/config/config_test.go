package config

import (
	"os"
	"testing"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

func TestGetBool(t *testing.T) {
	os.Setenv("TEST_BOOL", "true")

	expect := true

	if got := Get[bool]("TEST_BOOL"); got != expect {
		t.FailNow()
	}
}

func TestGetDuration(t *testing.T) {
	os.Setenv("TEST_DURATION", "6m")

	expect := time.Minute * 6

	if got := Get[time.Duration]("TEST_DURATION"); got != expect {
		t.FailNow()
	}
}

func TestGetFloat64(t *testing.T) {
	os.Setenv("TEST_FLOAT64", "78.51")

	expect := 78.51

	if got := Get[float64]("TEST_FLOAT64"); got != expect {
		t.FailNow()
	}
}

func TestGetInt(t *testing.T) {
	os.Setenv("TEST_INT", "9798")

	expect := 9798

	if got := Get[int]("TEST_INT"); got != expect {
		t.FailNow()
	}
}

func TestGetInt64(t *testing.T) {
	os.Setenv("TEST_INT64", "15478965")

	expect := int64(15478965)

	if got := Get[int64]("TEST_INT64"); got != expect {
		t.FailNow()
	}
}

func TestGetString(t *testing.T) {
	os.Setenv("TEST_STRING", "foo")

	expect := "foo"

	if got := Get[string]("TEST_STRING"); got != expect {
		t.FailNow()
	}
}

func TestGetUnknownType(t *testing.T) {
	os.Setenv("TEST_UNKNOWN_TYPE", "foo")

	var expect int32

	if got := Get[int32]("TEST_UNKNOWN_TYPE"); got != expect {
		t.FailNow()
	}
}

func TestGetInvalidKey(t *testing.T) {
	expect := int(0)

	if got := Get[int]("TEST_INVALID_KEY"); got != expect {
		t.FailNow()
	}
}
