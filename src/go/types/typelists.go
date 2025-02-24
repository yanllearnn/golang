// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

import "bytes"

// TParamList holds a list of type parameters.
type TParamList struct{ tparams []*TypeParam }

// Len returns the number of type parameters in the list.
// It is safe to call on a nil receiver.
func (l *TParamList) Len() int { return len(l.list()) }

// At returns the i'th type parameter in the list.
func (l *TParamList) At(i int) *TypeParam { return l.tparams[i] }

// list is for internal use where we expect a []*TypeParam.
// TODO(rfindley): list should probably be eliminated: we can pass around a
// TParamList instead.
func (l *TParamList) list() []*TypeParam {
	if l == nil {
		return nil
	}
	return l.tparams
}

// TypeList holds a list of types.
type TypeList struct{ types []Type }

// NewTypeList returns a new TypeList with the types in list.
func NewTypeList(list []Type) *TypeList {
	if len(list) == 0 {
		return nil
	}
	return &TypeList{list}
}

// Len returns the number of types in the list.
// It is safe to call on a nil receiver.
func (l *TypeList) Len() int { return len(l.list()) }

// At returns the i'th type in the list.
func (l *TypeList) At(i int) Type { return l.types[i] }

// list is for internal use where we expect a []Type.
// TODO(rfindley): list should probably be eliminated: we can pass around a
// TypeList instead.
func (l *TypeList) list() []Type {
	if l == nil {
		return nil
	}
	return l.types
}

func (l *TypeList) String() string {
	if l == nil || len(l.types) == 0 {
		return "[]"
	}
	var buf bytes.Buffer
	newTypeWriter(&buf, nil).typeList(l.types)
	return buf.String()
}

// ----------------------------------------------------------------------------
// Implementation

func bindTParams(list []*TypeParam) *TParamList {
	if len(list) == 0 {
		return nil
	}
	for i, typ := range list {
		if typ.index >= 0 {
			panic("type parameter bound more than once")
		}
		typ.index = i
	}
	return &TParamList{tparams: list}
}
