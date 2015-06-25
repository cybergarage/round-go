// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import (
//"errors"
)

// A MethodManager represents a manager for node methods.
type MethodManager struct {
}

// NewMethodManager returns a new MethodManager.
func NewMethodManager() *MethodManager {
	methodMgr := &MethodManager{}
	return methodMgr
}
