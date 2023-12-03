package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	nombre := "Carla"
	want := regexp.MustCompile(`\b` + nombre + `\b`)
	msg, err := Hello("Carla")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Carla") = %q, %v, quiere coincidencia para %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, quiere "", error`, msg, err)
	}
}