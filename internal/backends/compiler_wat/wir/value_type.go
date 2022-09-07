package wir

import (
	"github.com/wa-lang/wa/internal/backends/compiler_wat/wir/wat"
	"github.com/wa-lang/wa/internal/logger"
)

/**************************************
Void:
**************************************/
type Void struct {
}

func (t Void) byteSize() int        { return 0 }
func (t Void) Raw() []wat.ValueType { return []wat.ValueType{} }
func (t Void) Equal(u ValueType) bool {
	if _, ok := u.(Void); ok {
		return true
	}
	return false
}

/**************************************
Int32:
**************************************/
type Int32 struct {
}

func (t Int32) byteSize() int        { return 4 }
func (t Int32) Raw() []wat.ValueType { return []wat.ValueType{wat.I32{}} }
func (t Int32) Equal(u ValueType) bool {
	if _, ok := u.(Int32); ok {
		return true
	}
	return false
}

/**************************************
Int64:
**************************************/
type Int64 struct {
}

func (t Int64) byteSize() int        { return 8 }
func (t Int64) Raw() []wat.ValueType { return []wat.ValueType{wat.I64{}} }
func (t Int64) Equal(u ValueType) bool {
	if _, ok := u.(Int64); ok {
		return true
	}
	return false
}

/**************************************
Pointer:
**************************************/
type Pointer struct {
	Base ValueType
}

func NewPointer(base ValueType) Pointer { return Pointer{Base: base} }
func (t Pointer) byteSize() int         { return 4 }
func (t Pointer) Raw() []wat.ValueType  { return []wat.ValueType{wat.I32{}} }
func (t Pointer) Equal(u ValueType) bool {
	if ut, ok := u.(Pointer); ok {
		return t.Base.Equal(ut.Base)
	}
	return false
}

/**************************************
Struct:
**************************************/
type Struct struct {
	name    string
	Members []Field
}

type Field struct {
	name string
	typ  ValueType
}

func NewField(n string, t ValueType) *Field { return &Field{name: n, typ: t} }
func (i Field) Name() string                { return i.name }
func (i Field) Type() ValueType             { return i.typ }
func (i Field) Equal(u Field) bool          { return i.name == u.name && i.typ.Equal(u.typ) }

func NewStruct(name string, m []Field) *Struct {
	return &Struct{name: name, Members: m}
}

func (t Struct) byteSize() int { logger.Fatal("Todo"); return 0 }
func (t Struct) Raw() []wat.ValueType {
	var r []wat.ValueType
	for _, f := range t.Members {
		r = append(r, f.Type().Raw()...)
	}
	return r
}
func (t Struct) Equal(u ValueType) bool {
	if u, ok := u.(Struct); ok {
		if len(t.Members) != len(u.Members) {
			return false
		}

		for i := range t.Members {
			if !t.Members[i].Equal(u.Members[i]) {
				return false
			}
		}

		return true
	}
	return false
}

/**************************************
Ref:
**************************************/
type Ref struct {
	Base ValueType
}

func NewRef(base ValueType) Ref    { return Ref{Base: base} }
func (t Ref) byteSize() int        { return 8 }
func (t Ref) Raw() []wat.ValueType { return []wat.ValueType{wat.I32{}, wat.I32{}} }
func (t Ref) Equal(u ValueType) bool {
	if ut, ok := u.(Ref); ok {
		return t.Base.Equal(ut.Base)
	}
	return false
}