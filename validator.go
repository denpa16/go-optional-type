package main

import (
	"github.com/go-playground/validator/v10"
	"reflect"
)

type OptionalTypes []any

func NewValidator() *validator.Validate {
	validate := validator.New()
	return validate
}

func Validate(requestData interface{}, optionalTypes OptionalTypes) error {
	validate := NewValidator()
	RegisterDefaultOptionalTypesValidators(validate)

	if len(optionalTypes) > 0 {
		validate.RegisterCustomTypeFunc(
			func(field reflect.Value) any {
				getValueMethod := field.MethodByName("ValidateValue")
				if getValueMethod.IsValid() {
					results := getValueMethod.Call(nil)
					if len(results) > 0 {
						return results[0].Interface()
					}
				}
				return nil
			},
			optionalTypes...,
		)
	}

	err := validate.Struct(requestData)
	if err != nil {
		return err
	}
	return nil
}

func RegisterDefaultOptionalTypesValidators(validate *validator.Validate) {
	optionalTypes := []any{
		OptionalInt{},
		OptionalInt8{},
		OptionalInt16{},
		OptionalInt32{},
		OptionalInt64{},
		OptionalUint{},
		OptionalUint8{},
		OptionalUint16{},
		OptionalUint32{},
		OptionalUint64{},
		OptionalFloat32{},
		OptionalFloat64{},
		OptionalString{},
		OptionalBool{},
		OptionalUUID{},
	}
	validate.RegisterCustomTypeFunc(
		func(field reflect.Value) any {
			getValueMethod := field.MethodByName("ValidateValue")
			if getValueMethod.IsValid() {
				results := getValueMethod.Call(nil)
				if len(results) > 0 {
					return results[0].Interface()
				}
			}
			return nil
		},
		optionalTypes...,
	)
}
