// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package common

import (
	"crypto/sha256"
)

// A ContentHashObjecter represents a listener of ContentHashObject.
type ContentHashObjecter interface {
	GetHashSeed() string
}

// A ContentHashObject represents a ContentHashObject.
type ContentHashObject struct {
}

// NewContentHashObject returns a new ContentHashObject.
func NewContentHashObject() *ContentHashObject {
	cluster := &ContentHashObject{}
	return cluster
}

// IsContentHashObject check whether the specified object has ContentHashObject interface.
func IsContentHashObject(obj interface{}) (ContentHashObjecter, bool) {
	hashObj, ok := obj.(ContentHashObjecter)
	return hashObj, ok
}

// GetHashCode returns a hash code of the specified object.
func (self *ContentHashObject) GetHashCode() (string, bool) {
	hashObj, ok := IsContentHashObject(self)
	if !ok {
		return "", false
	}
	hashSeed := hashObj.GetHashSeed()
	hashByte := sha256.Sum256([]byte(hashSeed))
	return string(hashByte[:]), true
}
