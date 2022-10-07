package mapper

import "testing"

// Run these by running `go test` in the directory containing this file.
func TestMapString(t *testing.T) {
	s := NewSkipString(3, "Aspiration.com")
	MapString(&s)
	if s.Value != "asPirAtiOn.cOm" {
		t.Errorf("Expected 'asPirAtiOn.cOm', got '%s'", s.Value)
	}
}

func TestMapStringAllCaps(t *testing.T) {
	s := NewSkipString(1, "Aspiration.com")
	MapString(&s)
	if s.Value != "ASPIRATION.COM" {
		t.Errorf("Expected 'ASPIRATION.COM', got '%s'", s.Value)
	}
}

func TestMapStringZero(t *testing.T) {
	// Sometimes if I run into a bug, I like to create a unit test to reproduce it going forward.
	// This helps avoid regressions and also provides a level of documentation.
	s := NewSkipString(0, "Aspiration.com")
	MapString(&s)
	if s.Value != "Aspiration.com" {
		t.Errorf("Expected 'Aspiration.com', got '%s'", s.Value)
	}
}

func TestMapStringDoubleMap(t *testing.T) {
	// Ideally you would programatically avoid this case, but it's good to test it anyway.
	s := NewSkipString(3, "Aspiration.com")
	MapString(&s)
	MapString(&s)
	if s.Value != "asPirAtiOn.cOm" {
		t.Errorf("Expected 'asPirAtiOn.cOm', got '%s'", s.Value)
	}
}

func TestMapStringLeadingNonAlphanumeric(t *testing.T) {
	s := NewSkipString(3, ".Aspiration.com")
	MapString(&s)
	if s.Value != ".asPirAtiOn.cOm" {
		t.Errorf("Expected '.asPirAtiOn.cOm', got '%s'", s.Value)
	}
}

func TestMapStringManyNonAlphanumeric(t *testing.T) {
	s := NewSkipString(3, "....abc.....def....ghi....")
	MapString(&s)
	if s.Value != "....abC.....deF....ghI...." {
		t.Errorf("Expected '....abC.....deF....ghI....', got '%s'", s.Value)
	}
}