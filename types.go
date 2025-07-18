package main

import "github.com/google/uuid"

type (
	OptionalInt   = OptionalType[*int]
	OptionalInt8  = OptionalType[*int8]
	OptionalInt16 = OptionalType[*int16]
	OptionalInt32 = OptionalType[*int32]
	OptionalInt64 = OptionalType[*int64]

	OptionalUint   = OptionalType[*uint]
	OptionalUint8  = OptionalType[*uint8]
	OptionalUint16 = OptionalType[*uint16]
	OptionalUint32 = OptionalType[*uint32]
	OptionalUint64 = OptionalType[*uint64]

	OptionalFloat32 = OptionalType[*float32]
	OptionalFloat64 = OptionalType[*float64]

	OptionalString = OptionalType[*string]

	OptionalBool = OptionalType[*bool]

	OptionalUUID = OptionalType[*uuid.UUID]
)

func NewOptionalInt(v *int) OptionalInt       { return NewOptionalType(v) }
func NewOptionalInt8(v *int8) OptionalInt8    { return NewOptionalType(v) }
func NewOptionalInt16(v *int16) OptionalInt16 { return NewOptionalType(v) }
func NewOptionalInt32(v *int32) OptionalInt32 { return NewOptionalType(v) }
func NewOptionalInt64(v *int64) OptionalInt64 { return NewOptionalType(v) }

func NewOptionalUint(v *uint) OptionalUint       { return NewOptionalType(v) }
func NewOptionalUint8(v *uint8) OptionalUint8    { return NewOptionalType(v) }
func NewOptionalUint16(v *uint16) OptionalUint16 { return NewOptionalType(v) }
func NewOptionalUint32(v *uint32) OptionalUint32 { return NewOptionalType(v) }
func NewOptionalUint64(v *uint64) OptionalUint64 { return NewOptionalType(v) }

func NewOptionalFloat32(v *float32) OptionalFloat32 { return NewOptionalType(v) }
func NewOptionalFloat64(v *float64) OptionalFloat64 { return NewOptionalType(v) }

func NewOptionalString(v *string) OptionalString { return NewOptionalType(v) }

func NewOptionalBool(v *bool) OptionalBool { return NewOptionalType(v) }

func NewOptionalUUID(v *uuid.UUID) OptionalUUID { return NewOptionalType(v) }
