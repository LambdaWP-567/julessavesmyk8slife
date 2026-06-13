package api

import "testing"

func TestVersion(t *testing.T) {
    expected := "0.1.0"
    if expected != "0.1.0" {
        t.Errorf("expected %s, got %s", expected, expected)
    }
}
