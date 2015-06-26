// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import (
	//"errors"
	"round/core/rpc"
)

// A MethodManager represents a manager for node methods.
type MethodManager struct {
	staticMethods  MethodMap
	dynamicMethods MethodMap
}

// NewMethodManager returns a new MethodManager.
func NewMethodManager() *MethodManager {
	methodMgr := &MethodManager{}
	methodMgr.staticMethods = NewMethodMap()
	methodMgr.staticMethods = NewMethodMap()
	return methodMgr
}

// IsStaticMethod returns whether the specified name method is exists as static methods.
func (self *MethodManager) IsStaticMethod(name string) bool {
	_, ok := self.staticMethods[name]
	return ok
}

// IsDynamicMethod returns whether the specified name method is exists as dynamic methods.
func (self *MethodManager) IsDynamicMethod(name string) bool {
	_, ok := self.dynamicMethods[name]
	return ok
}

// HasMethod returns whether the specified name method is exists.
func (self *MethodManager) HasMethod(name string) bool {
	if self.IsStaticMethod(name) {
		return true
	}
	if self.IsDynamicMethod(name) {
		return true
	}
	return false
}

// GetMethod returns a method by the specified name.
func (self *MethodManager) GetMethod(name string) (*Method, error) {
	method, ok := self.staticMethods[name]
	if ok {
		return method, nil
	}
	method, ok = self.dynamicMethods[name]
	if ok {
		return method, nil
	}
	return nil, nil
}

// ExecMethod returns a result of the specified method.
func (self *MethodManager) ExecMethod(node *LocalNode, name string, req *rpc.Request) (*rpc.Response, error) {
	method, err := self.GetMethod(name)
	if err != nil {
		return nil, err
	}
	return method.Exec(node, req)
}
