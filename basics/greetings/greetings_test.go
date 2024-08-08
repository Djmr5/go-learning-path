package greetings

import (
	"regexp"
	"testing"
)

// Test call of greetings.Hello with a name
// check for a valid return type.
func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello(name)

	// Check both the message is in the correct format
	// and there is no error returned
	if !want.MatchString(msg) || err != nil {
		t.Fatalf("Hello(\"Gladys\") = %q, %v, want match for %#q, nil", msg, err, want)
	}

}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
