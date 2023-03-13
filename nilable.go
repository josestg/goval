package goval

import "context"

type NilValidator[T any] FunctionValidator[*T]

func Nil[T any]() NilValidator[T] { return NopValueValidator[*T] }

func (f NilValidator[T]) Build(value *T) Validator {
	return validatorOf(f, value)
}

func (f NilValidator[T]) Required() NilValidator[T] {
	return Chain(f, func(ctx context.Context, value *T) error {
		if value == nil {
			return NewRuleError(NilRequired, value)
		}
		return nil
	})
}

func (f NilValidator[T]) Optional(builder Builder[T]) NilValidator[T] {
	return Chain(f, func(ctx context.Context, value *T) error {
		if value != nil {
			return builder.Build(*value).Validate(ctx)
		}
		return nil
	})
}

func (f NilValidator[T]) Next(builder Builder[T]) NilValidator[T] {
	return Chain(f, func(ctx context.Context, value *T) error {
		return builder.Build(*value).Validate(ctx)
	})
}
