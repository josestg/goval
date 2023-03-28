package goval_test

import (
	"github.com/pkg-id/goval"
	"testing"
)

func TestErrors(t *testing.T) {
	t.Run("Error: using var", func(t *testing.T) {
		var errs goval.Errors
		str := errs.Error()
		if str != "null" {
			t.Errorf("expect string Error: %q; got %q", "null", str)
		}
	})

	t.Run("Error: using make", func(t *testing.T) {
		errs := make(goval.Errors, 0)
		str := errs.Error()
		if str != "[]" {
			t.Errorf("expect string Error: %q; got %q", "[]", str)
		}
	})
}
