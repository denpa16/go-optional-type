package main

import (
	"encoding/json"
	"fmt"
)

type FieldEnum string

const FieldEnum1 FieldEnum = "first"
const FieldEnum2 FieldEnum = "second"

type Example struct {
	ID           OptionalUUID            `json:"id"`
	EnumField    OptionalType[FieldEnum] `json:"enum_field" validate:"required,oneof=first second"`
	IntField     OptionalInt             `json:"int_field" validate:"required,min=50"`
	Int64Field   OptionalInt64           `json:"int64_field"`
	Float32Field OptionalFloat32         `json:"float32_field"`
	Float64Field OptionalFloat64         `json:"float64_field"`
	StringField  OptionalString          `json:"string_field"`
	BoolField    OptionalBool            `json:"bool_field"`
}

func main() {
	jsonData := []byte(
		`{
			"field":"first",
			"int_field":10,
			"int64_field":20,
			"string_field":"hello",
			"bool_field":true,
			"enum_field":"first",
			"float64_field":3.14159265359
		}`,
	)
	example := Example{}
	err := json.Unmarshal(jsonData, &example)
	if err != nil {
		panic(err)
	}

	err = Validate(example, OptionalTypes{OptionalType[FieldEnum]{}})
	if err != nil {
		panic(err)
	}

	fmt.Println("ID Value:", example.ID.Value)
	fmt.Println("ID Defined:", example.ID.Defined)
	fmt.Println("ID Valid:", example.ID.Valid)

	fmt.Println("EnumField Value:", example.EnumField.Value)
	fmt.Println("EnumField Defined:", example.EnumField.Defined)
	fmt.Println("EnumField Valid:", example.EnumField.Valid)

	fmt.Println("IntField Value:", example.IntField.Value)
	fmt.Println("IntField Defined:", example.IntField.Defined)
	fmt.Println("IntField Valid:", example.IntField.Valid)

	fmt.Println("Int64Field Value:", example.Int64Field.Value)
	fmt.Println("Int64Field Defined:", example.Int64Field.Defined)
	fmt.Println("Int64Field Valid:", example.Int64Field.Valid)

	fmt.Println("Float32Field Value:", example.Float32Field.Value)
	fmt.Println("Float32Field Defined:", example.Float32Field.Defined)
	fmt.Println("Float32Field Valid:", example.Float32Field.Valid)

	fmt.Println("Float64Field Value:", example.Float64Field.Value)
	fmt.Println("Float64Field Defined:", example.Float64Field.Defined)
	fmt.Println("Float64Field Valid:", example.Float64Field.Valid)

	fmt.Println("StringField Value:", example.StringField.Value)
	fmt.Println("StringField Defined:", example.StringField.Defined)
	fmt.Println("StringField Valid:", example.StringField.Valid)

	fmt.Println("BoolField Value:", example.BoolField.Value)
	fmt.Println("BoolField Defined:", example.BoolField.Defined)
	fmt.Println("BoolField Valid:", example.BoolField.Valid)

}
