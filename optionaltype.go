package main

import (
	"encoding/json"
)

type OptionalType[T any] struct {
	Value   T
	Defined bool
	Valid   bool
}

func NewOptionalType[T any](v *T) OptionalType[*T] {
	if v == nil {
		return OptionalType[*T]{nil, true, false}
	}
	return OptionalType[*T]{v, true, true}
}

func (o *OptionalType[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		var zero T
		o.Value = zero
		o.Defined = true
		o.Valid = false
		return nil
	}

	var value T
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	o.Value = value
	o.Defined = true
	o.Valid = true
	return nil
}

func (o OptionalType[T]) MarshalJSON() ([]byte, error) {
	if !o.Defined {
		return []byte("null"), nil
	}
	return json.Marshal(o.Value)
}

func (o OptionalType[T]) ValidateValue() T {
	return o.Value
}
