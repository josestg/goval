package goval

import (
	"context"
	"fmt"
	"github.com/pkg-id/goval/funcs"
	"strings"
)

// StringValidator is a FunctionValidator that validates string.
type StringValidator FunctionValidator[string]

// String returns a StringValidator with no rules.
func String() StringValidator {
	return NopFunctionValidator[string]
}

// Validate executes the validation rules immediately.
func (f StringValidator) Validate(ctx context.Context, value string) error {
	return validatorOf(f, value).Validate(ctx)
}

// With attaches the next rule to the chain.
func (f StringValidator) With(next StringValidator) StringValidator {
	return Chain(f, next)
}

// Required ensures the string is not empty.
func (f StringValidator) Required() StringValidator {
	return f.With(func(ctx context.Context, value string) error {
		if value == "" {
			return NewRuleError(StringRequired)
		}
		return nil
	})
}

// Min ensures the length of the string is not less than the given length.
func (f StringValidator) Min(length int) StringValidator {
	return f.With(func(ctx context.Context, value string) error {
		if len(value) < length {
			return NewRuleError(StringMin, length)
		}
		return nil
	})
}

// Max ensures the length of the string is not greater than the given length.
func (f StringValidator) Max(length int) StringValidator {
	return f.With(func(ctx context.Context, value string) error {
		if len(value) > length {
			return NewRuleError(StringMax, length)
		}
		return nil
	})
}

// Match ensures the string matches the given pattern.
// If pattern cause panic, will be recovered.
func (f StringValidator) Match(pattern Pattern) StringValidator {
	return f.With(func(ctx context.Context, value string) (err error) {
		defer func() {
			if rec := recover(); rec != nil {
				err = fmt.Errorf("panic: %v", rec)
			}
		}()

		exp := pattern.RegExp()
		if !exp.MatchString(value) {
			return NewRuleError(StringMatch, exp.String())
		}

		return err
	})
}

// In ensures that the provided string is one of the specified options.
// This validation is case-sensitive, use InFold to perform a case-insensitive In validation.
func (f StringValidator) In(options ...string) StringValidator {
	return f.With(func(ctx context.Context, value string) error {
		ok := funcs.Contains(options, func(opt string) bool { return value == opt })
		if !ok {
			return NewRuleError(StringIn, options)
		}
		return nil
	})
}

// InFold ensures that the provided string is one of the specified options with case-insensitivity.
func (f StringValidator) InFold(options ...string) StringValidator {
	return f.With(func(ctx context.Context, value string) error {
		ok := funcs.Contains(options, func(opt string) bool { return strings.EqualFold(value, opt) })
		if !ok {
			return NewRuleError(StringInFold, options)
		}
		return nil
	})
}

// When adds validation logic to the chain based on a condition for string values.
//
// If the predicate returns true, the result of the mapper function is added to the chain,
// and the input value is validated using the new chain. Otherwise, the original chain is returned unmodified.
//
// The mapper function takes a StringValidator instance and returns a new StringValidator instance with
// additional validation logic.
func (f StringValidator) When(p Predicate[string], m Mapper[string, StringValidator]) StringValidator {
	return whenLinker(f, p, m)
}
