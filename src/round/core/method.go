// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import (
	"round"
	"round/core/rpc"
)

const (
	errorBadExecMethod = "Method::Exec() is not overridded"
)

// A Method represents a node method.
type Method struct {
	Name string
}

// NewMethod returns a new Method.
func NewMethod(name string) *Method {
	method := &Method{Name: name}
	return method
}

// Exec runs the specified request on the local node.
func (self *Method) Exec(node *LocalNode, req *rpc.Request) (*rpc.Response, *round.Error) {
	return nil, round.NewError()
}
