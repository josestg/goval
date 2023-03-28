package goval_test

import (
	"context"
	"errors"
	"github.com/pkg-id/goval"
	"testing"
)

func TestNamed(t *testing.T) {
	t.Run("when validation fails", func(t *testing.T) {
		ctx := context.Background()
		err := goval.Named("field-name", "", goval.String().Required()).Validate(ctx)

		var exp *goval.KeyError
		if !errors.As(err, &exp) {
			t.Fatalf("expect error type: %T; got error type: %T", exp, err)
		}

		if exp.Key != "field-name" {
			t.Errorf("expect error key is field-name; got %s", exp.Key)
		}

		if exp.Err == nil {
			t.Errorf("expect error is not nil")
		}
	})

	t.Run("when validation ok", func(t *testing.T) {
		ctx := context.Background()
		err := goval.Named("field-name", "a", goval.String().Required()).Validate(ctx)
		if err != nil {
			t.Fatalf("expect not error")
		}
	})
}
