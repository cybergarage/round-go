// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import (
	"round/core/rpc"
)

// A Methoder represents a node interface.
type Method interface {
	GetName() string
	Exec(*LocalNode, *rpc.Request) (*rpc.Response, *rpc.Error)
}

// A Method represents a node method.
type BaseMethod struct {
	Name string
}

// NewBaseMethod returns a new BaseMethod.
func NewBaseMethod(name string) *BaseMethod {
	method := &BaseMethod{Name: name}
	return method
}

// GetName returns a name.
func (self *BaseMethod) GetName() string {
	return self.Name
}
