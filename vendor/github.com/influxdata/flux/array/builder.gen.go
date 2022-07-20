// Generated by tmpl
// https://github.com/benbjohnson/tmpl
//
// DO NOT EDIT!
// Source: builder.gen.go.tmpl

package array

import (
	"github.com/apache/arrow/go/v7/arrow/array"
	"github.com/apache/arrow/go/v7/arrow/memory"
)

type (
	Int     = array.Int64
	Uint    = array.Uint64
	Float   = array.Float64
	Boolean = array.Boolean
)

type IntBuilder struct {
	b *array.Int64Builder
}

func NewIntBuilder(mem memory.Allocator) *IntBuilder {
	return &IntBuilder{
		b: array.NewInt64Builder(mem),
	}
}
func (b *IntBuilder) Retain() {
	b.b.Retain()
}
func (b *IntBuilder) Release() {
	b.b.Release()
}
func (b *IntBuilder) Len() int {
	return b.b.Len()
}
func (b *IntBuilder) Cap() int {
	return b.b.Cap()
}
func (b *IntBuilder) Append(v int64) {
	b.b.Append(v)
}
func (b *IntBuilder) AppendValues(v []int64, valid []bool) {
	b.b.AppendValues(v, valid)
}
func (b *IntBuilder) UnsafeAppend(v int64) {
	b.b.UnsafeAppend(v)
}
func (b *IntBuilder) NullN() int {
	return b.b.NullN()
}
func (b *IntBuilder) AppendNull() {
	b.b.AppendNull()
}
func (b *IntBuilder) UnsafeAppendBoolToBitmap(isValid bool) {
	b.b.UnsafeAppendBoolToBitmap(isValid)
}
func (b *IntBuilder) Reserve(n int) {
	b.b.Reserve(n)
}
func (b *IntBuilder) Resize(n int) {
	b.b.Resize(n)
}
func (b *IntBuilder) NewArray() Array {
	return b.NewIntArray()
}
func (b *IntBuilder) NewIntArray() *Int {
	return b.b.NewInt64Array()
}

type UintBuilder struct {
	b *array.Uint64Builder
}

func NewUintBuilder(mem memory.Allocator) *UintBuilder {
	return &UintBuilder{
		b: array.NewUint64Builder(mem),
	}
}
func (b *UintBuilder) Retain() {
	b.b.Retain()
}
func (b *UintBuilder) Release() {
	b.b.Release()
}
func (b *UintBuilder) Len() int {
	return b.b.Len()
}
func (b *UintBuilder) Cap() int {
	return b.b.Cap()
}
func (b *UintBuilder) Append(v uint64) {
	b.b.Append(v)
}
func (b *UintBuilder) AppendValues(v []uint64, valid []bool) {
	b.b.AppendValues(v, valid)
}
func (b *UintBuilder) UnsafeAppend(v uint64) {
	b.b.UnsafeAppend(v)
}
func (b *UintBuilder) NullN() int {
	return b.b.NullN()
}
func (b *UintBuilder) AppendNull() {
	b.b.AppendNull()
}
func (b *UintBuilder) UnsafeAppendBoolToBitmap(isValid bool) {
	b.b.UnsafeAppendBoolToBitmap(isValid)
}
func (b *UintBuilder) Reserve(n int) {
	b.b.Reserve(n)
}
func (b *UintBuilder) Resize(n int) {
	b.b.Resize(n)
}
func (b *UintBuilder) NewArray() Array {
	return b.NewUintArray()
}
func (b *UintBuilder) NewUintArray() *Uint {
	return b.b.NewUint64Array()
}

type FloatBuilder struct {
	b *array.Float64Builder
}

func NewFloatBuilder(mem memory.Allocator) *FloatBuilder {
	return &FloatBuilder{
		b: array.NewFloat64Builder(mem),
	}
}
func (b *FloatBuilder) Retain() {
	b.b.Retain()
}
func (b *FloatBuilder) Release() {
	b.b.Release()
}
func (b *FloatBuilder) Len() int {
	return b.b.Len()
}
func (b *FloatBuilder) Cap() int {
	return b.b.Cap()
}
func (b *FloatBuilder) Append(v float64) {
	b.b.Append(v)
}
func (b *FloatBuilder) AppendValues(v []float64, valid []bool) {
	b.b.AppendValues(v, valid)
}
func (b *FloatBuilder) UnsafeAppend(v float64) {
	b.b.UnsafeAppend(v)
}
func (b *FloatBuilder) NullN() int {
	return b.b.NullN()
}
func (b *FloatBuilder) AppendNull() {
	b.b.AppendNull()
}
func (b *FloatBuilder) UnsafeAppendBoolToBitmap(isValid bool) {
	b.b.UnsafeAppendBoolToBitmap(isValid)
}
func (b *FloatBuilder) Reserve(n int) {
	b.b.Reserve(n)
}
func (b *FloatBuilder) Resize(n int) {
	b.b.Resize(n)
}
func (b *FloatBuilder) NewArray() Array {
	return b.NewFloatArray()
}
func (b *FloatBuilder) NewFloatArray() *Float {
	return b.b.NewFloat64Array()
}

type BooleanBuilder struct {
	b *array.BooleanBuilder
}

func NewBooleanBuilder(mem memory.Allocator) *BooleanBuilder {
	return &BooleanBuilder{
		b: array.NewBooleanBuilder(mem),
	}
}
func (b *BooleanBuilder) Retain() {
	b.b.Retain()
}
func (b *BooleanBuilder) Release() {
	b.b.Release()
}
func (b *BooleanBuilder) Len() int {
	return b.b.Len()
}
func (b *BooleanBuilder) Cap() int {
	return b.b.Cap()
}
func (b *BooleanBuilder) Append(v bool) {
	b.b.Append(v)
}
func (b *BooleanBuilder) AppendValues(v []bool, valid []bool) {
	b.b.AppendValues(v, valid)
}
func (b *BooleanBuilder) UnsafeAppend(v bool) {
	b.b.UnsafeAppend(v)
}
func (b *BooleanBuilder) NullN() int {
	return b.b.NullN()
}
func (b *BooleanBuilder) AppendNull() {
	b.b.AppendNull()
}
func (b *BooleanBuilder) UnsafeAppendBoolToBitmap(isValid bool) {
	b.b.UnsafeAppendBoolToBitmap(isValid)
}
func (b *BooleanBuilder) Reserve(n int) {
	b.b.Reserve(n)
}
func (b *BooleanBuilder) Resize(n int) {
	b.b.Resize(n)
}
func (b *BooleanBuilder) NewArray() Array {
	return b.NewBooleanArray()
}
func (b *BooleanBuilder) NewBooleanArray() *Boolean {
	return b.b.NewBooleanArray()
}
