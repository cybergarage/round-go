// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

// A MethodMap represents a method map.
type MethodMap map[string]Method

// NewMethodMap returns a new MethodMap.
func NewMethodMap() MethodMap {
	methodMap := MethodMap{}
	return methodMap
}

// Set sets a method with the specified name.
func (self MethodMap) Set(method Method) error {
	self[method.GetName()] = method
	return nil
}

// Get gets a method by the specified name.
func (self MethodMap) Get(name string) (Method, bool) {
	method, ok := self[name]
	return method, ok
}

// Remove removes a method by the specified name.
func (self MethodMap) Remove(name string) bool {
	delete(self, name)
	return true
}
