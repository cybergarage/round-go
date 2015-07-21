// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package common

import (
	"crypto/sha256"
)

// A HashObjecter represents a listener of HashObjecter.
type HashObjecter interface {
	GetHashSeed() string
}

// A HashObject represents a ContentHashObject.
type HashObject struct {
}

// NewHashObject returns a new NewHashObject.
func NewHashObject() *HashObject {
	hashObj := &HashObject{}
	return hashObj
}

// IsHashObject check whether the specified object has HashObjecter interface.
func IsHashObject(obj interface{}) (HashObjecter, bool) {
	hashObj, ok := obj.(HashObjecter)
	return hashObj, ok
}

// GetHashCode returns a hash code of the specified object.
func (self *HashObject) GetHashCode() (string, bool) {
	hashObj, ok := IsHashObject(self)
	if !ok {
		return "", false
	}
	hashSeed := hashObj.GetHashSeed()
	hashByte := sha256.Sum256([]byte(hashSeed))
	return string(hashByte[:]), true
}
