// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package common

import (
	"fmt"
)

// TestStaticContentHashObject

type TestStaticContentHashObject struct {
	*ContentHashObject
	Code int
}

func NewTestStaticContentHashObject(code int) *TestStaticContentHashObject {
	hashObj := &TestStaticContentHashObject{
		ContentHashObject: NewContentHashObject(),
		Code:              code,
	}
	return hashObj
}

func (self *TestStaticContentHashObject) GetHashSeed() string {
	return ""
}

func (self *TestStaticContentHashObject) GetHashCode() string {
	return fmt.Sprintf("%d", self.Code)
}
