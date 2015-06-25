// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

// A Method represents a node method.
type Method struct {
	Name string
}

// NewMethod returns a new Method.
func NewMethod(name string) *Method {
	method := &Method{Name: name}
	return method
}
