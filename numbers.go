package goval

import (
	"context"
	"github.com/pkg-id/goval/constraints"
	"github.com/pkg-id/goval/funcs"
)

// NumberConstraint is set types that treats as numbers.
type NumberConstraint interface {
	constraints.Integer | constraints.Float
}

// NumberValidator is a validator that validates numbers.
type NumberValidator[T NumberConstraint] FunctionValidator[T]

// Number returns a NumberValidator with no rules.
func Number[T NumberConstraint]() NumberValidator[T] {
	return NopFunctionValidator[T]()
}

// Build builds the validator chain and attaches the value to it.
func (nv NumberValidator[T]) Build(value T) Validator {
	return validatorOf(nv, value)
}

// With attaches the next rule to the chain.
func (nv NumberValidator[T]) With(next NumberValidator[T]) NumberValidator[T] {
	return Chain(nv, next)
}

// Required ensures the number is not zero.
func (nv NumberValidator[T]) Required() NumberValidator[T] {
	return nv.With(func(ctx context.Context, value T) error {
		var zero T
		if value == zero {
			return NewRuleError(NumberRequired)
		}
		return nil
	})
}

// Min ensures the number is not less than the given min.
func (nv NumberValidator[T]) Min(min T) NumberValidator[T] {
	return nv.With(func(ctx context.Context, value T) error {
		if value < min {
			return NewRuleError(NumberMin, min)
		}
		return nil
	})
}

// Max ensures the number is not greater than the given max.
func (nv NumberValidator[T]) Max(max T) NumberValidator[T] {
	return nv.With(func(ctx context.Context, value T) error {
		if value > max {
			return NewRuleError(NumberMax, max)
		}
		return nil
	})
}

// In ensures that the provided number is one of the specified options.
func (nv NumberValidator[T]) In(options ...T) NumberValidator[T] {
	return nv.With(func(ctx context.Context, value T) error {
		ok := funcs.Contains(options, func(opt T) bool { return opt == value })
		if !ok {
			return NewRuleError(NumberIn, options)
		}
		return nil
	})
}
