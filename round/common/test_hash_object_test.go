// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package common

import (
	"fmt"
	"math/rand"
)

// TestHashObject

type TestHashObject struct {
	*HashBaseObject
}

func NewTestHashObject() *TestHashObject {
	hashObj := &TestHashObject{}
	hashObj.HashBaseObject = NewHashBaseObject()
	return hashObj
}

func (self *TestHashObject) GetHashSeed() string {
	return fmt.Sprintf("%d", rand.Int)
}
