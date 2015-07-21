// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import (
	"fmt"
	"testing"

	"github.com/cybergarage/round-go/round/core/rpc"
)

const (
	errorMethodMapBadObject = "invalid object (%p) : expected (%p)"
	errorMethodMapBadLength = "invalid method len (%d) : expected (%d)"
)

type TestMethod struct {
	*BaseMethod
}

func NewTestMethod(name string) *TestMethod {
	method := &TestMethod{NewBaseMethod(name)}
	return method
}

func (self *TestMethod) Exec(*LocalNode, *rpc.Request) (*rpc.Response, *rpc.Error) {
	return nil, nil
}

func TestNewMethodMap(t *testing.T) {
	methodMap := NewMethodMap()

	if len(methodMap) != 0 {
		t.Errorf(errorMethodMapBadLength, len(methodMap), 0)
	}

	testLoopCnt := 10
	nameFmt := "name%d"
	methods := make([]Method, 0)

	// Add

	for n := 0; n < testLoopCnt; n++ {
		name := fmt.Sprintf(nameFmt, n)
		method := NewTestMethod(name)
		methodMap.Set(method)
		methods = append(methods, method)
		if len(methodMap) != (n + 1) {
			t.Errorf(errorMethodMapBadLength, len(methodMap), (n + 1))
		}
	}

	if len(methodMap) != testLoopCnt {
		t.Errorf(errorMethodMapBadLength, len(methodMap), testLoopCnt)
	}

	for n := 0; n < testLoopCnt; n++ {
		name := fmt.Sprintf(nameFmt, n)
		method, ok := methodMap.Get(name)
		if !ok {
			t.Errorf(errorMethodMapBadObject, method, methods[n])
		}
		if method != methods[n] {
			t.Errorf(errorMethodMapBadObject, method, methods[n])
		}
	}

	// Remove

	for n := 0; n < testLoopCnt; n++ {
		name := fmt.Sprintf(nameFmt, n)
		methodMap.Remove(name)
		if len(methodMap) != (testLoopCnt - (n + 1)) {
			t.Errorf(errorMethodMapBadLength, len(methodMap), (testLoopCnt - (n + 1)))
		}
	}

	if len(methodMap) != 0 {
		t.Errorf(errorMethodMapBadLength, len(methodMap), 0)
	}
}
